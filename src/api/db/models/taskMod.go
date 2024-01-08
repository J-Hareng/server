package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskLink struct {
	TITLE string `json:"title,omitempty" bson:"title,omitempty"`
	URL   string `json:"url,omitempty" bson:"url,omitempty"`
}
type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	NAME        string             `json:"name" bson:"name"`
	DESCRIPTION string             `json:"des" bson:"des"`
	INPROGRESS  string             `json:"inprog,omitempty" bosn:"inprog,omitempty"`

	COLLECTION primitive.ObjectID `json:"coll" bson:"coll"`
	GROUPID    string             `json:"groupid,omitempty" bson:"groupid,omitempty"`
}

func CreateTask(name string, des string) Task {
	id := primitive.NewObjectID()

	return Task{
		ID:          id,
		NAME:        name,
		DESCRIPTION: des,
	}
}
