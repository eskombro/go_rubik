package boltdb

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var Bolt BoltDB

type BoltDB struct {
	DB     *bolt.DB
	Bucket *BboltBucket
}

type BboltBucket struct {
	Bucket    *bolt.Bucket
	Name      string
	LastFlush time.Time
}

// NORMAL USAGE:
// bolt.CreateDB()
// bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
// bolt.CreateBucket(bolt.Bolt.Bucket)
// fmt.Println(bolt.Bolt.Bucket)
// bolt.Put(bolt.Bolt.Bucket, "SomeKey", "SomeValue")
// bolt.View(bolt.Bolt.Bucket, "SomeKey")
// bolt.Put(bolt.Bolt.Bucket, "SomeKey2", "SomeValue2")
// bolt.Put(bolt.Bolt.Bucket, "SomeKey3", "SomeValue3")
// bolt.PrintBucket(bolt.Bolt.Bucket)
// bolt.Delete(bolt.Bolt.Bucket, "SomeKey2")
// bolt.PrintBucket(bolt.Bolt.Bucket)
// bolt.Flush(bolt.Bolt.Bucket)
// bolt.PrintBucket(bolt.Bolt.Bucket)
// bolt.CloseDB()

func handleError(err error) {
	if err != nil {
		log.Print(err)
	}
}

func CreateDB() {
	var err error
	opt := &bolt.Options{
		Timeout: 1 * time.Second,
	}
	Bolt.DB, err = bolt.Open("cache/Cache.bolt", 0600, opt)
	handleError(err)
}

func CreateBucket(bucket *BboltBucket) {
	err := Bolt.DB.Update(func(tx *bolt.Tx) error {
		var err error
		bucket.Bucket, err = tx.CreateBucketIfNotExists([]byte(bucket.Name))
		bucket.LastFlush = time.Now()
		handleError(err)
		return nil
	})
	handleError(err)
}

func Put(bucket *BboltBucket, key string, value string) {
	err := Bolt.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket.Name))
		handleError(err)
		err = b.Put([]byte(key), []byte(value))
		handleError(err)
		return nil
	})
	handleError(err)
}

func View(bucket *BboltBucket, key string) {
	err := Bolt.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket.Name))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucket.Name)
		}
		val := b.Get([]byte(key))
		fmt.Println(string(val))
		return nil
	})
	handleError(err)
}

func Get(bucket *BboltBucket, key string) string {
	str := "none"
	err := Bolt.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket.Name))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucket.Name)
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

func Delete(bucket *BboltBucket, key string) {
	err := Bolt.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket.Name))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucket.Name)
		}
		err := b.Delete([]byte(key))
		handleError(err)
		return nil
	})
	handleError(err)
}

func PrintBucket(bucket *BboltBucket) {
	fmt.Println("Cacche bucket content:")
	err := Bolt.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket.Name))
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		handleError(err)
		return nil
	})
	handleError(err)
}

func Flush(bucket *BboltBucket) {
	err := Bolt.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket.Name))
		err := b.ForEach(func(k, v []byte) error {
			b.Delete(k)
			return nil
		})
		handleError(err)
		return nil
	})
	handleError(err)
}

func CloseDB() {
	Bolt.DB.Close()
}
