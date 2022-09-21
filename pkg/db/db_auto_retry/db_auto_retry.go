package db_auto_retry

import (
	"github.com/avast/retry-go"
	"github.com/recative/recative-backend/pkg/db/db_err"
)

func SerializableTransactionAutoRetry(retryTimes uint, errorReturnFunc func() error) error {
	err := retry.Do(
		errorReturnFunc,
		retry.Attempts(retryTimes),
		retry.RetryIf(func(err error) bool {
			if db_err.SerializationFailure.Is(err) {
				return true
			}
			return false
		}),
	)

	retryError, ok := err.(retry.Error)
	if !ok {
		return err
	}

	// find first error
	for _, err := range []error(retryError) {
		if err != nil {
			return err
		}
	}

	return err
}
