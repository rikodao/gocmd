package curl

import (
	"io"
	"log"
	"net/http"
)

type method string

const (
	GET    method = "GET"
	PUT           = "PUT"
	POST          = "POST"
	DELETE        = "DELETE"
)

type Option struct {
	method method
}

func (srv *Option) Method(method method) *Option {
	srv.method = method
	return srv
}

func NewOption() *Option {
	return &Option{
		method: GET,
	}
}

type Params struct {
	Url    string
	Option Option
	Body   io.Reader
}

func Run(params Params) ([]byte, error) {
	var resp *http.Response
	client := &http.Client{}

	req, err := http.NewRequest(string(params.Option.method), params.Url, params.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil

}
