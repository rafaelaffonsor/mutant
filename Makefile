TAG=latest
GOPATH=$(shell pwd)/vendor:$(shell pwd)

start:
	- docker-compose up -d

mysql-ip:
	- docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' meli_mysql