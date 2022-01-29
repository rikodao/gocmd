package pkg

import (
	"testing"

	"github.com/rikodao/gocmd/pkg/ls"
)

func TestLs(t *testing.T) {
	options := ls.NewOption()

	params := ls.Params{
		Path:   ".",
		Option: *options,
	}

	files, err := ls.Run(params)
	if err != nil {
		t.Error(err)
	}

	for _, f := range files {
		if f.Name()[0:1] == "." {
			t.Error(err)
		}
	}
}

func TestLsAll(t *testing.T) {
	options := ls.NewOption()
	options.All(true)

	params := ls.Params{
		Path:   ".",
		Option: *options,
	}

	files, err := ls.Run(params)
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		println(f.Name())
	}
}
func TestLsDirectory(t *testing.T) {
	options := ls.NewOption()
	options.Directory(true)

	params := ls.Params{
		Path:   ".",
		Option: *options,
	}

	files, err := ls.Run(params)
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			println(f.Name())
			t.Error()

		}
	}
}
