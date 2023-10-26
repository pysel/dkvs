package db

import "github.com/syndtr/goleveldb/leveldb"

type DB interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	Close() error
}

// LevelDB is a wrapper around GoLevelDB to implement the DB interface.
type LevelDB struct {
	*leveldb.DB
}

func (ldb *LevelDB) Get(key []byte) ([]byte, error) {
	return ldb.DB.Get(key, nil)
}

func (ldb *LevelDB) Put(key []byte, value []byte) error {
	return ldb.DB.Put(key, value, nil)
}

func (ldb *LevelDB) Delete(key []byte) error {
	return ldb.DB.Delete(key, nil)
}

// NewGoLevelDB returns a new instance of GoLevelDB.
func NewLevelDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
