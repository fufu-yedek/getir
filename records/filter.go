package records

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//Filter stores filterable fields of the record
type Filter struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  uint
	MaxCount  uint
}

func (f Filter) Validate() error {
	if f.MinCount > 0 && f.MaxCount > 0 {
		if f.MinCount > f.MaxCount {
			return fmt.Errorf("min_count must be less than or equeal to the max_count")
		}
	}

	if !f.StartDate.IsZero() && !f.EndDate.IsZero() {
		if f.StartDate.After(f.EndDate) {
			return fmt.Errorf("start_date must be less than or equeal to the end_date")
		}
	}

	return nil
}

//GenerateMongoQuery generates a mongo db query from given filters
func GenerateMongoQuery(from Filter) bson.M {
	var q []bson.M

	if !from.StartDate.IsZero() {
		q = append(q, bson.M{"createdAt": bson.M{"$gte": from.StartDate}})
	}

	if !from.EndDate.IsZero() {
		q = append(q, bson.M{"createdAt": bson.M{"$lte": from.EndDate}})
	}

	if from.MinCount > 0 {
		q = append(q, bson.M{"totalCount": bson.M{"$gte": from.MinCount}})
	}

	if from.MaxCount > 0 {
		q = append(q, bson.M{"totalCount": bson.M{"$lte": from.MaxCount}})
	}

	if len(q) > 0 {
		return bson.M{"$and": q}
	}

	return nil
}
