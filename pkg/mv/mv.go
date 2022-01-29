package mv

import (
	"io/fs"
	"log"
	"os"
	"sort"
)

func Ls(path string, option map[string]string) {
	dir, err := readDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dir {
		printDir(v, option)
	}

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
func printDir(file fs.FileInfo, option map[string]string) {
	fileName := file.Name()

	// 不可視ファイルの表示制御
	if _, ok := option["-a"]; !ok {
		firstchar := fileName[0:1]
		if firstchar == "." {
			return
		}
	}
	// ディレクトリフィルタの制御
	if _, ok := option["-d"]; ok {
		if !file.IsDir() {
			return
		}
	}

	println(file.Name())

}
