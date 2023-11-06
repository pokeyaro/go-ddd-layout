package context

import (
	"context"
	"errors"

	"server/domain/user/entity/valueobj"
)

func GetUsername(ctx context.Context) (string, error) {
	username, ok := ctx.Value(valueobj.CtxUsername).(string)
	if !ok {
		return "", errors.New("failed to retrieve the username from the request context")
	} else {
		return username, nil
	}
}

func GetUID(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(valueobj.CtxUserID).(int)
	if !ok {
		return 0, errors.New("failed to retrieve the uid from the request context")
	} else {
		return userID, nil
	}
}

func GetEmployeeID(ctx context.Context) (string, error) {
	empID, ok := ctx.Value(valueobj.CtxEmpNO).(string)
	if !ok {
		return "", errors.New("failed to retrieve the employee ID from the request context")
	} else {
		return empID, nil
	}
}

func GetRoleList(ctx context.Context) ([]string, error) {
	roles, ok := ctx.Value(valueobj.CtxRoles).([]string)
	if !ok {
		return nil, errors.New("failed to retrieve the role list from the request context")
	} else {
		return roles, nil
	}
}

func GetIsAdminPerm(ctx context.Context) (bool, error) {
	isAdmin, ok := ctx.Value(valueobj.CtxIsAdmin).(bool)
	if !ok {
		return false, errors.New("failed to retrieve the admin permission flag from the request context")
	} else {
		return isAdmin, nil
	}
}
