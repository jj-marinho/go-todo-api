package config

import "gopkg.in/mgo.v2"

// TODO: This should be in a .env
var MONGO_URI string = "mongodb://localhost:27017/SampleDB"

func StartDBConn(url string) (*mgo.Session, error) {
	conn, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
