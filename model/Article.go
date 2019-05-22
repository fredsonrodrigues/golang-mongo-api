package model

import "gopkg.in/mgo.v2/bson"

type Article struct {
	ID          bson.ObjectId 	`bson:"_id" json:"id"`
	Title 		string 			`bson:"title" json:"title"`
	Subtitle 	string 			`bson:"subtitle" json:"subtitle"`
	Slug 		string 			`bson:"slug" json:"slug"`
	Content 	string 			`bson:"content" json:"content"`
	Comments 	[]Comments 		`bson:"comments" json:"comments"`
}

type Comments struct {
	Author 		string 			`bson:"author" json:"author"`
	Timestamp 	string 			`bson:"timestamp" json:"timestamp"`
	Email 		string 			`bson:"email" json:"email"`
	Content 	string 			`bson:"content" json:"content"`
	Photo 		string 			`bson:"photo" json:"photo"`
}