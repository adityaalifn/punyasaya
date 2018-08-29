package server

import (
	"punyasaya/pkg/controller"

	"github.com/tokopedia/tdk/go/app/http"
)

type HttpService struct {
}

func NewHttpServer() HttpService {
	return HttpService{}
}

func (s HttpService) RegisterHandler(r *http.Router) {
	r.HandleFunc("/", index, "GET")
	r.HandleFunc("/articles", controller.HandleGetArticle, "GET")
	r.HandleFunc("/articles", controller.HandlePostArticle, "POST")
}

func index(ctx http.TdkContext) error {
	ctx.Writer().Write([]byte("This is index"))
	return nil
}
