package context

import (
	"context"
	"strconv"

	"server/domain/user/entity/valueobj"
	"server/infrastructure/common/jwt"
)

// Store stores user-related information in the request context
func Store(ctx context.Context, claims *jwt.PayloadClaims) context.Context {
	// Store username in request context
	ctx = context.WithValue(ctx, valueobj.CtxUsername, claims.Username)

	// Store userID in request context
	userID, _ := strconv.Atoi(claims.Audience)
	ctx = context.WithValue(ctx, valueobj.CtxUserID, userID)

	// Store employee ID in request context
	ctx = context.WithValue(ctx, valueobj.CtxEmpNO, claims.EmpNO)

	// Store user roles in request context
	ctx = context.WithValue(ctx, valueobj.CtxRoles, claims.Roles)

	// Check for admin privileges
	isAdmin := valueobj.IsAdminFalse
	for _, role := range claims.Roles {
		if role == valueobj.AdminRoleName || role == valueobj.RootRoleName {
			isAdmin = valueobj.IsAdminTrue
			break
		}
	}

	// Store admin privilege flag in request context
	ctx = context.WithValue(ctx, valueobj.CtxIsAdmin, isAdmin)

	return ctx
}
