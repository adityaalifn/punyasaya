package controller

import "github.com/tokopedia/tdk/go/app/http"

func HandleGetArticle(ctx http.TdkContext) error {
	ctx.Writer().Write([]byte("This will give article"))
	return nil
}

func HandlePostArticle(ctx http.TdkContext) error {
	ctx.Writer().Write([]byte("This will store articles"))
	return nil
}
