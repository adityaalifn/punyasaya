package controller

import (
	"encoding/json"

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
	_ = json.Unmarshal(body, &request)
	user := model.User{Username: request.Username, Password: request.Password}
	returnObject, _ := json.Marshal(user)
	ctx.Writer().Write(returnObject)
	return nil
}
