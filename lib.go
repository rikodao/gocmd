package main

import (
	"io/fs"

	"github.com/rikodao/gocmd/pkg/cat"
	"github.com/rikodao/gocmd/pkg/curl"
	"github.com/rikodao/gocmd/pkg/ls"
)

type Gocmd struct {
	Curl func(curl.Params) ([]byte, error)
	Ls func(ls.Params) ([]fs.FileInfo, error)
	Cat func(cat.Params) 
}

func New() *Gocmd {
	return &Gocmd{
		Curl: curl.Curl,
		Ls: ls.Ls,
		Cat: cat.Cat,
	}

}
