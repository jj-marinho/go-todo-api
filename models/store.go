package models

import "gopkg.in/mgo.v2"

// Store is used mainly as a global pointer to our database
// It can be extended in the future to include other kinds of data
type Store struct {
	*mgo.Session
}

var S *Store = NewStore()

func NewStore() *Store {
	return &Store{}
}

func (S *Store) Connect(conn *mgo.Session) {
	S.Session = conn
}
