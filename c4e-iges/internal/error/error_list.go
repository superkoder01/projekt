package error

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	InternalError = fmt.Errorf("internal error")
	a             = errors.Wrap(InternalError, "aaa")
)
