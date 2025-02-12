package apm

import (
	"net/http"

	apmhttp "go.elastic.co/apm/module/apmhttp/v2"
	apm "go.elastic.co/apm/v2"
)

func WrapClient(client *http.Client) *http.Client {
	return apmhttp.WrapClient(client)
}

func FormatTraceParentHeader(c apm.TraceContext) string {
	return apmhttp.FormatTraceparentHeader(c)
}

func TraceParentHeader() string {
	return apmhttp.W3CTraceparentHeader
}
