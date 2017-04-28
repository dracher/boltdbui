package backend

import (
	"bytes"
	"fmt"

	"github.com/boltdb/bolt"
)

// DB is connection to database instance
var DB *bolt.DB

func dbStatus() bolt.Stats {
	return DB.Stats()
}

func listBucketByName(bucketName string) [][2]string {
	res := [][2]string{}
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		count := 0
		for k, v := c.Last(); k != nil && count < 50; k, v = c.Prev() {
			count++
			res = append(res, [2]string{string(k), string(v)})
		}
		return nil
	})
	return res
}

func listAllBuckets() []string {
	res := []string{}

	DB.View(func(tx *bolt.Tx) error {

		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			b := []string{string(name)}
			res = append(res, b...)
			return nil
		})
	})
	return res
}

func createBucket(bucketName string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
	return err
}

func deleteBucket(bucketName string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
	return err
}

func getValFromBucket(key, bucketName string) ([]byte, error) {
	var res []byte
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		res = b.Get([]byte(key))
		return nil
	})
	return res, nil
}

func putValToBucket(key, val, bucketName string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Put([]byte(key), []byte(val))
	})
}

func prefixScan(bucketName, prefix string) ([][2]string, error) {
	res := [][2]string{}
	DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucketName)).Cursor()

		prefix := []byte(prefix)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			res = append(res, [2]string{string(k), string(v)})
		}
		return nil
	})
	return res, nil
}

func deleteKeyFromBucket(key, bucketName string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(key))
	})
	return err
}
