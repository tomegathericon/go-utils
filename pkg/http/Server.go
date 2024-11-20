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
	err := config.Load(context.Background())
	if err != nil {
		panic(err)
	}
	lf, lfErr := log.GetFormat(os.Getenv("LOG_FORMAT"))
	if lfErr != nil {
		panic(lfErr)
	}
	l := log.Must(lf).Create()
	ctx := l.NewContext(context.Background())
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
			"message": os.Getenv("DEFAULT_ROUTE_MESSAGE"),
			"code":    http.StatusOK,
		})
	})
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))); err != nil {
		panic(err)
	}
}
