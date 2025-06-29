package executor

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	MIN_TIME   = 3 * 60 // в секундах
	DELTA_TIME = 2 * 60
)

// по тз ждём ответа 3-5 минут
func HardIOBoundWork(payload any) (any, error) {
	duration := time.Duration(MIN_TIME+rand.Intn(DELTA_TIME+1)) * time.Second
	time.Sleep(duration)

	result := rand.Intn(10)
	resStr := fmt.Sprintf("io work result = %d", rand.Intn(10))
	if result == 9 {
		return nil, errors.New("some problems in I/O bound task")
	}

	return resStr, nil
}
