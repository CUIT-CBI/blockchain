package kvstore

import (
	"merkletree/hash"

	"github.com/syndtr/goleveldb/leveldb"
)

type KVStore interface {
	Has(key []byte) (bool, error)
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
}

type LevelDB struct {
	db *leveldb.DB
}

func NewLevelDB(path string) *LevelDB {
	db, _ := leveldb.OpenFile(path, nil)
	return &LevelDB{
		db: db,
	}
}

func (db *LevelDB) Close() error {
	return db.db.Close()
}

func (db *LevelDB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

func (db *LevelDB) Put(value []byte) ([]byte, error) {
	h := hash.Sha3Sum256(value)
	err := db.db.Put(h, value, nil)
	return h, err
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	return db.db.Get(key, nil)
}

func (db *LevelDB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}
