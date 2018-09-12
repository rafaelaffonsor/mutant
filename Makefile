TAG=latest
GOPATH=$(shell pwd)/vendor:$(shell pwd)

start:
	- docker network create --gateway 172.25.0.1 --subnet 172.25.0.0/24 mutant-network
	- docker-compose up -d

mysql-ip:
	- docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mutant_mysql