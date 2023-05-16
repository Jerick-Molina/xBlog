package database

import "go.mongodb.org/mongo-driver/mongo"

type SubcriberCollection struct {
	DBTX
}

func NewSubcribe(db *mongo.Database) *SubcriberCollection {
	return &SubcriberCollection{db.Collection("User")}
}

const followUser = `Create `

func FollowUser() error {

	return nil
}

const unfollowUser = ``

func (q *Queries) UnfollowUser() error {

	return nil
}
