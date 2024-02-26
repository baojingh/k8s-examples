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
```





# pod

```bash
# 查看pod信息
kg pod nginx-deployment-55888b446c-qdl6x 
kg pod  nginx-deployment-55888b446c-tpl9z -owide
kg pod nginx-deployment-55888b446c-qdl6x -oyaml

# 删除pod【如果其对应的deployment没有删除，则此pod被删除后会被马上创建】
k delete pod nginx-deployment-55888b446c-tpl9z

#

```







