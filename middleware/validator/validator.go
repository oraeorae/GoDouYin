package validator

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"go_douyin/global/consts"
	"go_douyin/global/variable"
	"go_douyin/utils/response"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login Login

		if err := c.ShouldBindJSON(&login); err != nil {
			variable.ZapLog.Info(err.Error())
			response.Fail(c, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, gin.H{})
			c.Abort()
			return
		}
		variable.ZapLog.Info(login.Username)
		if err := validation.ValidateStruct(&login,
			validation.Field(&login.Username, validation.Required),
			validation.Field(&login.Password, validation.Required),
		); err != nil {
			variable.ZapLog.Info(err.Error())
			response.Fail(c, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, gin.H{})
			c.Abort()
			return
		}
		c.Set("login", login)
		c.Next()
	}
}
