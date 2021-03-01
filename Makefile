start:
	go run main.go


# Running Database Server with docker container enverontment 

# MySQL
# Pull MySQL from docker hub
mysqlpull:
	docker pull mysql:8.0.23

# run local mysql docker container
# use command 'docker ps images' to check postgres version
mysql:
	docker run --name=mysql8 -e MYSQL_ROOT_PASSWORD=keepsecret -e MYSQL_DATABASE=gogolang -p 3306:3306 -d mysql:8.0.23


# Postgresql
postgrespull:
	docker pull postgres:latest

# run local postgress docker container
# use command 'docker ps images' to check postgres version
postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

# This is temporary use less as configs still not complete.
# MongoDB
mongopull:
	docker pull mongo:4.4.4

# run local mongodb docker container
# use command 'docker ps images' to check postgres version
mangodb:
	# docker run --name postgres13 -p 27017:27017 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d mongo:4.4.4

.PHONY: start mysqlpull mysql postgrespull postgres 