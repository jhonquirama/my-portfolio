package apm

import (
	"context"

	apm "go.elastic.co/apm/v2"
)

func CaptureError(ctx context.Context, err error) {
	if apmError := apm.CaptureError(ctx, err); apmError != nil && apmError.ErrorData != nil {
		apmError.SetStacktrace(3)
		apmError.Send()
	}
}
