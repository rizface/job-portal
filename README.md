# HOW TO RUN THIS PROJECT (This project is not finished yet)

1. run docker-compose
```shell
// move to job portal service
cd job-portal

// run docker-compose
sudo docker-compose -f docker-compose.yaml up -d
```

2. execute mongodb container
```shell
sudo docker container exec -it job-portal bash
```

3. setting auth for mongodb
```shell
// use database admin
use admin 

// create user
db.createUser({user:"root", pwd: "root", roles:["userAdminAnyDatabase", "readWriteAnyDatabase"]})
```

4. re-create mongodb
```shell
sudo docker-compose -f docker-compose.yaml stop
```

5. modify docker-compose.yaml
```shell
// change mongodb database command
// from this
mongod
//to this
mongod --auth
```

6. run docker-compose
```shell
sudo docker-compose -f docker-compose.yaml up -d
```

7. move to email service folder
```shell
cd email-service
```

8. run email-service
```shell
go run main.go
```

9. run job-portal
```shell
cd job-portal
go run main.go
```


