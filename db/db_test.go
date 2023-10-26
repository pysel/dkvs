package db

import (
	"os"
	"testing"
)

func TestNewLevelDB(t *testing.T) {
	db, err := NewLevelDB("test")
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	if db == nil {
		t.Error("db is nil")
	}

	err = os.RemoveAll("test")
	if err != nil {
		t.Error(err)
	}
}
