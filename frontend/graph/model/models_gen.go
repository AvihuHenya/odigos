// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Action interface {
	IsAction()
	GetID() string
	GetType() string
	GetName() *string
	GetNotes() *string
	GetDisable() bool
	GetSignals() []SignalType
}

type ActionInput struct {
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details string       `json:"details"`
}

type AddClusterInfoAction struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Name    *string        `json:"name,omitempty"`
	Notes   *string        `json:"notes,omitempty"`
	Disable bool           `json:"disable"`
	Signals []SignalType   `json:"signals"`
	Details []*ClusterInfo `json:"details"`
}

func (AddClusterInfoAction) IsAction()              {}
func (this AddClusterInfoAction) GetID() string     { return this.ID }
func (this AddClusterInfoAction) GetType() string   { return this.Type }
func (this AddClusterInfoAction) GetName() *string  { return this.Name }
func (this AddClusterInfoAction) GetNotes() *string { return this.Notes }
func (this AddClusterInfoAction) GetDisable() bool  { return this.Disable }
func (this AddClusterInfoAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type APIToken struct {
	Token     string `json:"token"`
	Name      string `json:"name"`
	IssuedAt  int    `json:"issuedAt"`
	ExpiresAt int    `json:"expiresAt"`
}

type ClusterCollectorAnalyze struct {
	Enabled              *EntityProperty `json:"enabled"`
	CollectorGroup       *EntityProperty `json:"collectorGroup"`
	Deployed             *EntityProperty `json:"deployed,omitempty"`
	DeployedError        *EntityProperty `json:"deployedError,omitempty"`
	CollectorReady       *EntityProperty `json:"collectorReady,omitempty"`
	DeploymentCreated    *EntityProperty `json:"deploymentCreated"`
	ExpectedReplicas     *EntityProperty `json:"expectedReplicas,omitempty"`
	HealthyReplicas      *EntityProperty `json:"healthyReplicas,omitempty"`
	FailedReplicas       *EntityProperty `json:"failedReplicas,omitempty"`
	FailedReplicasReason *EntityProperty `json:"failedReplicasReason,omitempty"`
}

type ClusterInfo struct {
	AttributeName        string  `json:"attributeName"`
	AttributeStringValue *string `json:"attributeStringValue,omitempty"`
}

type CodeAttributes struct {
	Column     *bool `json:"column,omitempty"`
	FilePath   *bool `json:"filePath,omitempty"`
	Function   *bool `json:"function,omitempty"`
	LineNumber *bool `json:"lineNumber,omitempty"`
	Namespace  *bool `json:"namespace,omitempty"`
	Stacktrace *bool `json:"stacktrace,omitempty"`
}

type CodeAttributesInput struct {
	Column     *bool `json:"column,omitempty"`
	FilePath   *bool `json:"filePath,omitempty"`
	Function   *bool `json:"function,omitempty"`
	LineNumber *bool `json:"lineNumber,omitempty"`
	Namespace  *bool `json:"namespace,omitempty"`
	Stacktrace *bool `json:"stacktrace,omitempty"`
}

type ComputePlatform struct {
	ComputePlatformType  ComputePlatformType    `json:"computePlatformType"`
	APITokens            []*APIToken            `json:"apiTokens"`
	K8sActualNamespaces  []*K8sActualNamespace  `json:"k8sActualNamespaces"`
	K8sActualNamespace   *K8sActualNamespace    `json:"k8sActualNamespace,omitempty"`
	Sources              *PaginatedSources      `json:"sources"`
	Destinations         []*Destination         `json:"destinations"`
	Actions              []*PipelineAction      `json:"actions"`
	InstrumentationRules []*InstrumentationRule `json:"instrumentationRules"`
}

type Condition struct {
	Status             ConditionStatus `json:"status"`
	Type               string          `json:"type"`
	Reason             *string         `json:"reason,omitempty"`
	Message            *string         `json:"message,omitempty"`
	LastTransitionTime *string         `json:"lastTransitionTime,omitempty"`
}

type ContainerRuntimeInfoAnalyze struct {
	ContainerName  *EntityProperty   `json:"containerName"`
	Language       *EntityProperty   `json:"language"`
	RuntimeVersion *EntityProperty   `json:"runtimeVersion"`
	EnvVars        []*EntityProperty `json:"envVars"`
}

type ContainerWorkloadManifestAnalyze struct {
	ContainerName *EntityProperty   `json:"containerName"`
	Devices       *EntityProperty   `json:"devices"`
	OriginalEnv   []*EntityProperty `json:"originalEnv"`
}

type CustomReadDataLabel struct {
	Condition string `json:"condition"`
	Title     string `json:"title"`
	Value     string `json:"value"`
}

type DbQueryPayloadCollection struct {
	MaxPayloadLength    *int  `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool `json:"dropPartialPayloads,omitempty"`
}

type DbQueryPayloadCollectionInput struct {
	MaxPayloadLength    *int  `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool `json:"dropPartialPayloads,omitempty"`
}

type DeleteAttribute struct {
	AttributeName string `json:"attributeName"`
}

type DeleteAttributeAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details []string     `json:"details"`
}

