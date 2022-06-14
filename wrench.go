package wrench

import (
	"github.com/TheAlyxGreen/wrench/paths"
	"time"

	"go.etcd.io/bbolt"
)

// A Wrench is used to interact with a bolt database
type Wrench struct {
	filePath    string
	ReadOnly    bool
	errorLog    []ErrorContext
	BoltOptions *bbolt.Options
}

// New creates a new Wrench object from the specified File
func New(file string) (Wrench, error) {
	out := Wrench{
		filePath:    file,
		ReadOnly:    false,
		errorLog:    make([]ErrorContext, 5),
		BoltOptions: nil,
	}
	if !fileExists(file) {
		return out, ErrorFileCannotBeFound
	}
	return out, nil
}

// ErrMsg returns an ErrorContext with more information about the last error that occurred
func (w Wrench) ErrMsg() ErrorContext {
	return w.errorLog[len(w.errorLog)-1]
}

// ErrLog returns a list of all the errors encountered by the Wrench
func (w Wrench) ErrLog() []ErrorContext {
	return w.errorLog
}

// // newErrorCtx returns an ErrorContext with some fields autofilled
// func (w Wrench) newErrorCtx(p paths.Path, action string, err error) ErrorContext {
// 	ctx := ErrorContext{
// 		Timestamp: time.Now(),
// 		File:      w.filePath,
// 		Path:      p,
// 		Action:    action,
// 	}
// 	ctx.setErrors(err)
// 	return ctx
// }
//
// // logErrorCtx appends an ErrorContext to the Wrench's error log
// func (w Wrench) logErrorCtx(ctx ErrorContext) {
// 	ctx.Timestamp = time.Now()
// 	w.errorLog = append(w.errorLog, ctx)
// }

// logError creates a new ErrorContext and appends it to the Wrench's error log
func (w *Wrench) logError(p paths.Path, action string, err error) error {
	ctx := ErrorContext{
		Timestamp: time.Now(),
		File:      w.filePath,
		Path:      p,
		Action:    action,
	}
	ctx.setErrors(err)
	w.errorLog = append(w.errorLog, ctx)
	return ctx.WrenchErr
}
