package executor

import "fmt"

const prefix = "executor: "

var (
	ErrInvalidKey = fmt.Errorf("%sinvalid key", prefix)
)
