package records

import (
	"context"
	"github.com/fufuceng/getir-challange/gerrors"
	internalMongo "github.com/fufuceng/getir-challange/mongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const mongoCollectionName = "records"

type Repository interface {
	FindWithCount(Filter) ([]RecordWithCount, error)
}

type mongoRepository struct {
	conn *mongo.Database
}

func (r mongoRepository) collection() *mongo.Collection {
	return r.conn.Collection(mongoCollectionName)
}

func (r mongoRepository) FindWithCount(f Filter) ([]RecordWithCount, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"_id":       0, // exclude
				"key":       1,
				"createdAt": 1,
				"totalCount": bson.M{
					"$sum": "$counts",
				},
			},
		},
		{
			"$group": bson.M{
				"_id": "$key",
				"createdAt": bson.M{
					"$first": "$createdAt",
				},
				"totalCount": bson.M{
					"$sum": "$totalCount",
				},
			},
		},
	}

	if match := GenerateMongoQuery(f); match != nil {
		pipeline = append(pipeline, bson.M{"$match": match})
	}

	cur, err := r.collection().Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			logrus.WithField("location", "mongoRepository - Find()").
				WithError(err).Error("error while closing cursor")
		}
	}()

	var items []RecordWithCount
	if err := cur.All(context.TODO(), &items); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, gerrors.ErrRecordNotFound
	}

	return items, nil
}

func NewMongoRepository(withDB *mongo.Database) Repository {
	return mongoRepository{conn: withDB}
}

func NewDefaultMongoRepository() Repository {
	return NewMongoRepository(internalMongo.Get())
}
