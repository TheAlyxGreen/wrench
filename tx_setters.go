package wrench

import (
	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
	"github.com/TheAlyxGreen/wrench/values"
)

// txCreateBucket creates a bucket unless the specified path already exists
func txCreateBucket(tx *bbolt.Tx, path paths.Path) error {
	// if it's root, throw an error
	if path.Length() == 0 {
		return ErrorPathCannotBeRoot
	}

	// grab the key of the new bucket
	key, _ := path.PeekKey()

	// if the bucket will be in the root directory, create it
	if path.Length() == 1 {
		_, err := tx.CreateBucket(key.GetBytes())
		return err
	}

	// otherwise grab the parent bucket that will hold the new bucket
	parentBucket, err := txGetParentBucket(tx, path)
	if err != nil {
		return err
	}

	// if the new bucket path does not exist
	if parentBucket.Bucket(key.GetBytes()) == nil {
		// create thew new bucket inside the parent bucket
		_, err := parentBucket.CreateBucket(key.GetBytes())
		return err
	} else {
		// throw an error if the bucket path already exists
		return ErrorPathCannotAlreadyExist
	}

}

// txSetValue sets the value if it already exists and throws an error if it does not
func txSetValue(tx *bbolt.Tx, path paths.Path, value values.Value) error {
	// if it's root, throw an error
	if path.Length() == 0 {
		return ErrorPathCannotBeRoot
	} else if path.Length() == 1 {
		return ErrorRootCannotContainValues
	}

	// get the parent bucket that hold the value
	parentBucket, err := txGetParentBucket(tx, path)
	if err != nil {
		return err
	}

	// grab the key of the new bucket
	key, _ := path.PeekKey()

	// verify that the key is not a bucket
	if parentBucket.Bucket(key.GetBytes()) != nil {
		return ErrorPathCannotBeBucket
	}

	// verify that the key exists
	if parentBucket.Get(key.GetBytes()) == nil {
		return ErrorPathCannotBeFound
	}

	// write the value to the path
	err = parentBucket.Put(key.GetBytes(), value.GetBytes())

	return err
}
