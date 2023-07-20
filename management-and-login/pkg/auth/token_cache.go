package auth

import (
	"context"
)

const (
	ACCESS  = "Access"
	REFRESH = "Refresh"
)

func (a *auth) saveAccessToken(id string) error {
	status := a.Redis.SetEX(context.Background(), id, ACCESS, a.AccessTokenExpirationTime)
	if status != nil {
		return status.Err()
	}
	return ErrRedisSet
}

func (a *auth) saveRefreshToken(id string) error {
	status := a.Redis.SetEX(context.Background(), id, REFRESH, a.RefreshTokenExpirationTime)
	if status != nil {
		return status.Err()
	}
	return ErrRedisSet
}

func (a *auth) removeToken(tokenID string) error {
	status := a.Redis.Del(context.Background(), tokenID)
	if status != nil {
		return status.Err()
	}
	return ErrRedisDel
}

func (a *auth) getToken(id string, kind string) error {
	val := a.Redis.Get(context.Background(), id)
	if val == nil || val.Val() == "" {
		return ErrRedisKeyNotFound
	} else if kind == val.Val() {
		return nil
	}
	return ErrRedisWrongTokenTypeUsed
}
