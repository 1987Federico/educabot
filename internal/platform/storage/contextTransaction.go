package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type contextKey int

const DBContextKey contextKey = iota

func NewContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, DBContextKey, db)
}

func FromContext(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return nil
	}

	h, exists := ctx.Value(DBContextKey).(*gorm.DB)
	if exists {
		return h
	}

	panic(errors.New("database context doesn't exists"))
}

func RequestWithDBContext(req *http.Request, db *gorm.DB) *http.Request {
	ctx := req.Context()
	ctx = NewContext(ctx, db)
	return req.WithContext(ctx)
}
