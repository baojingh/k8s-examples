[toc]



# deployment







```bash
# 创建deployment
k create deployment mynginx --image=nginx --port=80

# 查看到deployment信息
kg deploy mynginx -owide

# 查看deployment详情信息，yaml展示
 kg deploy mynginx -oyaml


# 删除deployment
k delete deploy nginx-deployment
```





# pod

```
# 删除pod
k delete pod mynginx
```







