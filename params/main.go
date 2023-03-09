package params

import (
	"errors"
	"flag"
)

type Params struct {
	Question    string
	ResetDialog bool
}

func Parse() (*Params, error) {
	var params Params

	flag.StringVar(&params.Question, "q", "", "your question")
	flag.BoolVar(&params.ResetDialog, "reset", false, "reset dialog")
	flag.Parse()
	if params.Question == "" {
		return &params, errors.New("please use the -q parameter to enter your question.")
	}

	return &params, nil
}
