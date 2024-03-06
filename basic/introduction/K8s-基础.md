# 基础知识

### metadata
spec.selector.matchlabels
这3个字段，共同构成了一个label selector，也就是标签选择器，
这里我们可以看到matchlabels字段后面跟的值是app：nginx。这是一个k-v结构，带有这个k-v结构的Pod对象，
将会被这个Deployment管理

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.7.9
          ports:
            - containerPort: 80

```




### Finalizer
一个常见的 Finalizer 的例子是 kubernetes.io/pv-protection， 
它用来防止意外删除 PersistentVolume 对象。 当一个 PersistentVolume 对象被 Pod 使用时，
Kubernetes 会添加 pv-protection Finalizer。 如果你试图删除 PersistentVolume，它将进入 Terminating 状态， 
但是控制器因为该 Finalizer 存在而无法删除该资源。 当 Pod 停止使用 PersistentVolume 时， 
Kubernetes 清除 pv-protection Finalizer，控制器就会删除该卷。
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: xyz
spec:
  finalizers:
    - kubernetes
```

