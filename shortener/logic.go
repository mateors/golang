package shortener

import (
	"errors"
	"fmt"
)

var (
	ErrRedirectNotFound = errors.New("Redirect not found")
	ErrRedirectInvalid  = errors.New("Redirect invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {

	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {

	fmt.Println("logic.Find()")
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {

	//validate logic goes here
	//redirect.Code = "3434234dfgdfgd"
	//redirect.CreateTime = time.Now().Unix()
	fmt.Println("logic.Store()")
	return r.redirectRepo.Store(redirect)
}

// test only
func (r *redirectService) Test(redirect *Redirect) error {

	return r.redirectRepo.Store(redirect)
}
