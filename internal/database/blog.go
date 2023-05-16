package database

import "go.mongodb.org/mongo-driver/mongo"

type BlogCollection struct {
	DBTX
}

func NewBlog(db *mongo.Database) *BlogCollection {
	return &BlogCollection{db.Collection("User")}
}
func Postblog() error {
	return nil
}

func Readblog() error {
	return nil
}

func Editblog() error {
	return nil
}

func Deleteblog() error {
	return nil
}
