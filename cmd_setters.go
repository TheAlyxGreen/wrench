package wrench

import (
	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
	"github.com/TheAlyxGreen/wrench/values"
)

func (w *Wrench) CreateBucket(path paths.Path) error {

	const operationAction = "CreateBucket" + actionSeparator
	var operationErr, otherErr, tmpErr error

	if !w.FilePathExists() {
		return w.logError(path, operationAction+"FilePathExists", ErrorFileCannotBeFound)
	} else if w.ReadOnly {
		return w.logError(path, operationAction+"IsReadOnly", ErrorWrenchCannotBeReadOnly)
	} else if path.Equals(paths.RootPath) {
		return w.logError(path, operationAction+"IsRoot", ErrorPathCannotBeRoot)
	}

	db, otherErr := bbolt.Open(w.filePath, 0666, w.BoltOptions)
	if otherErr != nil {
		return w.logError(path, operationAction+actionBoltOpen, otherErr)
	}

	tx, otherErr := db.Begin(true)
	if otherErr != nil {
		return w.logError(path, operationAction+actionBoltBegin, otherErr)
	}

	operationErr = txCreateBucket(tx, path)
	if operationErr != nil {
		operationErr = w.logError(path, operationAction+"txCreateBucket", operationErr)
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
		return operationErr
	} else {
		return otherErr
	}
}

func (w *Wrench) SetValue(path paths.Path, value values.Value) error {

	const operationAction = "SetValue" + actionSeparator
	var operationErr, otherErr, tmpErr error

	if !w.FilePathExists() {
		return w.logError(path, operationAction+"FilePathExists", ErrorFileCannotBeFound)
	} else if w.ReadOnly {
		return w.logError(path, operationAction+"ReadOnly", ErrorWrenchCannotBeReadOnly)
	}

	db, otherErr := bbolt.Open(w.filePath, 0666, w.BoltOptions)
	if otherErr != nil {
		return w.logError(path, operationAction+actionBoltOpen, otherErr)
	}

	tx, otherErr := db.Begin(true)
	if otherErr != nil {
		return w.logError(path, operationAction+actionBoltBegin, otherErr)
	}

	operationErr = txSetValue(tx, path, value)
	if operationErr != nil {
		operationErr = w.logError(path, operationAction+"txSetValue", operationErr)
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
		return operationErr
	} else {
		return otherErr
	}
}
