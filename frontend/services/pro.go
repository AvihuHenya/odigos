package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/odigos-io/odigos/api/k8sconsts"
	"github.com/odigos-io/odigos/frontend/graph"
	"github.com/odigos-io/odigos/frontend/graph/model"
	"github.com/odigos-io/odigos/frontend/kube"
	"github.com/odigos-io/odigos/k8sutils/pkg/env"
	"github.com/odigos-io/odigos/k8sutils/pkg/pro"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// type ReportPayload struct {
// 	CustomerID       string `json:"customer_id"`
// 	TotalNodesCount  int    `json:"node_count"`
// 	LabeledNodeCount int    `json:"labeled_node_count"`
// 	Timestamp        string `json:"timestamp"`
// 	ReportSource     string `json:"report_source"` // "UI-browser" or "component-periodic"
// }

func countNodesInfo(ctx context.Context) (totalNodesCount int, labeledNodeCount int, err error) {
	// List all nodes (no label selector)
	nodeList, err := kube.DefaultClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return 0, 0, fmt.Errorf("failed to list nodes: %w", err)
	}

	// Count how many nodes have the desired label
	for _, node := range nodeList.Items {
		if val, ok := node.Labels[k8sconsts.OdigletEnterpriseInstalledLabel]; ok && val == k8sconsts.OdigletInstalledLabelValue {
			labeledNodeCount++
		}
	}

	return len(nodeList.Items), labeledNodeCount, nil
}

func SendUsageDetailes(ctx context.Context) *model.OdigosUsage {
	totalNodesCount, labeledNodeCount, err := countNodesInfo(ctx)
	if err != nil {
		log.Fatal("Failed to count Kubernetes nodes for reporting", "error", err)
		return nil
	}

	ns := env.GetCurrentNamespace()
	secret, err := kube.DefaultClient.CoreV1().Secrets(ns).Get(ctx, k8sconsts.OdigosProSecretName, metav1.GetOptions{})
	if err != nil {
		// Todo
		return nil
	}

	tokenBytes, ok := secret.Data["token"]
	if !ok {
		log.Println("Token not found in Odigos Pro secret")
		return nil
	}

	payload, err := graph.ExtractJWTPayload(string(tokenBytes))
	if err != nil {
		log.Println("Failed to extract JWT payload from token", "error", err)
		return nil
	}

	iss, _ := payload["iss"].(string)
	sub, _ := payload["sub"].(string)
	aud, _ := payload["aud"].(string)

	// Parse and validate JWT for customer_id
	// (You'll need a JWT parsing library, e.g., github.com/golang-jwt/jwt/v5)
	// For simplicity, let's assume customerID is extracted
	customerID := "avihu" // Replace with actual JWT parsing

	payload := &model.OdigosUsage{
		CustomerID:       aud,
		TotalNodesCount:  totalNodesCount,
		LabeledNodeCount: labeledNodeCount,
		Timestamp:        time.Now().Format(time.RFC3339),
		ReportSource:     "UI-browser",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal report payload", "error", err)
		return nil
	}

	// todo: do we need to save it in secret
	// **IMPORTANT:** Replace with your actual AWS API Gateway endpoint
	// awsAPIGatewayURL := os.Getenv("AWS_API_GATEWAY_URL") // Get from environment variable
	// if awsAPIGatewayURL == "" {
	// 	log.Println("AWS_API_GATEWAY_URL environment variable not set. Usage report not sent.")
	// 	return
	// }

	req, err := http.NewRequestWithContext(ctx, "POST", "https://39er97hl30.execute-api.us-east-1.amazonaws.com", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Failed to create HTTP request to AWS API Gateway", "error", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second} // Add a timeout
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send usage report to AWS API Gateway", "error", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Println("AWS API Gateway returned non-success status", "status", resp.Status, "response", string(bodyBytes))
	} else {
		log.Println("Usage report sent successfully to AWS API Gateway", "nodeCount", totalNodesCount)
	}

	return payload
}

// func PeriodicReport(ctx context.Context) {
// 	ticker := time.NewTicker(24 * time.Hour)
// 	defer ticker.Stop()

// 	// Run immediately at startup
// 	log.Println("Sending initial usage report...")
// 	SendUsageReport(ctx)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Usage reporting goroutine shutting down...")
// 			return
// 		case <-ticker.C:
// 			log.Println("Sending scheduled usage report...")
// 			SendUsageReportToAWS(ctx)
// 		}
// 	}
// }

func UpdateToken(c *gin.Context) {
	var request pro.TokenPayload

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	onPremToken := request.OnpremToken
	if onPremToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "onprem-token is required",
		})
		return
	}
	ctx := c.Request.Context()

	err := pro.UpdateOdigosToken(ctx, kube.DefaultClient, env.GetCurrentNamespace(), onPremToken)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	statusCode := 200
	acceptHeader := c.GetHeader("Accept")

	if acceptHeader == "application/json" {
		c.JSON(statusCode, struct{}{})
	} else {
		c.String(statusCode, "")
	}
}
