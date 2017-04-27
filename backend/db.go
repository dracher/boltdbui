package backend

import (
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

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fmt.Printf("key=%s, value=%s\n", k, v)
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
