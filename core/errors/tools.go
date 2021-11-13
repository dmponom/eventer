package errors

import "errors"

func GetStatusCodeFromErr(err error) int {
	target := &CustomError{}
	if errors.As(err, &target) {
		return target.StatusCode
	}
	return 0
}

func IsEqual(err, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}

	return err.Error() == target.Error()
}
