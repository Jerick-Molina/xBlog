package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It defines the methods needed to interact with the database
type DBTX interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
}

type Repositories struct {
	UserCollection
	BlogCollection
	SubcriberCollection
}
type Queries struct {
	DBTX
}
