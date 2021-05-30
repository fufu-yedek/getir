package inmem

import "github.com/tidwall/buntdb"

var conn *buntdb.DB

func Initialize() error {
	var err error
	conn, err = buntdb.Open(":memory:")
	return err
}

func Get() *buntdb.DB {
	return conn
}
