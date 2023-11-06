package controller

import (
	userAppSrv "server/application/service/user"
	"server/interfaces/controller/sys"
	"server/interfaces/controller/user"
)

type apiGroup struct {
	Sys  sys.EndpointCtl
	User user.EndpointCtl
}

var APIs *apiGroup

func InitSrvInject(userSrv userAppSrv.Service) {
	APIs = &apiGroup{
		Sys:  sys.EndpointCtl{Srv: userSrv},
		User: user.EndpointCtl{Srv: userSrv},
	}
}
