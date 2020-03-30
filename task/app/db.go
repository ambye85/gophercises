package todo

import (
	"errors"
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "tasks.db"
const bucketName = "tasks"

func Add(task string) error {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
		}
		err = b.Put([]byte(task), []byte("false"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	return err
}

func Do(task int) (string, error) {
	tasks, err := List()
	if err != nil {
		return "", err
	}

	if task-1 >= len(tasks) {
		return "", errors.New("could not load tasks")
	}

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
		}
		err = bucket.Put([]byte(tasks[task-1]), []byte("true"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	return tasks[task-1], nil
}

func List() ([]string, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var tasks []string
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			log.Fatal(err)
			return err
		}
		err = bucket.ForEach(func(k, v []byte) error {
			if string(v) != "true" {
				tasks = append(tasks, string(k))
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	return tasks, err
}
