package pkg

import (
	"testing"

	"github.com/rikodao/gocmd/pkg/cat"
)

func TestCat(t *testing.T) {
	params := cat.Params{Dirname: "./cat_test.go"}

	file, err := cat.Run(params)
	if err != nil {
		t.Error(err)
	}
	println(file)

}
