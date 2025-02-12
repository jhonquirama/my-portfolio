package apm

import (
	"context"

	"github.com/sirupsen/logrus"
	apmlogrus "go.elastic.co/apm/module/apmlogrus"
)

func TraceContext(ctx context.Context) logrus.Fields {
	return apmlogrus.TraceContext(ctx)
}
