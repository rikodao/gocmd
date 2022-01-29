package ls

import (
	"io/fs"
	"os"
	"sort"
)

type optionalFunction func(*Option)

type Option struct {
	directory bool
	all       bool
}

func NewOption() *Option {
	return &Option{
		directory: false,
		all:       false,
	}

}
func (srv *Option) Directory(directory bool) *Option {
	srv.directory = directory
	return srv
}
func (srv *Option) All(all bool) *Option {
	srv.all = all
	return srv
}

type Params struct {
	Path   string
	Option Option
}

func Run(params Params) ([]fs.FileInfo, error) {
	f, err := os.Open(params.Path)
	if err != nil {
		return nil, err
	}
	allFile, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	var files []fs.FileInfo
	for _, file := range allFile {
		// 不可視ファイルの表示制御
		if !params.Option.all {
			fileName := file.Name()
			firstchar := fileName[0:1]
			if firstchar == "." {
				continue
			}
		}
		// ディレクトリフィルタの制御
		if params.Option.directory {
			if !file.IsDir() {
				continue
			}
		}
		files = append(files, file)
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	return files, nil

}
