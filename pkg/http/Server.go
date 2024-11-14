package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tomegathericon/go-utils/pkg/config"
	"github.com/tomegathericon/go-utils/pkg/log"
	"github.com/tomegathericon/go-utils/pkg/log/gin/middlewares"
	ginMiddleWare "github.com/tomegathericon/go-utils/pkg/tracing/gin/middlewares"
	"github.com/tomegathericon/go-utils/pkg/tracing/tracer"
	"go.opentelemetry.io/otel"
	"net/http"
	"os"
)

func NewServer() {
	l := log.Must(log.LOGFMT).Create()
	ctx := l.NewContext(context.Background())
	err := config.Load(ctx)
	if err != nil {
		panic(err)
	}
	tpc := tracer.NewTraceProviderConfig()
	tpc.SetServiceName(os.Getenv("SERVICE_NAME"))
	tpc.SetServiceVersion(os.Getenv("SERVICE_VERSION"))
	tp, tpErr := tracer.NewHttpTraceProvider(tpc)
	if tpErr != nil {
		panic(tpErr)
	}
	otel.SetTracerProvider(tp)
	router := gin.New()
	router.Use(ginMiddleWare.OpenTelemetryTracing(ctx))
	router.Use(middlewares.Middleware(nil))
	router.GET(os.Getenv("DEFAULT_ROUTE"), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I AM ALIVE!!!!!",
			"code":    http.StatusOK,
		})
	})
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))); err != nil {
		panic(err)
	}
}
