package runtime

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	"github.com/labstack/echo/v4"
)

type StrictIrisHandlerFunc func(ctx iris.Context, request interface{}) (response interface{}, err error)

type StrictIrisMiddlewareFunc func(f StrictIrisHandlerFunc, operationID string) StrictIrisHandlerFunc

type StrictEchoHandlerFunc func(ctx echo.Context, request interface{}) (response interface{}, err error)

type StrictEchoMiddlewareFunc func(f StrictEchoHandlerFunc, operationID string) StrictEchoHandlerFunc

type StrictHttpHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error)

type StrictHttpMiddlewareFunc func(f StrictHttpHandlerFunc, operationID string) StrictHttpHandlerFunc

type StrictGinHandlerFunc func(ctx *gin.Context, request interface{}) (response interface{}, err error)

type StrictGinMiddlewareFunc func(f StrictGinHandlerFunc, operationID string) StrictGinHandlerFunc
