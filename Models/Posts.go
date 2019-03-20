package Models

import "gopkg.in/mgo.v2/bson"

type (
	Post struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UID        string        `json:"uid" bson:"uid"`
		Name      string        `json:"name" bson:"name"`
		Title    string        `json:"title" bson:"title"`
		Content string        `json:"content" bson:"content"`
	}
)
