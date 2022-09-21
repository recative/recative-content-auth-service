package db_err

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type DBErrType struct {
	Code int
	Name string
}

func (D DBErrType) wrap(err error) error {
	return &DBErr{
		Code:    D.Code,
		Name:    D.Name,
		Message: err.Error(),
		Payload: err,
	}
}

func (D DBErrType) Is(err error) bool {
	e, ok := err.(*DBErr)
	if !ok {
		return false
	}
	return e.Code == D.Code
}

var (
	NotFound = DBErrType{
		Code: 404200,
		Name: "db_not_found",
	}

	IntegrityConstraintViolation = DBErrType{
		Code: 500200,
		Name: "db_integrity_constraint_violation",
	}

	RestrictViolation = DBErrType{
		Code: 500201,
		Name: "db_restrict_violation",
	}

	NotNullViolation = DBErrType{
		Code: 500202,
		Name: "db_not_null_violation",
	}

	ForeignKeyViolation = DBErrType{
		Code: 500203,
		Name: "db_foreign_key_violation",
	}

	UniqueViolation = DBErrType{
		Code: 500204,
		Name: "db_unique_violation",
	}

	CheckViolation = DBErrType{
		Code: 500205,
		Name: "db_check_violation",
	}

	ExclusionViolation = DBErrType{
		Code: 500206,
		Name: "db_exclusion_violation",
	}

	OperatorIntervention = DBErrType{
		Code: 500207,
		Name: "db_operator_intervention",
	}

	TransactionIntegrityConstraintViolation = DBErrType{
		Code: 500208,
		Name: "db_transaction_integrity_constraint_violation",
	}

	TransactionRollback = DBErrType{
		Code: 503200,
		Name: "db_transaction_rollback",
	}

	SerializationFailure = DBErrType{
		Code: 503201,
		Name: "db_serialization_failure",
	}

	DeadLockDetected = DBErrType{
		Code: 503202,
		Name: "db_deadlock_detected",
	}

	QueryCanceled = DBErrType{
		Code: 503203,
		Name: "db_query_canceled",
	}

	Conflict = DBErrType{
		Code: 409200,
		Name: "db_conflict",
	}
)

type DBErr struct {
	Code    int
	Name    string
	Message string
	Payload any
}

func (d *DBErr) Error() string {
	return fmt.Sprintf("%d: %s: %s: %v", d.Code, d.Name, d.Message, d.Payload)
}

var _ error = (*DBErr)(nil)

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NotFound.wrap(err)
	case !strings.Contains(msg, "SQLSTATE"):
		return err
	case strings.Contains(msg, "SQLSTATE 23000"):
		return IntegrityConstraintViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23001"):
		return RestrictViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23502"):
		return NotNullViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23503"):
		return ForeignKeyViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23505"):
		return UniqueViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23514"):
		return CheckViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 23P01"):
		return ExclusionViolation.wrap(err)
	case strings.Contains(msg, "SQLSTATE 57000"):
		return OperatorIntervention.wrap(err)
	case strings.Contains(msg, "SQLSTATE 40000"):
		return TransactionRollback.wrap(err)
	case strings.Contains(msg, "SQLSTATE 40001"):
		return SerializationFailure.wrap(err)
	case strings.Contains(msg, "SQLSTATE 40P01"):
		return DeadLockDetected.wrap(err)
	case strings.Contains(msg, "SQLSTATE 57014"):
		return QueryCanceled.wrap(err)
	case strings.Contains(msg, "SQLSTATE 42000"):
		return TransactionIntegrityConstraintViolation.wrap(err)
	}
	return err
}

func (d *DBErr) Unwrap() error {
	e, ok := d.Payload.(error)
	if !ok {
		return nil
	}
	return e
}
