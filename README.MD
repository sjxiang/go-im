

# IM 即时通讯

技术栈
数据库 NoSQL MongoDB 

```dockerfile

docker run -d  --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo


docker start some-mongo
docker stop  some-mongo


docker inspect some-mongo | grep IPAddress
```


```
IM 

用户模块
    用户登录
    发送验证码
    用户注册
    用户详情


通讯模块

    发送、接受消息

    一对一通讯
    多对多通讯
    消息列表
    聊天记录列表
    添加好友
    删除好友


```

```text

gorilla/websocket 于 2022-12-10 不再维护

```