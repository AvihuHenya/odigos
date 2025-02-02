package v1alpha1

import (
	"github.com/odigos-io/odigos/api/odigos/v1alpha1/instrumentationrules"
	"github.com/odigos-io/odigos/common"
	"go.opentelemetry.io/otel/attribute"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// InstrumentationConfig is the Schema for the instrumentationconfig API
type InstrumentationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstrumentationConfigSpec   `json:"spec,omitempty"`
	Status InstrumentationConfigStatus `json:"status,omitempty"`
}

type OtherAgent struct {
	Name string `json:"name,omitempty"`
}

// +kubebuilder:object:generate=true
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// +kubebuilder:object:generate=true
type RuntimeDetailsByContainer struct {
	ContainerName  string                     `json:"containerName"`
	Language       common.ProgrammingLanguage `json:"language"`
	RuntimeVersion string                     `json:"runtimeVersion,omitempty"`
	EnvVars        []EnvVar                   `json:"envVars,omitempty"`
	OtherAgent     *OtherAgent                `json:"otherAgent,omitempty"`
	LibCType       *common.LibCType           `json:"libCType,omitempty"`

	// Stores the error message from the CRI runtime if returned to prevent instrumenting the container if an error exists.
	CriErrorMessage *string `json:"criErrorMessage,omitempty"`
	// Holds the environment variables retrieved from the container runtime.
	EnvFromContainerRuntime []EnvVar `json:"envFromContainerRuntime,omitempty"`
	// A temporary variable used during migration to track whether the new runtime detection process has been executed. If empty, it indicates the process has not yet been run. This field may be removed later.
	RuntimeUpdateState *ProcessingState `json:"runtimeUpdateState,omitempty"`
}

type InstrumentationConfigStatus struct {
	// Capture Runtime Details for the workloads that this CR applies to.
	RuntimeDetailsByContainer []RuntimeDetailsByContainer `json:"runtimeDetailsByContainer,omitempty"`

	// Represents the observations of a InstrumentationConfig's current state.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" protobuf:"bytes,1,rep,name=conditions"`
}

// Config for the OpenTelemeetry SDKs that should be applied to a workload.
// The workload is identified by the owner reference
type InstrumentationConfigSpec struct {

	// the service.name property is used to populate the `service.name` resource attribute in the telemetry generated by this workload
	ServiceName string `json:"serviceName,omitempty"`

	// config for this workload.
	// the config is a list to allow for multiple config options and values to be applied.
	// the list is processed in order, and the first matching config is applied.
	Config []WorkloadInstrumentationConfig `json:"config,omitempty"`

	// Configuration for the OpenTelemetry SDKs that this workload should use.
	// The SDKs are identified by the programming language they are written in.
	// TODO: consider adding more granular control over the SDKs, such as community/enterprise, native/ebpf.
	SdkConfigs []SdkConfig `json:"sdkConfigs,omitempty"`
}

type SdkConfig struct {

	// The language of the SDK being configured
	Language common.ProgrammingLanguage `json:"language"`

	// configurations for the instrumentation libraries the the SDK should use
	InstrumentationLibraryConfigs []InstrumentationLibraryConfig `json:"instrumentationLibraryConfigs,omitempty"`

	// HeadSamplingConfig is a set sampling rules.
	// This config currently only applies to root spans.
	// In the Future we might add another level of configuration base on the parent span (ParentBased Sampling)
	HeadSamplingConfig *HeadSamplingConfig `json:"headSamplerConfig,omitempty"`

	DefaultPayloadCollection *instrumentationrules.PayloadCollection `json:"payloadCollection"`
}

// 'Operand' represents the attributes and values that an operator acts upon in an expression
type AttributeCondition struct {
	// attribute key (e.g. "url.path")
	Key attribute.Key `json:"key"`
	// currently only string values are supported.
	Val string `json:"val"`
	// The operator to use to compare the attribute value.
	Operator Operator `json:"operator,omitempty"`
}

// +kubebuilder:validation:Enum=equals;notEquals;endWith;startWith
// +kubebuilder:default:=equals
type Operator string

const (
	Equals    Operator = "equals"
	NotEquals Operator = "notEquals"
	EndWith   Operator = "endWith"
	StartWith Operator = "startWith"
)

