package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	NAME        string             `json:"name" bson:"name"`
	DESCRIPTION string             `json:"avalability" bson:"avalability"`
	USERS       []UserLink         `json:"users,omitempty" bson:"users,omitempty"`
	TASKS       []TaskLink         `json:"tasks,omitempty" bson:"tasks,omitempty"`
	DONETASKS   []TaskLink         `json:"donetasks,omitempty" bson:"donetasks,omitempty"`
	POSTS       []PostLink         `json:"posts,omitempty" bson:"posts,omitempty"`
	GROUPID     string             `json:"groupid,omitempty" bson:"groupid,omitempty"`
}
type TeamLink struct {
	ID string `json:"_id" bson:"_id"`
}
