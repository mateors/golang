package api

import (
	js "urlshortener/serializer/json"
	"urlshortener/shortener"
)

type Handler struct {
	RedirectService shortener.RedirectService
}

// type DataHandler interface{

// }

// func NewHandler(redirectService shortener.RedirectService) DataHandler {
// 	return &handler{redirectService: redirectService}
// }

func (h *Handler) Serializer(contentType string) shortener.RedirectSerializer {

	
	if contentType == "application/x-msgpack" {
		//return &ms.Redirect{}
	}
	return &js.Redirect{}
}
