package cat

import (
	"os"
)

type Params struct {
	Dirname string
}

func Run(params Params) ([]byte, error) {
	return os.ReadFile(params.Dirname)
}
