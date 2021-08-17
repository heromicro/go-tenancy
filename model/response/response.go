package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"status"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

const (
	SUCCESS           = 2000
	BAD_REQUEST_ERROR = 4000
	NEED_INIT_ERROR   = 4007
)

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	// 开始时间
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, "操作成功", ctx)
}

func OkWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(http.StatusOK, data, "操作成功", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusOK, data, message, ctx)
}

func Fail(ctx *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, "操作失败", ctx)
}

func UnauthorizedFailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusUnauthorized, map[string]interface{}{}, message, ctx)
}

func UnauthorizedFailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusUnauthorized, data, message, ctx)
}

func ForbiddenFailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusForbidden, map[string]interface{}{}, message, ctx)
}

func FailWithMessage(message string, ctx *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(http.StatusBadRequest, data, message, ctx)
}
func NeedInitWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(NEED_INIT_ERROR, data, message, ctx)
}
