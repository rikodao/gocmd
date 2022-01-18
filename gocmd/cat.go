package gocmd

import (
	"log"
	"os"
)
func Cat(dirname string)  {
	data,err := readFile(dirname)
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
func printFile(data []byte)  {
	os.Stdout.Write(data)
}