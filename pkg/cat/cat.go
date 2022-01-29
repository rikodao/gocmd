package cat

import (
	"log"
	"os"
)

type Params struct {
	Dirname    string
}

func Cat(params Params) {
	data, err := readFile(params.Dirname)
	if err != nil {
		log.Fatal(err)
	}
	printFile(data)

}
func readFile(dirname string) ([]byte, error) {
	data, err := os.ReadFile(dirname)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func printFile(data []byte) {
	os.Stdout.Write(data)
}
