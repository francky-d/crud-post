package main

import (
	"errors"
	"strconv"
	"time"
)

var ErrNotFound = errors.New("resource not found")
var ErrSomethingWentWrong = errors.New("something went wrong")

func currentTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
