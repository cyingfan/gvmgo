package cache

import (
	"os"
	"path"
	"time"

	"github.com/cyingfan/gvmgo/api"
	"github.com/syndtr/goleveldb/leveldb"
)

var Db *leveldb.DB = nil

func GetDb() *leveldb.DB {
	if Db == nil {
		apppath, found := os.LookupEnv("APPDATA")
		if !found {
			apppath = "."
		}
		dir := path.Join(apppath, "gvmgo")
		Db, _ = leveldb.OpenFile(dir, nil)
	}
	return Db
}

func GetLastUpdate() (time.Time, error) {
	db := GetDb()
	item, err := db.Get([]byte("last-update"), nil)
	if err != nil {
		return time.Time{}, err
	}
	t, err := time.Parse(time.RFC1123, string(item))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func GetLastUpdateISO() string {
	t, _ := GetLastUpdate()
	return t.Format(time.RFC1123)
}

func SetLastUpdate(t time.Time) {
	db := GetDb()
	db.Put([]byte("last-update"), []byte(t.Format(time.RFC1123)), nil)
}

func Update() error {
	candidates := api.GetCandidatesList()
	broadcast := api.GetBroadcast()

	db := GetDb()
	db.Put([]byte("candidates"), candidates, nil)
	db.Put([]byte("broadcast"), broadcast, nil)
	return nil
}

func GetCandidates() string {
	db := GetDb()
	candidates, _ := db.Get([]byte("candidates"), nil)
	return string(candidates)
}

func GetBroadcast() string {
	db := GetDb()
	broadcast, _ := db.Get([]byte("broadcast"), nil)
	return string(broadcast)
}
