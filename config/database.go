package config

import "fmt"

const DBDialect = "postgres"
const DBHost = "localhost"
const DBPort = "5432"
const DBUser = "postgres"
const DBName = "ecommerce"
const DBPassword = "darkside"

var DBURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
