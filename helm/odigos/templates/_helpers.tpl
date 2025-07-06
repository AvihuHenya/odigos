{{- define "utils.imagePrefix" -}}
{{- $defaultRegistry := "registry.odigos.io" -}}
{{- $redHatRegistry := "registry.connect.redhat.com/odigos" -}}
{{- if $.Values.imagePrefix -}}
    {{- $.Values.imagePrefix -}}
{{- else -}}
    {{- if $.Values.openshift.enabled -}}
        {{- $redHatRegistry -}}
    {{- else -}}
        {{- $defaultRegistry -}}
    {{- end -}}
{{- end -}}
{{- end -}}

{{- define "utils.imageName" -}}
{{- printf "%s/odigos-%s%s:%s" (include "utils.imagePrefix" .) .Component (ternary "-ubi9" "" $.Values.openshift.enabled) .Tag }}
{{- end -}}
{{/*
Returns "true" if any userInstrumentationEnvs.language is enabled or has env vars
*/}}
{{- define "utils.shouldRenderUserInstrumentationEnvs" -}}
  {{- $languages := .Values.userInstrumentationEnvs.languages | default dict }}
  {{- $shouldRender := false }}
  {{- range $lang, $config := $languages }}
    {{- if or $config.enabled $config.env }}
      {{- $shouldRender = true }}
    {{- end }}
  {{- end }}
  {{- print $shouldRender }}
{{- end }}

{{- define "odigos.secretExists" -}}
  {{- $sec   := lookup "v1" "Secret" .Release.Namespace "odigos-pro" -}}
  {{- $token := default "" .Values.onPremToken -}}
  {{- if or $sec (ne $token "") -}}
true
  {{- end -}}
{{- end -}}


{{/*
  Return cleaned Kubernetes version, keeping leading 'v', removing vendor suffix like -eks-...
  */}}
  {{- define "utils.cleanKubeVersion" -}}
  {{- regexReplaceAll "-.*" .Capabilities.KubeVersion.Version "" -}}
  {{- end }}


{{- define "odigos.profileDefaults" -}}
{{- $profile := "size_s" -}}
{{- if and .Values.profiles (gt (len .Values.profiles) 0) -}}
  {{- $profile = index .Values.profiles 0 -}}
{{- end -}}
{{- if eq $profile "size_s" }}
  requests:
    cpu: "150m"
    memory: "300Mi"
  limits:
    cpu: "300m"
    memory: "300Mi"
{{- else if eq $profile "size_m" }}
  requests:
    cpu: "250m"
    memory: "250Mi"
  limits:
    cpu: "500m"
    memory: "500Mi"
{{- else if eq $profile "size_l" }}
  requests:
    cpu: "500m"
    memory: "500Mi"
  limits:
    cpu: "750m"
    memory: "750Mi"
{{- end }}
{{- end }}

