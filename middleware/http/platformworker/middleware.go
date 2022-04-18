package platformworker

import (
	"fmt"
	"github.com/dapr/components-contrib/middleware"
	"github.com/dapr/kit/logger"
	"github.com/valyala/fasthttp"
	//"github.com/go-resty/resty/v2"
)

const (
	metaTagReportingAppID = "reportingAppID"
)

type workermetadata struct {
	ReportingHost            string `json:"reportingHost"`
	ReportingProcessPathURL  string `json:"reportingProcessPathURL"`
	ReportingRollbackPathURL string `json:"reportingRollbackPathURL"`
}

type Middleware struct {
	logger logger.Logger
}

func NewMiddleware(logger logger.Logger) *Middleware {
	return &Middleware{logger: logger}
}

func (m *Middleware) getNativeMetadata(metadata middleware.Metadata) (*workermetadata, error) {
	var middlewareMetadata workermetadata

	if val, ok := metadata.Properties[metaTagReportingAppID]; ok {
		middlewareMetadata.ReportingProcessPathURL = val
	}

	return &middlewareMetadata, nil
}

// GetHandler returns the HTTP handler provided by the middleware.
func (m *Middleware) GetHandler(metadata middleware.Metadata) (func(h fasthttp.RequestHandler) fasthttp.RequestHandler, error) {
	//meta, err := m.getNativeMetadata(metadata)
	//if err != nil {
	//	return nil, err
	//}

	//processURL := meta.ReportingHost + meta.ReportingProcessPathURL

	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			fmt.Println(">>>>>> REQUEST: ", ctx.Request.String())

			h(ctx)

			fmt.Println(">>>>>> RESPONSE: ", ctx.Response.String())
		}
	}, nil
}
