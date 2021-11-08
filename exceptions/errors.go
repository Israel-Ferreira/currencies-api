package exceptions

import "errors"

var ErrFromCoinEmptyOrNil = errors.New("err: from coin shouldn't empty or nil")

var ErrToCoinEmptyOrNil = errors.New("err: to coin shouldn't empty or nil")
