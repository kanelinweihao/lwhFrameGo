package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParams struct {
	typeInterface.EntityParams `validate:"-"`
	EntityAdminUser            *conf.EntityAdminUser `validate:"-"`
	AdminUserName              string                `validate:"required"`
	AdminUserPassword          string                `validate:"required"`
	AdminUserId                int                   `validate:"omitempty"`
	AdminUserRole              int                   `validate:"omitempty"`
	JwtToken                   string                `validate:"omitempty"`
}
