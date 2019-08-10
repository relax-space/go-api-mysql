
# ping-mysql

## environment
 - MYSQL_HOST=127.0.0.1
 - PORT=3306
 - USER_NAME=root
 - PASSWORD=1234

## docker-compose
```yml
version: "3"
services:
 mysql-server:
   container_name: test-mysql
   environment:
   - MYSQL_ROOT_PASSWORD=1234
   image: mysql:5.7.22
   ports:
   - 3306
   volumes:
   - ./database/mysql/:/docker-entrypoint-initdb.d
 ping-mysql-server:
   container_name: test-ping-mysql
   command: sh -c 'echo "wait mysql..." && /go/bin/wait-for.sh test-mysql:3306 -t 36000 -- ./ping-mysql'
   depends_on:
   - mysql-server
   environment:
   - MYSQL_HOST=test-mysql
   - PORT=3306
   - USER_NAME=root
   - PASSWORD=1234
   image: xiaoxinmiao/ping-mysql
   volumes: 
   - ./wait-for.sh:/go/bin/wait-for.sh
   ports:
   - 8080
```

[wait-for.sh](https://github.com/Eficode/wait-for)
