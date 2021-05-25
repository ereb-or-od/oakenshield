package api

import (
	"github.com/ereb-or-od/kenobi/pkg/controller"
	"github.com/ereb-or-od/oakenshield/pkg/domain"
	"github.com/ereb-or-od/oakenshield/pkg/domain/contracts"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/labstack/echo/v4"
)
const emptyString = ""
type idController struct {
	idFactory         interfaces.OakenshildID
	idEncoderStrategy interfaces.OakenshildIDEncodeStrategy
}

func (i idController) Name() string {
	return emptyString
}

func (i idController) Prefix() string {
	return emptyString
}

func (i idController) Version() string {
	return emptyString
}

func (i idController) Endpoints() *map[string]map[string]echo.HandlerFunc {
	return &map[string]map[string]echo.HandlerFunc{
		"/next": {
			"GET": func(context echo.Context) error {
				contract := contracts.IDContract{
					RawID: i.idFactory.Next(),
				}
				if encodedId, err := i.idEncoderStrategy.Encode(context.QueryParam("encoder"), contract.RawID); err != nil{
					return context.JSON(400, &contracts.ApiContract{
						Message: err.Error(),
					})
				}else{
					contract.EncodedID = encodedId
				}
				return context.JSON(200, &contracts.ApiContract{
					Data: &contract,
				})
			},
		},
	}
}

func NewIdController() controller.HttpController {
	return idController{
		idFactory:         domain.NewOakenshildID(),
		idEncoderStrategy: domain.NewOakenshieldIDEncodeStrategy(),
	}
}
