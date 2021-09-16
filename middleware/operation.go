package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func OperationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body []byte

		var err error
		body, err = ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			g.TENANCY_LOG.Error("read body from request error:", zap.Any("err", err))
		} else {
			// ioutil.ReadAll 读取数据后重新回写数据
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		record := model.SysOperationRecord{
			BaseOperationRecord: model.BaseOperationRecord{
				Ip:           ctx.ClientIP(),
				Method:       ctx.Request.Method,
				Path:         ctx.Request.URL.Path,
				Agent:        ctx.Request.UserAgent(),
				Body:         string(body),
				UserID:       multi.GetUserId(ctx),
				SysTenancyID: multi.GetTenancyId(ctx),
			},
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = writer
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = ctx.Writer.Status()
		record.Latency = latency
		// 查询接口日志内容太多影响性能
		if ctx.Request.Method != http.MethodGet {
			record.Resp = writer.body.String()
		}

		if err := service.CreateSysOperationRecord(record); err != nil {
			g.TENANCY_LOG.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
