package server

import (
	"github.com/adityaalifn/punyasaya/pkg/controller"
	"github.com/tokopedia/tdk/go/app/http"
)

type HttpService struct {
}

func NewHttpServer() HttpService {
	return HttpService{}
}

func (s HttpService) RegisterHandler(r *http.Router) {
	r.HandleFunc("/", index, "GET")
	r.HandleFunc("/articles/get", controller.HandleGetArticle, "GET")
	r.HandleFunc("/articles/post", controller.HandlePostArticle, "POST")

	r.HandleFunc("/users/login", controller.HandleUserLogin, "POST")
}

func index(ctx http.TdkContext) error {
	ctx.Writer().Write([]byte("This is index"))
	return nil
}
