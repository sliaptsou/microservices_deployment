{{/* vim: set filetype=mustache: */}}

{{/*
Selector labels
*/}}
{{- define "example-chart.selectorLabels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Chart name and version as used by the chart label
*/}}
{{- define "example-chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "example-chart.labels" -}}
helm.sh/chart: {{ include "example-chart.chart" . }}
{{ include "example-chart.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Annotation to update pods on Secrets or ConfigMaps updates
*/}}
{{- define "example-chart.propertiesHash" -}}
{{- $backendConfig := include (print $.Template.BasePath "/backend-configMap.yaml") . | sha256sum -}}
{{- $gatewayConfig := include (print $.Template.BasePath "/gateway-configMap.yaml") . | sha256sum -}}
{{ print $backendConfig $gatewayConfig | sha256sum }}
{{- end -}}

{{/*
Names of backend tier components
*/}}
{{- define "example-chart.backend.defaultName" -}}
{{- printf "backend-%s" .Release.Name -}}
{{- end -}}

{{- define "example-chart.backend.deployment.name" -}}
{{- default (include "example-chart.backend.defaultName" .) .Values.backend.deployment.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.backend.container.name" -}}
{{- default (include "example-chart.backend.defaultName" .) .Values.backend.container.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.backend.service.name" -}}
{{- default (include "example-chart.backend.defaultName" .) .Values.backend.service.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.backend.hpa.name" -}}
{{- default (include "example-chart.backend.defaultName" .) .Values.backend.hpa.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Names of gateway tier components
*/}}
{{- define "example-chart.gateway.defaultName" -}}
{{- printf "gateway-%s" .Release.Name -}}
{{- end -}}

{{- define "example-chart.gateway.deployment.name" -}}
{{- default (include "example-chart.gateway.defaultName" .) .Values.gateway.deployment.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.gateway.container.name" -}}
{{- default (include "example-chart.gateway.defaultName" .) .Values.gateway.container.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.gateway.service.name" -}}
{{- default (include "example-chart.gateway.defaultName" .) .Values.gateway.service.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "example-chart.gateway.hpa.name" -}}
{{- default (include "example-chart.gateway.defaultName" .) .Values.gateway.hpa.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Names of other components
*/}}

{{/*{{- define "example-chart.backend.defaultName" -}}*/}}
{{/*{{- printf "backend-url-config-%s" .Release.Name -}}*/}}
{{/*{{- end -}}*/}}
{{/**/}}
{{/*{{- define "example-chart.gateway.defaultName" -}}*/}}
{{/*{{- printf "gateway-url-config-%s" .Release.Name -}}*/}}
{{/*{{- end -}}*/}}
