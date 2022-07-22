package middleware

import (
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Error struct {
	logger *zap.SugaredLogger
}

func NewError(logger *zap.SugaredLogger) *Error {
	return &Error{
		logger: logger,
	}
}

func (h *Error) Handler(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer h.recovery(ctx)

	ctx.Next()

	for _, err := range ctx.Errors {
		h.abortWithAPIError(ctx, err.Err)
		if ctx.IsAborted() {
			break
		}
	}
}

func (h *Error) recovery(ctx *gin.Context) {
	err := recover()
	if err != nil {
		switch er := err.(type) {
		case error:
			h.abortWithAPIError(ctx, er)
		default:
			h.logger.Error("Recovery from panic with errorCustom: ", er)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorCustom.InternalServerApiError("unknown error", "An errorCustom occurred"))
		}
	}
}

func (h *Error) abortWithAPIError(ctx *gin.Context, err error) {
	h.logger.Errorf("errorCustom: %v ", err)
	errorCustom.RespondError(ctx, err)
}
