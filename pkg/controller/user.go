package controller

import (
	"encoding/json"
	"errors"

	"github.com/adityaalifn/punyasaya/pkg/model"
	"github.com/tokopedia/tdk/go/app/http"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleUserLogin(ctx http.TdkContext) error {
	body := ctx.Body()
	request := request{}
	err := json.Unmarshal(body, &request)
	if err != nil {
		return err
	}
	if request.Username != "kucing" || request.Password != "lucu" {
		return ctx.Error(401, errors.New(string("{status: \"401 Unauthorized\"}")))
	}
	user := model.User{
		Username: request.Username,
		Password: request.Password,
		Name:     "Kucing mania",
	}
	returnObject, _ := json.Marshal(user)
	ctx.Writer().Write(returnObject)
	return nil
}
