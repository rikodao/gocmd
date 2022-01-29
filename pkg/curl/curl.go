package curl

import (
	"io"
	"log"
	"net/http"
)

type optionalFunction func(*Option)

func Method(method method) optionalFunction {
	return func(o *Option) {
		o.Method = method
	}
}

type Option struct {
	Method method
}

func NewOption(options ...optionalFunction) *Option {

	srv := Option{
		Method: GET,
	}
	for _, o := range options {
		o(&srv)
	}
	return &srv
}

type method string

const (
	GET    method = "GET"
	PUT           = "PUT"
	POST          = "POST"
	DELETE        = "DELETE"
)

type Params struct {
	Url    string
	Option Option
	Body   io.Reader
}

func Curl(params Params) ([]byte, error) {
	return request(params.Url, params.Option.Method, params.Body)
}
func request(url string, method method, body io.Reader) ([]byte, error) {
	var resp *http.Response
	client := &http.Client{}

	req, err := http.NewRequest(string(method), url, body)
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

// func perseOption(url string,method string) ([]byte, error) {
// 	resp, _ := http.Get(url)
//   	defer resp.Body.Close()

//   byteArray, err := io.ReadAll(resp.Body)
//   if err != nil {
// 	  return nil, err
//   }
// 	return byteArray, nil
// }
