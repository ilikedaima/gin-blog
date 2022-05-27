package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form auth
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	authService := auth_service.Auth{
		Username: form.Username,
		Password: form.Password,

	}
	exists, err := authService.Check()

	if err!=nil {
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_CHECK_TOKEN_FAIL,nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK,e.ERROR_AUTH,nil)
		return
	}

	
	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_CHECK_TOKEN_FAIL,nil)
		return
	}

	
	appG.Response(http.StatusOK,e.SUCCESS,map[string]interface{}{
		"token": token,
	})


	
}