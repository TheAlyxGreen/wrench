package wrench

import (
	"os"

	"go.etcd.io/bbolt"
)

const testDatabasePathString = "./testing/test.db"
const testValuePathString = "/testBucket/testKey"

func init() {
	err := resetTestDatabase()
	if err != nil {
		panic(err)
	}
}

func resetTestDatabase() error {
	_ = os.Remove(testDatabasePathString)
	db, err := bbolt.Open(testDatabasePathString, 0666, nil)
	if err != nil {
		return err
	}
	err = db.Update(
		func(tx *bbolt.Tx) error {
			b, err := tx.CreateBucket([]byte("this"))
			if err != nil {
				return err
			}
			b, err = b.CreateBucket([]byte("is"))
			if err != nil {
				return err
			}
			b, err = b.CreateBucket([]byte("a"))
			if err != nil {
				return err
			}
			b, err = b.CreateBucket([]byte("test"))
			if err != nil {
				return err
			}
			b, err = tx.CreateBucket([]byte("testBucket"))
			if err != nil {
				return err
			}
			err = b.Put([]byte("testKey"), []byte("testValue"))
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func indent(amount int) string {
	if amount < 1 {
		return ""
	} else if amount == 1 {
		return "\t| "
	} else {
		out := "\t|"
		for i := 1; i < amount; i++ {
			out = out + "\t"
		}
		out = out + "> "
		return out
	}
}
