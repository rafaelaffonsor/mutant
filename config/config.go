package config

const (
	DB_HOST = "172.0.0.1"
	DB_DRIVER = "mysql"
	DBUSER = "meli"
	DBPASSWORD = "meli"
	DBNAME = "meli"
	DBPORT = "3306"
	DB_STRING_CONNECTION = DBUSER + ":" + DBPASSWORD + "@tcp("+ DB_HOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8&parseTime=True"
)