func (DeleteAttributeAction) IsAction()              {}
func (this DeleteAttributeAction) GetID() string     { return this.ID }
func (this DeleteAttributeAction) GetType() string   { return this.Type }
func (this DeleteAttributeAction) GetName() *string  { return this.Name }
func (this DeleteAttributeAction) GetNotes() *string { return this.Notes }
func (this DeleteAttributeAction) GetDisable() bool  { return this.Disable }
func (this DeleteAttributeAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type DestinationDetails struct {
	Type      string `json:"type"`
	URLString string `json:"urlString"`
	Fields    string `json:"fields"`
}

type DestinationInput struct {
	Name            string                `json:"name"`
	Type            string                `json:"type"`
	ExportedSignals *ExportedSignalsInput `json:"exportedSignals"`
	Fields          []*FieldInput         `json:"fields"`
}

type EntityProperty struct {
	Name    string  `json:"name"`
	Value   string  `json:"value"`
	Status  *string `json:"status,omitempty"`
	Explain *string `json:"explain,omitempty"`
}

type ErrorSamplerAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details string       `json:"details"`
}

func (ErrorSamplerAction) IsAction()              {}
func (this ErrorSamplerAction) GetID() string     { return this.ID }
func (this ErrorSamplerAction) GetType() string   { return this.Type }
func (this ErrorSamplerAction) GetName() *string  { return this.Name }
func (this ErrorSamplerAction) GetNotes() *string { return this.Notes }
func (this ErrorSamplerAction) GetDisable() bool  { return this.Disable }
func (this ErrorSamplerAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type ExportedSignalsInput struct {
	Traces  bool `json:"traces"`
	Metrics bool `json:"metrics"`
	Logs    bool `json:"logs"`
}

type Field struct {
	Name                 string                 `json:"name"`
	DisplayName          string                 `json:"displayName"`
	ComponentType        string                 `json:"componentType"`
	ComponentProperties  string                 `json:"componentProperties"`
	Secret               bool                   `json:"secret"`
	InitialValue         string                 `json:"initialValue"`
	RenderCondition      []string               `json:"renderCondition"`
	HideFromReadData     []string               `json:"hideFromReadData"`
	CustomReadDataLabels []*CustomReadDataLabel `json:"customReadDataLabels"`
}

type FieldInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetConfigResponse struct {
	Installation InstallationStatus `json:"installation"`
	Readonly     bool               `json:"readonly"`
}

type GetDestinationDetailsResponse struct {
	Fields []*Field `json:"fields"`
}

type HTTPPayloadCollection struct {
	MimeTypes           []*string `json:"mimeTypes,omitempty"`
	MaxPayloadLength    *int      `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool     `json:"dropPartialPayloads,omitempty"`
}

type HTTPPayloadCollectionInput struct {
	MimeTypes           []*string `json:"mimeTypes,omitempty"`
	MaxPayloadLength    *int      `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool     `json:"dropPartialPayloads,omitempty"`
}

type InstrumentationConfigAnalyze struct {
	Created    *EntityProperty                `json:"created"`
	CreateTime *EntityProperty                `json:"createTime,omitempty"`
	Containers []*ContainerRuntimeInfoAnalyze `json:"containers"`
}

type InstrumentationDeviceAnalyze struct {
	StatusText *EntityProperty                     `json:"statusText"`
	Containers []*ContainerWorkloadManifestAnalyze `json:"containers"`
}

type InstrumentationInstanceAnalyze struct {
	Healthy               *EntityProperty   `json:"healthy"`
	Message               *EntityProperty   `json:"message,omitempty"`
	IdentifyingAttributes []*EntityProperty `json:"identifyingAttributes"`
}

type InstrumentationLabelsAnalyze struct {
	Instrumented     *EntityProperty `json:"instrumented"`
	Workload         *EntityProperty `json:"workload,omitempty"`
	Namespace        *EntityProperty `json:"namespace,omitempty"`
	InstrumentedText *EntityProperty `json:"instrumentedText,omitempty"`
}

type InstrumentationLibraryGlobalID struct {
	Name     string               `json:"name"`
	SpanKind *SpanKind            `json:"spanKind,omitempty"`
	Language *ProgrammingLanguage `json:"language,omitempty"`
}

type InstrumentationLibraryGlobalIDInput struct {
	Name     string               `json:"name"`
	SpanKind *SpanKind            `json:"spanKind,omitempty"`
	Language *ProgrammingLanguage `json:"language,omitempty"`
}

type InstrumentationRule struct {
	RuleID                   string                            `json:"ruleId"`
	RuleName                 *string                           `json:"ruleName,omitempty"`
	Notes                    *string                           `json:"notes,omitempty"`
	Disabled                 *bool                             `json:"disabled,omitempty"`
	Mutable                  bool                              `json:"mutable"`
	ProfileName              string                            `json:"profileName"`
	Workloads                []*PodWorkload                    `json:"workloads,omitempty"`
	InstrumentationLibraries []*InstrumentationLibraryGlobalID `json:"instrumentationLibraries,omitempty"`
	PayloadCollection        *PayloadCollection                `json:"payloadCollection,omitempty"`
	CodeAttributes           *CodeAttributes                   `json:"codeAttributes,omitempty"`
}

type InstrumentationRuleInput struct {
	RuleName                 *string                                `json:"ruleName,omitempty"`
	Notes                    *string                                `json:"notes,omitempty"`
	Disabled                 *bool                                  `json:"disabled,omitempty"`
	Workloads                []*PodWorkloadInput                    `json:"workloads,omitempty"`
	InstrumentationLibraries []*InstrumentationLibraryGlobalIDInput `json:"instrumentationLibraries,omitempty"`
	PayloadCollection        *PayloadCollectionInput                `json:"payloadCollection,omitempty"`
	CodeAttributes           *CodeAttributesInput                   `json:"codeAttributes,omitempty"`
}

type K8sActualNamespace struct {
	Name             string             `json:"name"`
	Selected         bool               `json:"selected"`
	K8sActualSources []*K8sActualSource `json:"k8sActualSources"`
}

type K8sActualSource struct {
	Namespace         string                           `json:"namespace"`
	Name              string                           `json:"name"`
	Kind              K8sResourceKind                  `json:"kind"`
	NumberOfInstances *int                             `json:"numberOfInstances,omitempty"`
	Selected          *bool                            `json:"selected,omitempty"`
	OtelServiceName   *string                          `json:"otelServiceName,omitempty"`
	Containers        []*SourceContainerRuntimeDetails `json:"containers,omitempty"`
	Conditions        []*Condition                     `json:"conditions,omitempty"`
}

type K8sDesiredNamespaceInput struct {
	AutoInstrument *bool `json:"autoInstrument,omitempty"`
}

type K8sDesiredSourceInput struct {
	ServiceName    *string `json:"serviceName,omitempty"`
	AutoInstrument *bool   `json:"autoInstrument,omitempty"`
}

type K8sNamespaceID struct {
	Name string `json:"name"`
}

type K8sSourceID struct {
	Namespace string          `json:"namespace"`
	Kind      K8sResourceKind `json:"kind"`
	Name      string          `json:"name"`
}

type LatencySamplerAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details []*string    `json:"details"`
}

