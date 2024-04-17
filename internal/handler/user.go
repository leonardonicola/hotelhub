package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/leonardonicola/tickethub/internal/dto"
)

// @Summary		Register
// @Description	Create an account
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CreateUserOutputDTO
// @Router			/register [post]
func RegisterHandler(ctx *gin.Context) {
	request := dto.CreateUserInputDTO{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	errs := validate.Struct(request)
	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errs.(validator.ValidationErrors).Translate(trans))
		return
	}
	res := dto.CreateUserOutputDTO{
		ID:      uuid.NewString(),
		Name:    request.Name,
		Address: request.Address,
		Surname: request.Surname,
		Email:   request.Email,
		CPF:     request.CPF,
	}
	log.Info("Usuário criado com sucesso!")
	ctx.JSON(200, res)
}
