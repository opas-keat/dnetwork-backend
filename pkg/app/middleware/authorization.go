package middleware

import (
	"ect/dnetwork/backend/pkg/common"
	"ect/dnetwork/backend/pkg/common/logger"

	"github.com/casbin/casbin/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUnauthorize = common.NewForbiddenError("you are not allowed to access this resource")
	ErrEnforce     = common.NewUnexpectedError("error occurred while enforce")
)

func Authorize(enforcer *casbin.Enforcer) common.HandleFunc {
	return func(c common.HContext) error {
		// ถ้าเป็น public route ให้ข้ามการตรวจสอบไป
		public := c.Locals("public").(bool)
		if public {
			return c.Next()
		}

		// ดึงค่า role ออกมา
		u := c.Locals("user").(jwt.MapClaims)
		role := u["role"].(string)

		// ส่ง request เข้าไปตรวจสอบ
		ok, err := enforcer.Enforce(role, c.Path(), c.Method())

		if err != nil {
			logger.ErrorWithReqId(err.Error(), c.RequestId())
			return common.ResponseError(c, ErrEnforce)
		}

		if !ok {
			return common.ResponseError(c, ErrUnauthorize)
		}

		return c.Next()
	}
}
