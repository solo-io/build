package envutils

import (
	"github.com/pkg/errors"
	"github.com/solo-io/build/pkg/constants"
)

func StringForBoolToEnv(b bool) string {
	if b {
		return constants.PrintEnvTrue
	}
	return constants.PrintEnvFalse
}

func BoolFromEnvString(s string) (bool, error) {
	if s == constants.PrintEnvTrue {
		return true, nil
	}
	if s == constants.PrintEnvFalse {
		return false, nil
	}
	return false, errors.Errorf("invalid boolean env var: %s", s)
}
