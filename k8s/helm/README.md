# Helm 使用

* Helm 能够动态生成Kubernetes配置文件（Deployment.yaml, service.yaml等），并调用kubectl自动执行部署；支持发布的版本管理；很大程度简化了Kubernetes应用的部署

## 相关概念

### Chart

* 应用的配置集合，包括各种 Kubernetes 对象的配置模板、参数定义、依赖关系、文档说 明等

### Release

* chart 的运行实例，代表了一个正在运行的应用
* 当 chart 被安装到 Kubernetes 集群，就生成 一个 release
* chart 能够多次安装到同一个集群，每次安装都是一个 release

## 相关命令

| 操作类型                                    | 命令                                          |
| ------------------------------------------- | --------------------------------------------- |
| 创建Chart                                   | helm create NAME [flags]                      |
| 添加仓库                                    | helm repo add repo-name repo-source           |
| 更新仓库                                    | helm repo update                              |
| 查看helm仓库列表                            | helm repo list                                |
| 查看本地已安装的包                          | helm list                                     |
| 查看helm版本                                | helm version                                  |
| 删除release                                 | helm delete  --purge loki                     |
| 指定charts版本                              | helm version                                  |
| 查看安装历史                                | helm historys apps                            |
| 升级版本                                    | helm upgrade -f  values.yaml apps stable/apps |
| 版本回滚                                    | helm rollback apps 1                          |
| 打包chart                                   | helm package mychart                          |
| 检查chart是否存在问题                       | helm lint mysql                               |
| 查看release状态                             | helm status mysql                             |
| 查看指定release的历史版本部署时部分配置信息 | helm get --revision 1 mysql                   |
| 对chart的模板和配置进行测试                 | helm install --dry-run --debug ./             |
| 展示可以使用charts可用                      | helm search                                   |
| 查看chart详细信息                           | helm inspect stable/apps                      |
| 查看release详情信息                         | helm get stable/apps                          |

## Chart详解

* 命令：helm create myMysql
* 相关文件
  * values.yaml 默认值配置
  * templates/  配置文件模版
  * Chart.yaml  chart的描述

### template

#### functions
  
* [functions](https://helm.sh/docs/chart_template_guide/function_list/)

#### Flow Control
  
##### IF/ELSE

* 条件控制

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  drink: {{ .Values.favorite.drink | default "tea" | quote }}
  food: {{ .Values.favorite.food | upper | quote }}
  {{- if eq .Values.favorite.drink "coffee" }}
  mug: true
  {{ else }}
  mug: false
  {{- end }}
```

##### WITH

* 限制作用域
  * **$** 指向 root scope
  * **.** 指向 current scope

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  {{- with .Values.favorite }}
  drink: {{ .drink | default "tea" | quote }}
  food: {{ .food | upper | quote }}
  release: {{ $.Release.Name }}
  {{- end }}
  release: {{ .Release.Name }}
```

##### RANGE

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  {{- with .Values.favorite }}
  drink: {{ .drink | default "tea" | quote }}
  food: {{ .food | upper | quote }}
  release: {{ $.Release.Name }}
  {{- end }}
  release: {{ .Release.Name }}

  toppings: |-
    {{- range .Values.pizzaToppings }}
    - {{ . | title | quote }}
    {{- end }}
```

### 变量

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  {{- $relname := .Release.Name -}}
  {{- with .Values.favorite }}
  drink: {{ .drink | default "tea" | quote }}
  food: {{ .food | upper | quote }}
  release: {{ $relname }}
  {{- end }}
```

### 模版定义

* 在模版文件中创建有名模版

```yaml
{{- define "mychart.labels" }}
  labels:
    generator: helm
    date: {{ now | htmlDate }}
{{- end }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
  {{- template "mychart.labels" }}
data:
  myvalue: "Hello World"
  {{- range $key, $val := .Values.favorite }}
  {{ $key }}: {{ $val | quote }}
  {{- end }}
```

* defined template with scope

```yaml
{{- define "mychart.labels" }}
  labels:
    generator: helm
    date: {{ now | htmlDate }}
    chart: {{ .Chart.Name | default "-" }}
    version: {{ .Chart.Version | default "1.0.0" }}
{{- end }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
  {{- template "mychart.labels" . }}
data:
  myvalue: "Hello World"
  {{- range $key, $val := .Values.favorite }}
  {{ $key }}: {{ $val | quote }}
  {{- end }}
```

* include function handles output formatting

```yaml
{{- define "mychart.app" -}}
app_name: {{ .Chart.Name }}
app_version: "{{ .Chart.Version }}"
{{- end -}}

apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
  labels:
{{ include "mychart.app" . | indent 4 }}
data:
  myvalue: "Hello World"
  {{- range $key, $val := .Values.favorite }}
  {{ $key }}: {{ $val | quote }}
  {{- end }}
{{ include "mychart.app" . | indent 2 }}
```

### 文件读取

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  {{- $files := .Files }}
  {{- range tuple "config1.toml" "config2.toml" "config3.toml" }}
  {{ . }}: |-
    {{ $files.Get . }}
  {{- end }}
```

* Path
  * functions: Base, Dir, Ext, IsAbs, Clean

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  {{- $files := .Files }}
  {{- range tuple "config1.toml" "config2.toml" "config3.toml" }}
  {{ . }}: |-
    {{ $files.Get . }}
  {{- end }}
```

## Subchart

* 见mychart/charts/mysubchart
