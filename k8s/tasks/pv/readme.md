# 配置 Pod 以使用 PersistentVolume 作为存储

```bash

echo 'Hello workd' > /mnt/data/index.html


kubectl apply -f pv.yaml
kubectl apply -f pv-pod.yaml

```
