[toc]



# pod

```bash
# 查看pod信息
kg pod nginx-deployment-55888b446c-qdl6x 
kg pod  nginx-deployment-55888b446c-tpl9z -owide
kg pod nginx-deployment-55888b446c-qdl6x -oyaml

# 删除pod【如果其对应的deployment没有删除，则此pod被删除后会被马上创建】
k delete pod nginx-deployment-55888b446c-tpl9z

# 查看pod的label
kg pod  -owide --show-labels
kg pod         --show-labels



# 创建pod
k create 


# 查看pod详情信息
k describe deploy  nginx1-deployment-55788c949-kmxtm

#k8s v1.18之前创建pod的方式，只有pod，不存在depolyment
kubectl run --image=nginx nginx

```



基于yaml文件创建pod

```bash
# 创建pod
k apply -f busybox.yaml
```



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-sleep
spec:
  containers:
  - name: busybox
    image: busybox:1.28
    args:
    - sleep
    - "1000000"
```







# deployment



```bash
# 创建deployment
k create deployment mynginx --image=nginx --port=80
k create deployment mynginx --image nginx --port 80

# 模拟创建deployment，不执行
k create deployment nginx-deployment --image=nginx:latest --dry-run=server/client -o yaml



# 查看到deployment信息
kg deploy mynginx -owide
# 查看deployment详情信息，yaml展示
 kg deploy mynginx -oyaml



# 修改deployment信息
k edit deploy  mynginx 



# 删除deployment
k delete deploy nginx-deployment

# 查看deployment详情信息
k describe deploy  nginx1-deployment
```





# Service

```bash
# 创建service，通过服务公开端口；nginx-app是deployment名称
kubectl expose deployment nginx-app --port=80 --name=nginx-http
```





# namespace

```bash
# namespace


# 获取集群列表
kg ns
kg namespace

# 创建ns
k create ns hello

# 删除ns
k delete ns hello

```





# top

```bash
# 查看pod资源
k top pods

# 查看node资源
k top nodes
```



# api-versions

```bash
# API组名/版本号
k api-versions
```



# configMap

```bash
# 查看configmap列表
kg cm


# 查看cm详情信息
k get cm kube-root-ca.crt -oyaml
```



# contexts

```bash

# context列表
 k config get-contexts
 
 # 当前的context名称
  k config current-context
  
  # 设置默认的上下文为 my-cluster-name
  kubectl config use-context my-cluster-name
  
  
```



# job

```bash
# 创建一个打印 “Hello World” 的 Job，创建的是pod，查看具体输出： k logs -f  hello-dzbt7
k create job hello --image=busybox -- echo "hello"

# 删除job
 k delete job hello

# 创建一个打印 “Hello World” 间隔 1 分钟的 CronJob；每分钟创建一个pod，查看pod日志即可看到输出。
 k create cronjob --image=busybox  --schedule="*/1 * * * *" -- echo "hello"

# 删除cronjob
k delete cronjob hello

```



# 集群信息

```bash
k cluster-info
```



# explain

```bash
kubectl explain pods.spec.containers
```







# REF

https://blog.csdn.net/u011095110/article/details/83545350

https://kubernetes.io/zh-cn/docs/reference/kubectl/quick-reference/#kubectl-context-and-configuration

