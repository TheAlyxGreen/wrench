package wrench

import (
	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
)

func txGetParentBucket(tx *bbolt.Tx, path paths.Path) (*bbolt.Bucket, error) {

	if path.Length() <= 1 {
		return nil, ErrorPathCannotBeRoot
	}
	_, err := path.PopKey()
	if err != nil {
		return nil, err
	}
	bucket, err := txGetBucket(tx, path)

	return bucket, err

}
