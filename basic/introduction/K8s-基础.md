



### Finalizer
一个常见的 Finalizer 的例子是 kubernetes.io/pv-protection， 
它用来防止意外删除 PersistentVolume 对象。 当一个 PersistentVolume 对象被 Pod 使用时，
Kubernetes 会添加 pv-protection Finalizer。 如果你试图删除 PersistentVolume，它将进入 Terminating 状态， 
但是控制器因为该 Finalizer 存在而无法删除该资源。 当 Pod 停止使用 PersistentVolume 时， 
Kubernetes 清除 pv-protection Finalizer，控制器就会删除该卷。
