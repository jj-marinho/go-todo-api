package config

import "gopkg.in/mgo.v2"

// TODO: This should be in a .env
var MONGO_URI string = "mongodb://localhost:27017/SampleDB"

var Mongo *mgo.Session = nil

func StartMongoDB(url string) error {
	conn, err := mgo.Dial(url)
	if err != nil {
		return err
	}

	Mongo = conn

	return nil
}
