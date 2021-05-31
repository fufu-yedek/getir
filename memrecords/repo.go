package memrecords

import (
	"errors"
	"github.com/fufu-yedek/getir-challange/bunt"
	"github.com/fufu-yedek/getir-challange/gerrors"
	"github.com/tidwall/buntdb"
)

type Repository interface {
	CreateOrUpdate(record Record) (Record, error)
	FindOne(f Filter) (Record, error)
}

type inMemRepository struct {
	conn *buntdb.DB
}

func (r inMemRepository) CreateOrUpdate(record Record) (Record, error) {
	err := r.conn.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(record.Key, record.Value, nil)
		return err
	})

	if err != nil {
		return Record{}, err
	}

	return record, err
}

func (r inMemRepository) FindOne(f Filter) (Record, error) {
	var target string

	err := r.conn.View(func(tx *buntdb.Tx) (err error) {
		target, err = tx.Get(f.Key)
		return err
	})

	if err != nil {
		if errors.Is(err, buntdb.ErrNotFound) {
			err = gerrors.ErrRecordNotFound
		}

		return Record{}, err
	}

	return Record{
		Key:   f.Key,
		Value: target,
	}, nil
}

func NewInMemRepository(with *buntdb.DB) Repository {
	return inMemRepository{conn: with}
}

func NewDefaultInMemRepository() Repository {
	return NewInMemRepository(bunt.DB())
}
