package r

import (
	"errors"
	"fmt"
	"net/http"
)

type Ex interface {
	IsDatabaseError() bool
	IsRepositoryError() bool
	IsServiceError() bool
	IsDataNotFound() bool
	Error() string
}

type errs struct {
	isDatabaseError   bool
	isRepositoryError bool
	isServiceError    bool
	isDataNotFound    bool
	Message           string
}

func (e *errs) IsDatabaseError() bool {
	return e.isDatabaseError
}

func (e *errs) IsRepositoryError() bool {
	return e.isRepositoryError
}

func (e *errs) IsServiceError() bool {
	return e.isServiceError
}

func (e *errs) IsDataNotFound() bool {
	return e.isDataNotFound
}

func newDatabaseError(message string) errs {
	return errs{
		isDatabaseError:   true,
		isRepositoryError: false,
		isServiceError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newRepositoryError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: true,
		isServiceError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newServiceError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isServiceError:    true,
		isDataNotFound:    false,
		Message:           message,
	}
}

func newNFError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isServiceError:    false,
		isDataNotFound:    true,
		Message:           message,
	}
}

func newOError(message string) errs {
	return errs{
		isDatabaseError:   false,
		isRepositoryError: false,
		isServiceError:    false,
		isDataNotFound:    false,
		Message:           message,
	}
}

func NewErrorDataNotFound(coll string, err error) Ex {
	return &ErrorRepository{
		errs:       newNFError(err.Error()),
		Collection: coll,
	}
}

func NewErrorRepository(coll string, err error) Ex {
	return &ErrorRepository{
		errs:       newRepositoryError(err.Error()),
		Collection: coll,
	}
}

type ErrorRepository struct {
	errs
	Collection string
}

func (e *ErrorRepository) Error() string {
	return fmt.Sprintf("error repository (%s) : %s", e.Collection, e.Message)
}

func NewErrorService(err error) Ex {
	return &ErrorService{newServiceError(err.Error())}
}

type ErrorService struct {
	errs
}

func (e *ErrorService) Error() string {
	return fmt.Sprintf("error service : %s", e.Message)
}

func NewErrorDatabase(err error) Ex {
	return &ErrorDatabase{newDatabaseError(err.Error())}
}

type ErrorDatabase struct {
	errs
}

func (e *ErrorDatabase) Error() string {
	return fmt.Sprintf("error database : %s", e.Message)
}

func NewBodyParseError(err error) Ex {
	return &ErrorResult{
		errs:   newOError(err.Error()),
		Status: http.StatusUnprocessableEntity,
	}
}

func NewErr(err string) Ex {
	return NewErrorService(errors.New(err))
}

type ErrorResult struct {
	errs
	Status int32
}

func (e *ErrorResult) Error() string {
	return fmt.Sprintf("error (%d) : %s", e.Status, e.Message)
}
