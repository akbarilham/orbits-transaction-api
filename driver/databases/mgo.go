package database

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
)

var (
	MgoDatabase *mgo.Database
	MgoSession  *mgo.Session
)

func ConnectMgo() {
	const (
		hosts    = "180.250.19.191:27017"
		username = "orbits"
		password = "sigma2018*!"
	)
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Username: username,
		Password: password,
	}
	var err1 error
	log.Println("Connecting to mongodb...")
	MgoSession, err1 = mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}
	log.Println("Connected to mongodb!")
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
