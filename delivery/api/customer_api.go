package api

import (
	"net/http"

	"enigmacamp.com/bank/delivery/appreq"
	"enigmacamp.com/bank/delivery/tokenauth"
	"enigmacamp.com/bank/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
	BaseApi
	loginUseCase    usecase.LoginUseCase
	transferUseCase usecase.TransferUseCase
	tokenauth       tokenauth.Token
}

func (a *CustomerApi) userLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var custReq appreq.CustomerRequestLogin
		err := a.ParseRequestBody(c, &custReq)
		if err != nil {
			c.Error(err)
			return
		}
		err = a.loginUseCase.Login(custReq.AccountNumber, custReq.UserPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "LOGIN FAILED")
			return
		}
		token := a.tokenauth.GetToken()
		c.JSON(http.StatusOK, token)
	}
}

func (a *CustomerApi) Transfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var custReq appreq.CustomerRequestPayment
		err := a.ParseRequestBody(c, &custReq)
		if err != nil {
			c.Error(err)
			return
		}
		if custReq.Token != a.tokenauth.GetToken() {
			c.JSON(http.StatusBadRequest, "INVALID TOKEN")
			return
		}
		a.transferUseCase.Transfer(custReq.SenderAccountNumber, custReq.ReceiverAccountNumber, custReq.AmountTransfer)
		c.JSON(http.StatusOK, "TRANSFER SUCCESS")
	}
}

func NewCustomerApi(customerRoute *gin.RouterGroup, loginUseCase usecase.LoginUseCase, transferUseCase usecase.TransferUseCase) {
	api := CustomerApi{
		loginUseCase:    loginUseCase,
		transferUseCase: transferUseCase,
	}
	customerRoute.POST("/login", api.userLogin())
	customerRoute.POST("/transfer", api.Transfer())
}
