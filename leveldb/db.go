package db

import "github.com/syndtr/goleveldb/leveldb"

type DB interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	Has(key []byte) bool
	Close() error
}

// LevelDB is a wrapper around GoLevelDB to implement the DB interface.
type LevelDB struct {
	*leveldb.DB
}

func (ldb *LevelDB) Get(key []byte) ([]byte, error) {
	val, err := ldb.DB.Get(key, nil)
	if err != nil && err == leveldb.ErrNotFound {
		return val, nil
	}

	return val, err
}

func (ldb *LevelDB) Set(key []byte, value []byte) error {
	return ldb.DB.Put(key, value, nil)
}

func (ldb *LevelDB) Delete(key []byte) error {
	return ldb.DB.Delete(key, nil)
}

func (ldb *LevelDB) Has(key []byte) bool {
	has, _ := ldb.DB.Has(key, nil)
	return has
}

// NewGoLevelDB returns a new instance of GoLevelDB.
func NewLevelDB(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &LevelDB{db}, nil
}
