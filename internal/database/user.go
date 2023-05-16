package database

import "go.mongodb.org/mongo-driver/mongo"

type UserCollection struct {
	DBTX
}

func NewUser(db *mongo.Database) *UserCollection {
	return &UserCollection{db.Collection("User")}
}

func (usr *UserCollection) Test() {

}
