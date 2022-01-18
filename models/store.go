package models

import "gopkg.in/mgo.v2"

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
