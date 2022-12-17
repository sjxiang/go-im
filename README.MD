

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