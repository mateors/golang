package json

import (
	"encoding/json"
	"errors"
	"urlshortener/shortener"
)

type Redirect struct{}

// Encode(input *Redirect) ([]byte, error)
// 	Decode(input []byte) (*Redirect, error)

func (s *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {

	msg, err := json.Marshal(input)
	if err != nil {

		return nil, errors.New("unable to encode")
	}
	return msg, nil
}

func (s *Redirect) Decode(input []byte) (*shortener.Redirect, error) {

	redirect := &shortener.Redirect{}
	err := json.Unmarshal(input, redirect)
	if err != nil {
		return nil, errors.New("unable to decode")
	}
	return redirect, nil
}
