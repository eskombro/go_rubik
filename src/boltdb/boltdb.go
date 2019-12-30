package boltdb

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

/*
NORMAL USAGE OF WRAPPER (Example):

bolt.CreateDB("dbName")
bolt.CreateBucket("dbName", "bucketName")
bolt.Put("dbName", "bucketName", "SomeKey", "SomeValue")
bolt.View("dbName", "bucketName", "SomeKey")
bolt.Put("dbName", "bucketName", "SomeKey2", "SomeValue2")
bolt.Put("dbName", "bucketName", "SomeKey3", "SomeValue3")
bolt.PrintBucket("dbName", "bucketName")
bolt.Delete("dbName", "bucketName", "SomeKey2")
bolt.PrintBucket("dbName", "bucketName")
bolt.Flush("dbName", "bucketName")
bolt.PrintBucket("dbName", "bucketName")
bolt.CloseDB("dbName")
*/

var Bolt map[string]*BoltDB

type BoltDB struct {
	DB     *bolt.DB
	Bucket map[string]*BboltBucket
}

type BboltBucket struct {
	Bucket    *bolt.Bucket
	Name      string
	LastFlush time.Time
}

func handleError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func CreateDB(dbName string) *BoltDB {

	var err error
	opt := &bolt.Options{
		Timeout: 1 * time.Second,
	}

	if Bolt == nil {
		Bolt = make(map[string]*BoltDB, 10)
	}
	db := BoltDB{}
	db.DB, err = bolt.Open(dbName, 0600, opt)
	handleError(err)
	if db.DB != nil {
		Bolt[dbName] = &db
	}
	return &db
}

func CreateBucket(dbName string, bucketName string) {
	db := Bolt[dbName]
	if db.Bucket == nil {
		db.Bucket = make(map[string]*BboltBucket, 10)
	}
	err := db.DB.Update(func(tx *bolt.Tx) error {
		var err error
		db.Bucket[bucketName] = &BboltBucket{Name: bucketName}
		db.Bucket[bucketName].Bucket, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		db.Bucket[bucketName].LastFlush = time.Now()
		handleError(err)
		return nil
	})
	handleError(err)
}

func Put(dbName string, bucketName string, key string, value string) {
	db := Bolt[dbName]
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		handleError(err)
		err = b.Put([]byte(key), []byte(value))
		handleError(err)
		return nil
	})
	handleError(err)
}

func PutMap(dbName string, bucketName string, insertMap map[string]byte) {
	db := Bolt[dbName]
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		handleError(err)
		for k, v := range insertMap {
			err = b.Put([]byte(k), []byte{v})
		}
		handleError(err)
		return nil
	})
	handleError(err)
}

func View(dbName string, bucketName string, key string) {
	db := Bolt[dbName]
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucketName)
		}
		val := b.Get([]byte(key))
		fmt.Println(string(val))
		return nil
	})
	handleError(err)
}

func Get(dbName string, bucketName string, key string) string {
	str := "none"
	db := Bolt[dbName]
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucketName)
		}
		val := b.Get([]byte(key))
		if val != nil {
			str = string(val)
		}
		return nil
	})
	handleError(err)
	return str
}

func Delete(dbName string, bucketName string, key string) {
	db := Bolt[dbName]
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucketName)
		}
		err := b.Delete([]byte(key))
		handleError(err)
		return nil
	})
	handleError(err)
}

func PrintBucket(dbName string, bucketName string) {
	db := Bolt[dbName]
	fmt.Println("Cache bucket content:")
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		handleError(err)
		return nil
	})
	handleError(err)
}

func CountBucket(dbName string, bucketName string) int {
	ct := 0
	db := Bolt[dbName]
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.ForEach(func(k, v []byte) error {
			ct++
			return nil
		})
		handleError(err)
		return nil
	})
	handleError(err)
	return ct
}

func Flush(dbName string, bucketName string) {
	db := Bolt[dbName]
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.ForEach(func(k, v []byte) error {
			b.Delete(k)
			return nil
		})
		handleError(err)
		return nil
	})
	handleError(err)
}

func CloseDB(dbName string) {
	db := Bolt[dbName]
	db.DB.Close()
}
