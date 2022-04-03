package mgutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

// IDField define
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

// UpdatedAtField defines
type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

// NewObjectID generates a new object id.
var NewObjectID = primitive.NewObjectID

// UpdatedAt returns a value suitable for UpdatedAt field
var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a $set update document
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert  returns a $set update document
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
