# 基础知识

### metadata
在 Kubernetes 配置文件中，metadata.labels 和 spec.selector 通常在一起使用，。他们起着不同的作用：
metadata.labels：这些标签（labels）附加在你创建的对象（例如 Pod、Service、或 Deployment）上。标签是键值对，可以被用来组织和分类这些对象。
spec.selector：这个字段定义了如何找到你想要该 Kubernetes 对象（例如 Service 或 Deployment）管理的 Pod。
Selector 包含一组键值对，只有那些标签与 selector 完全匹配的 Pod 才会被选中。

所以，metadata.labels 是你给你的 Kubernetes 对象打的标签，而 spec.selector 是你定义的规则，告诉 Kubernetes 去找哪些 Pod。

以下是一个 Deployment 的配置文件示例，用于演示这两个字段的使用：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  labels:
    app: my-app-deployment-label
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
        - name: my-redis
          image: redis
```
在这个示例中，我们创建了一个名为 “my-deployment” 的 Deployment，它的标签为 {app: "my-app-deployment-label"}。
在 spec.selector.matchLabels 中，我们定义了 {app: "my-app"}，
这意味着这个 Deployment 将管理所有标签为 {app: "my-app"} 的 Pod。

在 spec.template.metadata.labels 中，我们定义了这个 Deployment 创建的 Pod 的标签也为 {app: "my-app"}。

因此，这个 Deployment 将管理所有标签为 {app: "my-app"} 的 Pod，并且它创建的所有新 Pod 的标签都将为 {app: "my-app"}。



```yaml
$ kgd    --show-labels
NAME            READY   UP-TO-DATE   AVAILABLE   AGE     LABELS
my-deployment   3/3     3            3           13m     app=my-app-deployment-label

$ kgp    --show-labels
NAME                            READY   STATUS             RESTARTS        AGE   LABELS
my-deployment-8c944cd96-4ffhr   1/1     Running            0               14m   app=my-app,pod-template-hash=8c944cd96
my-deployment-8c944cd96-4p858   1/1     Running            0               14m   app=my-app,pod-template-hash=8c944cd96
my-deployment-8c944cd96-fn2rt   1/1     Running            0               14m   app=my-app,pod-template-hash=8c944cd96
```



```yaml
$ kde deploy my-deployment
Name:                   my-deployment
Namespace:              default
CreationTimestamp:      Thu, 07 Mar 2024 09:28:28 +0800
Labels:                 app=my-app-deployment-label
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               app=my-app

....

...
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

