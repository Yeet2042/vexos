package xerror

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type AppError struct {
	Code    int
	Message error
	Source  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s]: %v", e.Source, e.Message)
}

func NewError(err error, statusCode int) error {
	_, file, line, ok := runtime.Caller(1)

	source := "unknown"
	if ok {
		cwd, _ := os.Getwd()
		if rel, err := filepath.Rel(cwd, file); err == nil {
			source = fmt.Sprintf("%s:%d", rel, line)
		} else {
			source = fmt.Sprintf("%s:%d", file, line)
		}
	}

	return &AppError{
		Code:    statusCode,
		Message: err,
		Source:  source,
	}
}
