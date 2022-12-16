

# IM 即时通讯

技术栈
数据库 MongoDB 


docker run -d  --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo

docker start some-mongo

docker inspect some-mongo | grep IPAddress

172.17.0.2:27017

