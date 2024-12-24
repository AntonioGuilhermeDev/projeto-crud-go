package controller

import (
	"github.com/AntonioGuilhermeDev/projeto-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada")
	ctx.JSON(err.Code, err)
}
