package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhanglp92/rep/config"
)

// desc: leveldb存储

var gDB *leveldb.DB

// SetDef ...
func SetDef(db *leveldb.DB) {
	gDB = db
}

// DB ...
func DB() *leveldb.DB {
	return gDB
}

// LevelDB 使用leveldb缓存
type LevelDB struct {
	config *config.Config

	db *leveldb.DB
}

// New ...
func New(config *config.Config) (a *LevelDB, err error) {
	a = &LevelDB{config: config}

	return a, a.init()
}

// Op ...
func (a *LevelDB) Op() *leveldb.DB {
	return a.db
}

// Close ...
func (a *LevelDB) Close() error {
	return a.db.Close()
}

// ---- internal ----

func (a *LevelDB) init() (err error) {
	if a.db, err = leveldb.OpenFile(a.config.LevelDBPath(), nil); err != nil {
		return
	}
	return
}
