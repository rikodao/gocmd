package main

import (
	"github.com/rikodao/gocmd/gocmd/ls"
)

func main() {
	// options := curl.NewOption(curl.Method(curl.GET))
	// params := curl.Params{
	// 	Url:    "https://www.youtube.com/watch?v=Ozc1xmgTwhI",
	// 	Option: *options,
	// 	Body:   nil,
	// }

	// bodyString, err := curl.Curl(params)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// println(string(bodyString))

	options := ls.NewOption(ls.All(true))
	params := ls.Params{
		Path:    ".",
		Option: *options,
	}

	ls.Ls(params)


}
