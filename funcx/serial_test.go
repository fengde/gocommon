package funcx

import (
	"errors"
	"fmt"
	"testing"
)

func TestSerial(t *testing.T) {
	err := Serial(
		func() error {
			fmt.Println(1)
			return nil
		},
		func() error {
			fmt.Println(2)
			return errors.New("happen error")
		},
		func() error {
			fmt.Println(3)
			return nil
		},
	)
	t.Log(err)
}
