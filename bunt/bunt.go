package bunt

import "github.com/tidwall/buntdb"

var conn *buntdb.DB

//Initialize crates a new BuntDB connection
func Initialize() error {
	var err error
	conn, err = buntdb.Open(":memory:")
	return err
}

func Get() *buntdb.DB {
	return conn
}
