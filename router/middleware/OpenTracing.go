package middleware

import (
	"fmt"
	"io"
	"time"
    "context"

	//"gin-api/app_const"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"github.com/gorefa/log"

	"github.com/pengsrc/go-shared/buffer"
)

func TraceSpan() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			c.Next()
		}
		buf := buffer.GlobalBytesPool().Get()
		buf.AppendString("HTTP ")
		buf.AppendString(c.Request.Method)

		span := opentracing.StartSpan(buf.String())
		rc := opentracing.ContextWithSpan(c.Request.Context(),span)

		if sc , ok := span.Context().(jaeger.SpanContext);ok {
			//rc = context.WithValue(rc,FromContext(c.Request.Context()),sc.TraceID().String())
			rc = context.WithValue(rc,c.Request.Header.Get("X-Request-ID"),sc.TraceID().String())
		}
		span.Finish()

		c.Next()
	}
}



func OpenTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span
		tracer, closer := NewJaegerTracer(viper.GetString("JaegerHostPort"))
		defer closer.Close()

		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		//spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			log.Error("spanCtx error",err)
			fmt.Println("opentracing:start")
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			fmt.Println("opentracing:extract")
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		c.Set("Tracer", tracer)
		c.Set("ParentSpanContext", parentSpan.Context())
		c.Next()
	}
}

func NewJaegerTracer(jaegerHostPort string) (opentracing.Tracer, io.Closer) {

	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			BufferFlushInterval: 1* time.Second,
			LocalAgentHostPort: jaegerHostPort,
		},

		ServiceName: viper.GetString("jaeger_server_name"),
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