func (LatencySamplerAction) IsAction()              {}
func (this LatencySamplerAction) GetID() string     { return this.ID }
func (this LatencySamplerAction) GetType() string   { return this.Type }
func (this LatencySamplerAction) GetName() *string  { return this.Name }
func (this LatencySamplerAction) GetNotes() *string { return this.Notes }
func (this LatencySamplerAction) GetDisable() bool  { return this.Disable }
func (this LatencySamplerAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type MessagingPayloadCollection struct {
	MaxPayloadLength    *int  `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool `json:"dropPartialPayloads,omitempty"`
}

type MessagingPayloadCollectionInput struct {
	MaxPayloadLength    *int  `json:"maxPayloadLength,omitempty"`
	DropPartialPayloads *bool `json:"dropPartialPayloads,omitempty"`
}

type Mutation struct {
}

type NodeCollectorAnalyze struct {
	Enabled        *EntityProperty `json:"enabled"`
	CollectorGroup *EntityProperty `json:"collectorGroup"`
	Deployed       *EntityProperty `json:"deployed,omitempty"`
	DeployedError  *EntityProperty `json:"deployedError,omitempty"`
	CollectorReady *EntityProperty `json:"collectorReady,omitempty"`
	DaemonSet      *EntityProperty `json:"daemonSet"`
	DesiredNodes   *EntityProperty `json:"desiredNodes,omitempty"`
	CurrentNodes   *EntityProperty `json:"currentNodes,omitempty"`
	UpdatedNodes   *EntityProperty `json:"updatedNodes,omitempty"`
	AvailableNodes *EntityProperty `json:"availableNodes,omitempty"`
}

type OdigosAnalyze struct {
	OdigosVersion        *EntityProperty          `json:"odigosVersion"`
	KubernetesVersion    *EntityProperty          `json:"kubernetesVersion"`
	Tier                 *EntityProperty          `json:"tier"`
	InstallationMethod   *EntityProperty          `json:"installationMethod"`
	NumberOfDestinations int                      `json:"numberOfDestinations"`
	NumberOfSources      int                      `json:"numberOfSources"`
	ClusterCollector     *ClusterCollectorAnalyze `json:"clusterCollector"`
	NodeCollector        *NodeCollectorAnalyze    `json:"nodeCollector"`
	IsSettled            bool                     `json:"isSettled"`
	HasErrors            bool                     `json:"hasErrors"`
}

type OverviewMetricsResponse struct {
	Sources      []*SingleSourceMetricsResponse      `json:"sources"`
	Destinations []*SingleDestinationMetricsResponse `json:"destinations"`
}

type PaginatedSources struct {
	NextPage string             `json:"nextPage"`
	Items    []*K8sActualSource `json:"items"`
}

type PatchSourceRequestInput struct {
	OtelServiceName string `json:"otelServiceName"`
}

type PayloadCollection struct {
	HTTPRequest  *HTTPPayloadCollection      `json:"httpRequest,omitempty"`
	HTTPResponse *HTTPPayloadCollection      `json:"httpResponse,omitempty"`
	DbQuery      *DbQueryPayloadCollection   `json:"dbQuery,omitempty"`
	Messaging    *MessagingPayloadCollection `json:"messaging,omitempty"`
}

type PayloadCollectionInput struct {
	HTTPRequest  *HTTPPayloadCollectionInput      `json:"httpRequest,omitempty"`
	HTTPResponse *HTTPPayloadCollectionInput      `json:"httpResponse,omitempty"`
	DbQuery      *DbQueryPayloadCollectionInput   `json:"dbQuery,omitempty"`
	Messaging    *MessagingPayloadCollectionInput `json:"messaging,omitempty"`
}

type PersistNamespaceItemInput struct {
	Name           string `json:"name"`
	FutureSelected bool   `json:"futureSelected"`
}

type PersistNamespaceSourceInput struct {
	Name     string          `json:"name"`
	Kind     K8sResourceKind `json:"kind"`
	Selected bool            `json:"selected"`
}

type PiiMaskingAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details []string     `json:"details,omitempty"`
}

func (PiiMaskingAction) IsAction()              {}
func (this PiiMaskingAction) GetID() string     { return this.ID }
func (this PiiMaskingAction) GetType() string   { return this.Type }
func (this PiiMaskingAction) GetName() *string  { return this.Name }
func (this PiiMaskingAction) GetNotes() *string { return this.Notes }
func (this PiiMaskingAction) GetDisable() bool  { return this.Disable }
func (this PiiMaskingAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type PipelineAction struct {
	ID         string       `json:"id"`
	Type       string       `json:"type"`
	Spec       string       `json:"spec"`
	Conditions []*Condition `json:"conditions,omitempty"`
}

type PodAnalyze struct {
	PodName    *EntityProperty        `json:"podName"`
	NodeName   *EntityProperty        `json:"nodeName"`
	Phase      *EntityProperty        `json:"phase"`
	Containers []*PodContainerAnalyze `json:"containers"`
}

type PodContainerAnalyze struct {
	ContainerName            *EntityProperty                   `json:"containerName"`
	ActualDevices            *EntityProperty                   `json:"actualDevices"`
	InstrumentationInstances []*InstrumentationInstanceAnalyze `json:"instrumentationInstances"`
}

type PodWorkload struct {
	Namespace string          `json:"namespace"`
	Name      string          `json:"name"`
	Kind      K8sResourceKind `json:"kind"`
}

type PodWorkloadInput struct {
	Namespace string          `json:"namespace"`
	Kind      K8sResourceKind `json:"kind"`
	Name      string          `json:"name"`
}

type ProbabilisticSamplerAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details string       `json:"details"`
}

func (ProbabilisticSamplerAction) IsAction()              {}
func (this ProbabilisticSamplerAction) GetID() string     { return this.ID }
func (this ProbabilisticSamplerAction) GetType() string   { return this.Type }
func (this ProbabilisticSamplerAction) GetName() *string  { return this.Name }
func (this ProbabilisticSamplerAction) GetNotes() *string { return this.Notes }
func (this ProbabilisticSamplerAction) GetDisable() bool  { return this.Disable }
func (this ProbabilisticSamplerAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type Query struct {
}

type RenameAttributeAction struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	Name    *string      `json:"name,omitempty"`
	Notes   *string      `json:"notes,omitempty"`
	Disable bool         `json:"disable"`
	Signals []SignalType `json:"signals"`
	Details string       `json:"details"`
}

func (RenameAttributeAction) IsAction()              {}
func (this RenameAttributeAction) GetID() string     { return this.ID }
func (this RenameAttributeAction) GetType() string   { return this.Type }
func (this RenameAttributeAction) GetName() *string  { return this.Name }
func (this RenameAttributeAction) GetNotes() *string { return this.Notes }
func (this RenameAttributeAction) GetDisable() bool  { return this.Disable }
func (this RenameAttributeAction) GetSignals() []SignalType {
	if this.Signals == nil {
		return nil
	}
	interfaceSlice := make([]SignalType, 0, len(this.Signals))
	for _, concrete := range this.Signals {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type RuntimeInfoAnalyze struct {
	Generation *EntityProperty                `json:"generation"`
	Containers []*ContainerRuntimeInfoAnalyze `json:"containers"`
}

type SingleDestinationMetricsResponse struct {
	ID            string `json:"id"`
	TotalDataSent int    `json:"totalDataSent"`
	Throughput    int    `json:"throughput"`
}

type SingleSourceMetricsResponse struct {
	Namespace     string `json:"namespace"`
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	TotalDataSent int    `json:"totalDataSent"`
	Throughput    int    `json:"throughput"`
}

type SourceAnalyze struct {
	Name                  *EntityProperty               `json:"name"`
	Kind                  *EntityProperty               `json:"kind"`
	Namespace             *EntityProperty               `json:"namespace"`
	Labels                *InstrumentationLabelsAnalyze `json:"labels"`
	RuntimeInfo           *RuntimeInfoAnalyze           `json:"runtimeInfo"`
	InstrumentationConfig *InstrumentationConfigAnalyze `json:"instrumentationConfig"`
	InstrumentationDevice *InstrumentationDeviceAnalyze `json:"instrumentationDevice"`
	TotalPods             int                           `json:"totalPods"`
	PodsPhasesCount       string                        `json:"podsPhasesCount"`
	Pods                  []*PodAnalyze                 `json:"pods"`
}

type SourceContainerRuntimeDetails struct {
	ContainerName  string  `json:"containerName"`
	Language       string  `json:"language"`
	RuntimeVersion string  `json:"runtimeVersion"`
	OtherAgent     *string `json:"otherAgent,omitempty"`
}

type TestConnectionResponse struct {
	Succeeded       bool    `json:"succeeded"`
	StatusCode      int     `json:"statusCode"`
	DestinationType *string `json:"destinationType,omitempty"`
	Message         *string `json:"message,omitempty"`
	Reason          *string `json:"reason,omitempty"`
}

type ComputePlatformType string

const (
	ComputePlatformTypeK8s ComputePlatformType = "K8S"
	ComputePlatformTypeVM  ComputePlatformType = "VM"
)

var AllComputePlatformType = []ComputePlatformType{
	ComputePlatformTypeK8s,
	ComputePlatformTypeVM,
}

func (e ComputePlatformType) IsValid() bool {
	switch e {
	case ComputePlatformTypeK8s, ComputePlatformTypeVM:
		return true
	}
	return false
}

func (e ComputePlatformType) String() string {
	return string(e)
}

func (e *ComputePlatformType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ComputePlatformType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ComputePlatformType", str)
	}
	return nil
}

func (e ComputePlatformType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConditionStatus string

const (
	ConditionStatusTrue    ConditionStatus = "True"
	ConditionStatusFalse   ConditionStatus = "False"
	ConditionStatusUnknown ConditionStatus = "Unknown"
)

var AllConditionStatus = []ConditionStatus{
	ConditionStatusTrue,
	ConditionStatusFalse,
	ConditionStatusUnknown,
}

func (e ConditionStatus) IsValid() bool {
	switch e {
	case ConditionStatusTrue, ConditionStatusFalse, ConditionStatusUnknown:
		return true
	}
	return false
}

func (e ConditionStatus) String() string {
	return string(e)
}

func (e *ConditionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConditionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConditionStatus", str)
	}
	return nil
}

func (e ConditionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InstallationStatus string

const (
	InstallationStatusNew          InstallationStatus = "NEW"
	InstallationStatusAppsSelected InstallationStatus = "APPS_SELECTED"
	InstallationStatusFinished     InstallationStatus = "FINISHED"
)

var AllInstallationStatus = []InstallationStatus{
	InstallationStatusNew,
	InstallationStatusAppsSelected,
	InstallationStatusFinished,
}

func (e InstallationStatus) IsValid() bool {
	switch e {
	case InstallationStatusNew, InstallationStatusAppsSelected, InstallationStatusFinished:
		return true
	}
	return false
}

func (e InstallationStatus) String() string {
	return string(e)
}

func (e *InstallationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InstallationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InstallationStatus", str)
	}
	return nil
}

func (e InstallationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type K8sResourceKind string

const (
	K8sResourceKindDeployment  K8sResourceKind = "Deployment"
	K8sResourceKindDaemonSet   K8sResourceKind = "DaemonSet"
	K8sResourceKindStatefulSet K8sResourceKind = "StatefulSet"
)

var AllK8sResourceKind = []K8sResourceKind{
	K8sResourceKindDeployment,
	K8sResourceKindDaemonSet,
	K8sResourceKindStatefulSet,
}

func (e K8sResourceKind) IsValid() bool {
	switch e {
	case K8sResourceKindDeployment, K8sResourceKindDaemonSet, K8sResourceKindStatefulSet:
		return true
	}
	return false
}

func (e K8sResourceKind) String() string {
	return string(e)
}

func (e *K8sResourceKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = K8sResourceKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid K8sResourceKind", str)
	}
	return nil
}

func (e K8sResourceKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProgrammingLanguage string

const (
	ProgrammingLanguageUnspecified ProgrammingLanguage = "Unspecified"
	ProgrammingLanguageJava        ProgrammingLanguage = "Java"
	ProgrammingLanguageGo          ProgrammingLanguage = "Go"
	ProgrammingLanguageJavaScript  ProgrammingLanguage = "JavaScript"
	ProgrammingLanguagePython      ProgrammingLanguage = "Python"
	ProgrammingLanguageDotNet      ProgrammingLanguage = "DotNet"
)

var AllProgrammingLanguage = []ProgrammingLanguage{
	ProgrammingLanguageUnspecified,
	ProgrammingLanguageJava,
	ProgrammingLanguageGo,
	ProgrammingLanguageJavaScript,
	ProgrammingLanguagePython,
	ProgrammingLanguageDotNet,
}

func (e ProgrammingLanguage) IsValid() bool {
	switch e {
	case ProgrammingLanguageUnspecified, ProgrammingLanguageJava, ProgrammingLanguageGo, ProgrammingLanguageJavaScript, ProgrammingLanguagePython, ProgrammingLanguageDotNet:
		return true
	}
	return false
}

func (e ProgrammingLanguage) String() string {
	return string(e)
}

func (e *ProgrammingLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgrammingLanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgrammingLanguage", str)
	}
	return nil
}

func (e ProgrammingLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SignalType string

const (
	SignalTypeTraces  SignalType = "TRACES"
	SignalTypeMetrics SignalType = "METRICS"
	SignalTypeLogs    SignalType = "LOGS"
)

var AllSignalType = []SignalType{
	SignalTypeTraces,
	SignalTypeMetrics,
	SignalTypeLogs,
}

func (e SignalType) IsValid() bool {
	switch e {
	case SignalTypeTraces, SignalTypeMetrics, SignalTypeLogs:
		return true
	}
	return false
}

func (e SignalType) String() string {
	return string(e)
}

func (e *SignalType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SignalType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SignalType", str)
	}
	return nil
}

func (e SignalType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SpanKind string

const (
	SpanKindInternal SpanKind = "Internal"
	SpanKindServer   SpanKind = "Server"
	SpanKindClient   SpanKind = "Client"
	SpanKindProducer SpanKind = "Producer"
	SpanKindConsumer SpanKind = "Consumer"
)

var AllSpanKind = []SpanKind{
	SpanKindInternal,
	SpanKindServer,
	SpanKindClient,
	SpanKindProducer,
	SpanKindConsumer,
}

func (e SpanKind) IsValid() bool {
	switch e {
	case SpanKindInternal, SpanKindServer, SpanKindClient, SpanKindProducer, SpanKindConsumer:
		return true
	}
	return false
}

func (e SpanKind) String() string {
	return string(e)
}

func (e *SpanKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SpanKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SpanKind", str)
	}
	return nil
}

func (e SpanKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