// AttributesAndSamplerRule is a set of AttributeCondition that are ANDed together.
// If all attribute conditions evaluate to true, the AND sampler evaluates to true,
// and the fraction is used to determine the sampling decision.
// If any of the attribute compare samplers evaluate to false,
// the fraction is not used and the rule is skipped.
// An "empty" AttributesAndSamplerRule with no attribute conditions is considered to always evaluate to true.
// and the fraction is used to determine the sampling decision.
// This entity is refered to a rule in Odigos terminology for head-sampling.
type AttributesAndSamplerRule struct {
	AttributeConditions []AttributeCondition `json:"attributeConditions"`
	// The fraction of spans to sample, in the range [0, 1].
	// If the fraction is 0, no spans are sampled.
	// If the fraction is 1, all spans are sampled.
	// +kubebuilder:default:=1
	Fraction float64 `json:"fraction"`
}

// HeadSamplingConfig is a set of attribute rules.
// The first attribute rule that evaluates to true is used to determine the sampling decision based on its fraction.
//
// If none of the rules evaluate to true, the fallback fraction is used to determine the sampling decision.
type HeadSamplingConfig struct {
	AttributesAndSamplerRules []AttributesAndSamplerRule `json:"attributesAndSamplerRules"`
	// Used as a fallback if all rules evaluate to false,
	// it may be empty - in this case the default value will be 1 - all spans are sampled.
	// it should be a float value in the range [0, 1] - the fraction of spans to sample.
	// a value of 0 means no spans are sampled if none of the rules evaluate to true.
	// +kubebuilder:default:=1
	FallbackFraction float64 `json:"fallbackFraction"`
}

type InstrumentationLibraryConfig struct {
	InstrumentationLibraryId InstrumentationLibraryId `json:"libraryId"`

	TraceConfig *InstrumentationLibraryConfigTraces `json:"traceConfig,omitempty"`

	PayloadCollection *instrumentationrules.PayloadCollection `json:"payloadCollection,omitempty"`
}

type InstrumentationLibraryId struct {
	// The name of the instrumentation library
	// - Node.js: The name of the npm package: `@opentelemetry/instrumentation-<name>`
	InstrumentationLibraryName string `json:"libraryName"`
	// SpanKind is only supported by Golang and will be ignored for any other SDK language.
	// In Go, SpanKind is used because the same instrumentation library can be utilized for different span kinds (e.g., client/server).
	SpanKind common.SpanKind `json:"spanKind,omitempty"`
}

type InstrumentationLibraryConfigTraces struct {
	// Whether the instrumentation library is enabled to record traces.
	// When false, it is expected that the instrumentation library does not produce any spans regardless of any other configuration.
	// When true, the instrumentation library should produce spans according to the other configuration options.
	// If not specified, the default value for this signal should be used (whether to enable libraries by default or not).
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkloadInstrumentationConfig defined a single config option to apply
// on a workload, along with it's value, filters and instrumentation libraries
type WorkloadInstrumentationConfig struct {

	// OptionKey is the name of the option
	// This value is transparent to the CRD and is passed as-is to the SDK.
	OptionKey string `json:"optionKey"`

	// This option allow to specify the config option for a specific span kind
	// for example, only to client spans or only to server spans.
	// it the span kind is not specified, the option will apply to all spans.
	SpanKind common.SpanKind `json:"spanKind,omitempty"`

	// OptionValueBoolean is the boolean value of the option if it is a boolean
	OptionValueBoolean bool `json:"optionValueBoolean,omitempty"`

	// a list of instrumentation libraries to apply this setting to
	// if a library is not in this list, the setting should not apply to it
	// and should be cleared.
	InstrumentationLibraries []InstrumentationLibrary `json:"instrumentationLibraries"`
}

// InstrumentationLibrary represents a library for instrumentation
type InstrumentationLibrary struct {
	// Language is the programming language of the library
	Language common.ProgrammingLanguage `json:"language"`

	// InstrumentationLibraryName is the name of the instrumentation library
	InstrumentationLibraryName string `json:"instrumentationLibraryName"`
}

// +kubebuilder:object:root=true

// InstrumentationConfigList contains a list of InstrumentationOption
type InstrumentationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstrumentationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstrumentationConfig{}, &InstrumentationConfigList{})
}

// Languages returns the set of languages that this configuration applies to
func (ic *InstrumentationConfig) Languages() map[common.ProgrammingLanguage]struct{} {
	langs := make(map[common.ProgrammingLanguage]struct{})
	for _, sdkConfig := range ic.Spec.SdkConfigs {
		langs[sdkConfig.Language] = struct{}{}
	}
	return langs
}
