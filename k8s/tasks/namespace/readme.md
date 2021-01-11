# 管理集群

## 管理内存，CPU 和 API 资源

### 为命名空间配置 内存、CPU 的 limit、require

- 创建命名空间 kubectl create namespace default-limit-example
- 创建 LimitRange 和 Pod

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
spec:
  limits:
  - default:
      memory: 1Gi
      cpu: 800m
    defaultRequest:
      memory: 1Gi
      cpu: 800m
    max:
      memory: 1Gi
      cpu: 800m
    min:
      memory: 500Mi
      cpu: 200m
    type: Container
```

- 内存限制
  - 如果 Container 未指定自己的内存请求和限制，将为它指定默认的内存请求和限制
  - 验证 Container 的内存请求是否大于或等于500 MiB
  - 验证 Container 的内存限制是否小于或等于1 GiB
- CPU限制
  - 如果容器没有声明自己的 CPU 请求和限制，将为容器指定默认 CPU 请求和限制。
  - 核查容器声明的 CPU 请求确保其大于或者等于 200 millicpu。
  - 核查容器声明的 CPU 限制确保其小于或者等于 800 millicpu。

- 执行限制 kubectl apply -f memory-defaults.yaml --namespace default-limit-example
- 查看namespace：kubectl describe namespaces default-limit-example

```bash
Resource Limits
 Type       Resource  Min    Max   Default Request  Default Limit  Max Limit/Request Ratio
 ----       --------  ---    ---   ---------------  -------------  -----------------------
 Container  cpu       200m   800m  800m             800m           -
 Container  memory    500Mi  1Gi   1Gi              1Gi            -
```

### 为命名空间配置 内存和 CPU 配额

- 创建命名空间 kubectl create namespace quota-mem-cpu-example
- 创建 ResourceQuota

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: mem-cpu-demo
spec:
  hard:
    requests.cpu: "1"
    requests.memory: 1Gi
    limits.cpu: "2"
    limits.memory: 2Gi
```

- 应用quota：kubectl apply -f quota-mem-cpu.yaml --namespace quota-mem-cpu-example
- 查看 ResourceQuota 详情：kubectl get resourcequota mem-cpu-demo --namespace=quota-mem-cpu-example --output=yaml
- ResourceQuota 在 quota-mem-cpu-example 命名空间中设置了如下要求
  - 每个容器必须有内存请求和限制，以及 CPU 请求和限制
  - 所有容器的内存请求总和不能超过1 GiB
  - 所有容器的内存限制总和不能超过2 GiB
  - 所有容器的 CPU 请求总和不能超过1 cpu
  - 所有容器的 CPU 限制总和不能超过2 cpu

### 创建POD

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app-pod
spec:
  containers:
  - name: app
    image: app:1.0.3
    imagePullPolicy: Never
    resources:
      limits:
        memory: "800Mi"
        cpu: "800m"
      requests:
        memory: "600Mi"
        cpu: "400m"
```

- 创建POD：kubectl apply -f pod.yaml --namespace quota-mem-cpu-example
- 再次查看 ResourceQuota 详情，可以看到 Pod 的内存和 CPU 请求值及限制值没有超过配额

```bash
spec:
  hard:
    limits.cpu: "2"
    limits.memory: 2Gi
    requests.cpu: "1"
    requests.memory: 1Gi
status:
  hard:
    limits.cpu: "2"
    limits.memory: 2Gi
    requests.cpu: "1"
    requests.memory: 1Gi
  used:
    limits.cpu: 800m
    limits.memory: 800Mi
    requests.cpu: 400m
```

### 配置命名空间下 Pod 配额

- 创建namespace：kubectl create namespace quota-pod-example
- 创建 ResourceQuota

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: pod-demo
spec:
  hard:
    pods: "2"
```

- 创建创建这个 ResourceQuota：kubectl apply -f quota-pod.yaml --namespace=quota-pod-example
- 查看资源配额：kubectl get resourcequota pod-demo --namespace=quota-pod-example --output=yaml

```yaml
spec:
  hard:
    pods: "2"
status:
  hard:
    pods: "2"
  used:
    pods: "0"
```

- 创建deployment
  - kubectl apply -f quota-pod-deployment.yaml --namespace=quota-pod-example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  selector:
    matchLabels:
      purpose: quota-demo
  replicas: 3
  template:
    metadata:
      labels:
        purpose: quota-demo
    spec:
      containers:
      - name: app
        image: app:1.0.3
        imagePullPolicy: Never
```

- 查看deployment详情
  - kubectl get deployment app --namespace=quota-pod-example --output=yaml

```yaml
spec:
  progressDeadlineSeconds: 600
  replicas: 3
  revisionHistoryLimit: 10
...
status:
  availableReplicas: 2
...
  lastUpdateTime: "2021-01-08T14:01:51Z"
  message: 'pods "app-569d8b5968-x7z4r" is forbidden: exceeded quota: pod-demo,
    requested: pods=1, used: pods=2, limited: pods=2'
```
