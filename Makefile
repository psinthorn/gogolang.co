start:
	go run main.go


# Running Database Server with docker container enverontment 

# -----------------------------
# Cassandra
# -----------------------------
# Pull cassandra image from docker hub
cassandrapull:
	docker pull cassandra:4.0

# run local cassandra docker container
# use command 'docker ps images' to check postgres version
cassandra:
	docker run --name cassandra4 -p 7000:7000 -p 7001:7001 -p 7199:7199 -p 9160:9160 -p 9042:9042 -d cassandra:4.0

# To access terminal for run bash command in docker 
cassandraterminal:
	docker exec -it cassandra4 bash && cqlsh

# Run cassandra sql command 
cqlsh:
	cqlsh


# Start and stop docker container by image
# start docker image 'docker start image_name'
# stop docker image 'docker stop image_name'
cassandrastart:
	docker start cassandra4

cassandrastop:
	docker stop cassandra4

# -----------------------------
# MySQL
# -----------------------------
# Pull MySQL from docker hub
mysqlpull:
	docker pull mysql:8.0.23

# run local mysql docker container
# use command 'docker ps images' to check postgres version
mysql:
	docker run --name=mysql8 -e MYSQL_ROOT_PASSWORD=keepsecret -e MYSQL_DATABASE=gogolang -p 3306:3306 -d mysql:8.0.23


# -----------------------------
# Postgres
# -----------------------------
postgrespull:
	docker pull postgres:latest

# run local postgress docker container
# use command 'docker ps images' to check postgres version
postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

# -----------------------------
# MongoDB
# -----------------------------
# This is temporary use less as configs still not complete. as we can use free tier at mongodb atlas service.
mongopull:
	docker pull mongo:4.4.4

# run local mongodb docker container
# use command 'docker ps images' to check postgres version
mangodb:
	# docker run --name postgres13 -p 27017:27017 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d mongo:4.4.4

.PHONY: start cassandrapull cassandra cassandraterminal cqlsh mysqlpull mysql postgrespull postgres 