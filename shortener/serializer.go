package shortener

type RedirectSerializer interface {
	Encode(input *Redirect) ([]byte, error)
	Decode(input []byte) (*Redirect, error)
}
