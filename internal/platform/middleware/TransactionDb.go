package middleware

import (
	"github.com/fede/golang_api/internal/platform/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"net/http"
)

var (
	logger, _   = zap.NewProduction()
	sugarLogger = logger.Sugar()
)

type DBTransaction struct {
	db *gorm.DB
}

func NewDBTransaction(db *gorm.DB) *DBTransaction {
	return &DBTransaction{
		db: db,
	}
}

func (t *DBTransaction) Handler(c *gin.Context) {
	tx := t.db.Begin()

	if c.Request.Method == http.MethodGet {
		c.Request = storage.RequestWithDBContext(c.Request, tx)
		c.Next()
		return
	}

	c.Request = storage.RequestWithDBContext(c.Request, tx)

	defer t.rollback(tx)

	c.Next()
	if len(c.Errors) == 0 {
		err := tx.Commit().Error
		if err != nil {
			sugarLogger.Error("failed in commit with error: ", err)
		}
	} else {
		err := tx.Rollback()
		if err != nil {
			sugarLogger.Error("failed in rollback with error: ", err)
		}
	}
}

func (t *DBTransaction) rollback(tx *gorm.DB) {
	if err := recover(); err != nil {
		e := tx.Rollback().Error
		if e != nil {
			sugarLogger.Error("failed in rollback", e)
		}
		panic(err)
	}
}
