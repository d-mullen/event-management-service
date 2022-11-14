package matchers

import "errors"

var (
	errBothArgsNil = errors.New("refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead")
	errExpectedNil = errors.New("refusing to compare expected to <nil>.\nBe explicit and use BeNil() instead")
)

func handleNilArgs(actual, expected interface{}) (ok bool, err error) {
	if actual == nil && expected == nil {
		return false, errBothArgsNil
	} else if expected == nil {
		return false, errExpectedNil
	} else if actual == nil {
		return false, nil
	}

	return true, nil
}
