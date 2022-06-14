package wrench

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"go.etcd.io/bbolt"

	"github.com/TheAlyxGreen/wrench/paths"
)

var (
	ErrorFileCannotBeFound   = errors.New("cannot find file")
	ErrorFileCannotBeOpened  = errors.New("cannot open file")
	ErrorFileCannotBeWritten = errors.New("cannot write to file")

	ErrorWrenchCannotBeReadOnly = errors.New("cannot write while wrench is set to ReadOnly")

	ErrorPathCannotBeRoot         = errors.New("cannot perform action on root path")
	ErrorPathCannotBeFound        = errors.New("cannot find path in database")
	ErrorPathCannotAlreadyExist   = errors.New("cannot perform action when path already exists")
	ErrorPathCannotBeIncompatible = errors.New("cannot perform action on specified path")
	ErrorPathCannotBeValue        = errors.New("cannot perform action on a value (must be bucket)")
	ErrorPathCannotBeBucket       = errors.New("cannot perform action on a bucket (must be value)")

	ErrorKeyCannotBeLong  = errors.New("cannot use keys longer than " + strconv.Itoa(bbolt.MaxKeySize))
	ErrorKeyCannotBeBlank = errors.New("cannot use an blank value as a key")

	ErrorRootCannotContainValues = errors.New("cannot write values in root bucket")

	ErrorBucketCannotBeFilled = errors.New("cannot perform action on filled bucket")
	ErrorBucketCannotBeEmpty  = errors.New("cannot perform action on empty bucket")

	ErrorUnknownError = errors.New("an unknown error occurred; check ErrMsg() for more details")
)

const (
	actionSeparator    = "->"
	actionBoltOpen     = "Bolt:Open"
	actionBoltClose    = "Bolt:Close"
	actionBoltBegin    = "Bolt:Begin"
	actionBoltCommit   = "Bolt:Commit"
	actionBoltRollback = "Bolt:Rollback"
)

type ErrorContext struct {
	Timestamp   time.Time
	File        string
	Path        paths.Path
	Action      string
	InternalErr error
	WrenchErr   error
}

func (ctx ErrorContext) String() string {
	return fmt.Sprintf(
		"Time: %s\n"+
			"File: %s\n"+
			"Path: %s\n"+
			"Action: %s\n"+
			"Internal Error Message: %v\n"+
			"Wrench Error Message: %v",
		ctx.Timestamp.String(),
		ctx.File,
		ctx.Path,
		ctx.Action,
		ctx.InternalErr,
		ctx.WrenchErr,
	)
}

// HasError returns true if either error value is not nil
func (ctx ErrorContext) HasError() bool {
	return ctx.InternalErr != nil || ctx.WrenchErr != nil
}

func (ctx *ErrorContext) setErrors(err error) {
	ctx.InternalErr = nil
	ctx.WrenchErr = err
	switch err {
	case bbolt.ErrDatabaseOpen:
	case bbolt.ErrTimeout:
	case bbolt.ErrInvalid:
	case bbolt.ErrVersionMismatch:
	case bbolt.ErrChecksum:
		ctx.WrenchErr = ErrorFileCannotBeOpened
		break
	case bbolt.ErrDatabaseNotOpen:
	case bbolt.ErrTxNotWritable:
	case bbolt.ErrTxClosed:
		ctx.WrenchErr = ErrorFileCannotBeWritten
		break
	case bbolt.ErrDatabaseReadOnly:
		ctx.WrenchErr = ErrorWrenchCannotBeReadOnly
		break
	case bbolt.ErrBucketNotFound:
		ctx.WrenchErr = ErrorPathCannotBeFound
		break
	case bbolt.ErrBucketExists:
		ctx.WrenchErr = ErrorPathCannotAlreadyExist
		break
	case bbolt.ErrKeyRequired:
	case bbolt.ErrBucketNameRequired:
		ctx.WrenchErr = ErrorKeyCannotBeBlank
		break
	case bbolt.ErrKeyTooLarge:
		ctx.WrenchErr = ErrorKeyCannotBeLong
		break
	case bbolt.ErrIncompatibleValue:
		ctx.WrenchErr = ErrorPathCannotBeIncompatible
		break
	}
}
