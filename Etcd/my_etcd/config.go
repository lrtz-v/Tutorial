package main

type EtcdConfig struct {
	Endpoints []string
	Username  string
	Password  string
	Prefix    string
}

var Endpoints = []string{"http://127.0.0.1:2379"}

const (
	Username = "user1"
	Password = "password"
	Appid    = "appid"
	Prefix   = "/appid"

	LogFileName = "my_etcd.log"
)

var Config = &EtcdConfig{Endpoints: Endpoints, Username: Username, Password: Password, Prefix: Prefix}


func (c *EtcdConfig) GetPrefix() string {
	return c.Prefix
}
