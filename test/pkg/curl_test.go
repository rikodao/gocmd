package pkg

import (
	"bytes"
	"testing"

	"github.com/rikodao/gocmd/pkg/curl"
	"golang.org/x/net/html"
)

func TestCurl(t *testing.T) {
	options := curl.NewOption()
	options.Method(curl.GET)
	params := curl.Params{
		Url:    "https://google.com",
		Option: *options,
		Body:   nil,
	}

	bodyString, err := curl.Run(params)
	if err != nil {
		t.Error(err)
	}

	_, err = html.Parse(bytes.NewReader(bodyString))
	if err != nil {
		t.Error(err)
	}

}
