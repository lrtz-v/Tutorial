# 配置 Pod 使用projected作存储

```yaml
# 创建包含用户名和密码的文件:
echo -n "admin" > ./username.txt
echo -n "1f2d1e2e67df" > ./password.txt-->

# 将上述文件引用到 Secret：
kubectl create secret generic user --from-file=./username.txt
kubectl create secret generic pass --from-file=./password.txt


kubectl create -f projected.yaml
```
