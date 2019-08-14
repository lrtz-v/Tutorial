package common

import (
	"testing"
)

func TestConnect(T *testing.T) {
	_, err := Connect()
	if err != nil {
		T.Fail()
	}
}
