package config

const (
	DB_HOST = "172.22.0.2"
	DB_DRIVER = "mysql"
	DBUSER = "mutant"
	DBPASSWORD = "mutant"
	DBNAME = "mutant"
	DBPORT = "3306"
	DB_STRING_CONNECTION = DBUSER + ":" + DBPASSWORD + "@tcp("+ DB_HOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8&parseTime=True"
)