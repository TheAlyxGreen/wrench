package wrench

import (
	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
	"github.com/TheAlyxGreen/wrench/values"
)

func txGetBucket(tx *bbolt.Tx, path paths.Path) (*bbolt.Bucket, error) {
	var bucket *bbolt.Bucket
	if path.Length() == 0 {
		return nil, ErrorPathCannotBeRoot
	}
	bucket = tx.Bucket(path.Keys()[0].GetBytes())
	if path.Length() == 1 {
		return bucket, nil
	} else {
		for i := 1; i < path.Length(); i++ {
			bucket = tx.Bucket(path.Keys()[i].GetBytes())
			if bucket == nil {
				return nil, ErrorPathCannotBeFound
			}
		}
	}
	return bucket, nil
}

func txGetBucketContents(tx *bbolt.Tx, path paths.Path) ([]paths.Path, error) {
	bucketContents := make([]paths.Path, 0)
	if path.Length() == 0 {
		_ = tx.ForEach(
			func(key []byte, _ *bbolt.Bucket) error {
				contentKey := values.NewFromBytes(key)
				contentPath := paths.Append(path, contentKey)
				bucketContents = append(bucketContents, contentPath)
				return nil
			},
		)
	} else {
		bucket, err := txGetBucket(tx, path)
		if err != nil {
			return nil, err
		}
		_ = bucket.ForEach(
			func(key, _ []byte) error {
				contentKey := values.NewFromBytes(key)
				contentPath := paths.Append(path, contentKey)
				bucketContents = append(bucketContents, contentPath)
				return nil
			},
		)
	}
	return bucketContents, nil
}
