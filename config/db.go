package config

import "gopkg.in/mgo.v2"

var MONGO_URI string = "mongodb://localhost:27017/SampleDB"

// Wrapper for mgo.dial()
func StartDBConn(url string) (*mgo.Session, error) {
	conn, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
