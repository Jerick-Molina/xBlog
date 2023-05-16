package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	db  *mongo.Database
	rep *Repositories
}

func NewStorage(db *mongo.Database) *Storage {

	return &Storage{
		db: db,
		rep: &Repositories{
			*NewUser(db),
			*NewBlog(db),
			*NewSubcribe(db),
		},
	}
}
