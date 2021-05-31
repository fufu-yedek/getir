package records

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"_id"`
	Key       string             `bson:"key"`
	CreatedAt time.Time          `bson:"createdAt"`
	Counts    []int64            `bson:"counts"`
	Value     string             `bson:"value"`
}

type RecordWithCount struct {
	Key        string    `bson:"_id"`
	CreatedAt  time.Time `bson:"createdAt"`
	TotalCount int       `bson:"totalCount"`
}
