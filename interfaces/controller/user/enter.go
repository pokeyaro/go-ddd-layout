package user

import (
	"server/application/service/user"
)

type EndpointCtl struct {
	Srv user.Service
}
