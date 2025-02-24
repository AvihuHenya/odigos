{{- define "utils.imageName" -}}
{{- $defaultRegistry := "docker.io/keyval" -}}
{{- $redHatRegistry := "registry.connect.redhat.com/odigos" -}}
{{- if $.Values.imagePrefix -}}
    {{- $.Values.imagePrefix -}}/keyval/
{{- else -}}
    {{- if $.Values.openshift.enabled -}}
        {{- $redHatRegistry -}}/
    {{- else -}}
        {{- $defaultRegistry -}}/
    {{- end -}}
{{- end -}}
odigos-
{{- .Component -}}
{{- if $.Values.openshift.enabled -}}
-ubi9
{{- end -}}
:
{{- .Tag -}}
{{- end -}}
