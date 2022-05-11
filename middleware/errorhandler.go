package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shelld1t/core/log"
	"github.com/shelld1t/core/model"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		if err == nil {
			return
		}
		if !ctx.Response().Committed {
			//todo: refact брать код ошибки из сущности error
			errorResponse := model.ErrorResp(err)
			encError, errEncode := errorResponse.Encode()
			if errEncode != nil {
				//todo: refact отдавать пятисотку нормально
				errEncode = ctx.Blob(http.StatusInternalServerError, echo.MIMEApplicationJSON, []byte("{internal httpServer error}"))
				log.Log(ctx.Request().Context()).Error("error encode errorResponse", errEncode)
				return
			}
			err = sendJson(ctx, errorResponse.Code(), encError)
			if err != nil {
				log.Log(ctx.Request().Context()).Error("write response error", err)
			}
		}
	}
}

func sendJson(ctx echo.Context, code int, content []byte) error {
	return ctx.Blob(code, echo.MIMEApplicationJSON, content)
}
