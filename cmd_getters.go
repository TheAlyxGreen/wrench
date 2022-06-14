package wrench

import (
	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
)

func (w *Wrench) GetBucketContents(path paths.Path) ([]paths.Path, error) {

	const operationAction = "GetBucketContents" + actionSeparator
	var operationErr, otherErr, tmpErr error

	if !w.FilePathExists() {
		return nil, w.logError(path, operationAction+"FilePathExists", ErrorFileCannotBeFound)
	} else if w.ReadOnly {
		return nil, w.logError(path, operationAction+"IsReadOnly", ErrorWrenchCannotBeReadOnly)
	}

	db, otherErr := bbolt.Open(w.filePath, 0666, w.BoltOptions)
	if otherErr != nil {
		return nil, w.logError(path, operationAction+actionBoltOpen, otherErr)
	}

	tx, otherErr := db.Begin(true)
	if otherErr != nil {
		return nil, w.logError(path, operationAction+actionBoltBegin, otherErr)
	}

	bucketContents, operationErr := txGetBucketContents(tx, path)
	if operationErr != nil {
		operationErr = w.logError(path, operationAction+"txGetBucketContents", operationErr)
	}

	if operationErr != nil {
		tmpErr = tx.Rollback()
		if tmpErr != nil {
			otherErr = w.logError(path, operationAction+actionBoltRollback, tmpErr)
		}
	} else {
		tmpErr = tx.Commit()
		if tmpErr != nil {
			otherErr = w.logError(path, operationAction+actionBoltCommit, tmpErr)
		}
	}

	tmpErr = db.Close()
	if tmpErr != nil {
		otherErr = w.logError(path, operationAction+actionBoltClose, tmpErr)
	}

	if operationErr != nil {
		return nil, operationErr
	} else if otherErr != nil {
		return nil, otherErr
	} else {
		return bucketContents, nil
	}
}
