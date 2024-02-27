[toc]



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







