package ls

import (
	"io/fs"
	"os"
	"sort"
)
type optionalFunction func(*Option)

func Directory(directory bool) optionalFunction {
	return func(o *Option) {
		o.Directory = directory
	}
}
func All(all bool) optionalFunction {
	return func(o *Option) {
		o.All = all
	}
}

type Option struct {
	Directory bool
	All       bool
}

func NewOption(options ...optionalFunction) *Option {

	srv := Option{
		Directory: false,
		All: false,
	}
	for _, o := range options {
		o(&srv)
	}
	return &srv
}




type Params struct {
	Path    string
	Option Option
}
func Ls(params Params) ([]fs.FileInfo, error) {
	dir, err := readDir(params.Path)
	if err != nil {
		return nil, err

	}
	for _, v := range dir {
		printDir(v, params.Option)
	}
	return dir, nil

}
func readDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
func printDir(file fs.FileInfo, option Option) {
	fileName := file.Name()

	// 不可視ファイルの表示制御
	if !option.All {
		firstchar := fileName[0:1]
		if firstchar == "." {
			return
		}
	}
	// ディレクトリフィルタの制御
	if option.Directory {
		if !file.IsDir() {
			return
		}
	}

	println(file.Name())

}
