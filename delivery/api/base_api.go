package api

import (
	"fmt"
	"net/http"

	"enigmacamp.com/bank/delivery/commonresp"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	err := c.ShouldBindJSON(body)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) Succes(c *gin.Context, message string, data interface{}) {
	commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("0", message, data))
}

func (b *BaseApi) FailedRequest(c *gin.Context, serviceName, errCode, message string) {
	e := commonresp.NewErrorMessage(http.StatusBadRequest, serviceName, errCode, message)
	c.Error(fmt.Errorf("%s", e.ToJson()))
}

func (b *BaseApi) Failed(c *gin.Context, serviceName, errCode, message string) {
	e := commonresp.NewErrorMessage(http.StatusInternalServerError, serviceName, errCode, message)
	c.Error(fmt.Errorf("%s", e.ToJson()))
}
