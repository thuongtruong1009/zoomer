package db

import (
	mgo "gopkg.in/mgo.v2"
	config "zoomer/configs"
)

var instance *mgo.Session

var err error

func GetMongoInstance(c *config.Configuration) *mgo.Session {
	if instance == nil {
		instance, err = mgo.Dial(c.DatabaseConnectionURL)
		if err != nil {
			panic(err)
		}
	}
	return instance.Copy()
}
