package models

import (
	"gopkg.in/mgo.v2/bson"
)

var todoC = "todos"

type Todo struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Done        bool          `json:"done"`
}

func GetAllTodos() ([]Todo, error) {
	sess := S.Session.Copy()
	defer sess.Close()

	todoCollection := sess.DB("").C(todoC)
	query := todoCollection.Find(nil)

	var todos []Todo
	err := query.All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodo(todo *Todo) error {
	todo.Id = bson.NewObjectId()

	sess := S.Session.Copy()
	defer sess.Close()

	todoCollection := sess.DB("").C(todoC)
	err := todoCollection.Insert(todo)
	if err != nil {
		return err
	}

	return nil
}

func GetTodo(id string) (*Todo, error) {
	sess := S.Session.Copy()
	defer sess.Close()

	var todo *Todo

	todoCollection := sess.DB("").C(todoC)
	err := todoCollection.FindId(bson.ObjectIdHex(id)).One(&todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func UpdateTodo(id string, new_todo *Todo) error {
	sess := S.Session.Copy()
	defer sess.Close()

	todoCollection := sess.DB("").C(todoC)
	err := todoCollection.UpdateId(bson.ObjectIdHex(id), new_todo)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id string) error {
	sess := S.Session.Copy()
	defer sess.Close()

	todoCollection := sess.DB("").C(todoC)
	err := todoCollection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil
}
