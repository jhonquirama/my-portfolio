package apm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	apm "go.elastic.co/apm/v2"
)

const (
	elasticAPMServerURL   = "ELASTIC_APM_SERVER_URL"
	elasticApmSecretToken = "ELASTIC_APM_SECRET_TOKEN" // nolint: gosec
	elasticApmLogFile     = "ELASTIC_APM_LOG_FILE"
	elasticApmLogLevel    = "ELASTIC_APM_LOG_LEVEL"
	elasticApmEnvironment = "ELASTIC_APM_ENVIRONMENT"

	serviceVersion = "1.0"
)

type (
	Config interface {
		ApmServiceName() string
		ApmServerHost() string
		ApmSecretToken() string
		ApmLogFile() string
		ApmLogLevel() string
		ApmEnvironment() string
	}

	Tracer struct {
		*apm.Tracer
	}
)

func (apm *Tracer) Close() {
	apm.Tracer.Close()
}

func (apm *Tracer) Flush(abort <-chan struct{}) {
	apm.Tracer.Flush(abort)
}

func NewAPM(cnf Config) (*Tracer, error) {
	var (
		tracer *apm.Tracer
		err    error
	)

	setEnv(cnf)

	defer unsetEnv()

	if tracer, err = apm.NewTracer(cnf.ApmServiceName(), serviceVersion); err != nil {
		return nil, err
	}

	return &Tracer{
		Tracer: tracer,
	}, nil
}

func setEnv(cnf Config) {
	os.Setenv(elasticAPMServerURL, cnf.ApmServerHost())
	os.Setenv(elasticApmSecretToken, cnf.ApmSecretToken())
	os.Setenv(elasticApmLogFile, cnf.ApmLogFile())
	os.Setenv(elasticApmLogLevel, cnf.ApmLogLevel())
	os.Setenv(elasticApmEnvironment, cnf.ApmEnvironment())
}

func unsetEnv() {
	os.Unsetenv(elasticAPMServerURL)
	os.Unsetenv(elasticApmSecretToken)
	os.Unsetenv(elasticApmLogFile)
	os.Unsetenv(elasticApmLogLevel)
	os.Unsetenv(elasticApmEnvironment)
}

func TraceIDContext(ctx context.Context) interface{} {
	tx := TransactionFromContext(ctx)

	if tx == nil {
		return nil
	}

	return tx.Transaction.TraceContext().Trace
}

func GetFuncNameFromCtx(ctx context.Context) *string {
	var (
		funcName string
	)

	if result := apm.SpanFromContext(ctx); result != nil && result.SpanData != nil {
		funcName = result.Type + "." + result.Name
	} else if result := apm.TransactionFromContext(ctx); result != nil && result.TransactionData != nil {
		funcName = result.Type + "." + result.Name
	}

	return &funcName
}

func GetCallerName(skip int) (string, int) {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", 0
	}

	f := runtime.FuncForPC(pc)
	if f == nil {
		return "", 0
	}

	_, line := f.FileLine(pc)

	return f.Name(), line
}

func DetachedContext(ctx context.Context) context.Context {
	return apm.DetachedContext(ctx)
}

func getLabels(key string, value interface{}) map[string]string {
	var (
		data   = make(map[string]string)
		length = 1020
	)

	bytes, _ := json.Marshal(value)
	newValue := string(bytes)
	if len(newValue) <= length {
		data[key] = newValue
	} else {
		for i := 0; i < len(newValue); i += length {
			end := i + length
			if end > len(newValue) {
				end = len(newValue)
			}

			data[fmt.Sprintf("%s_%d", key, i)] = newValue[i:end]
		}
	}

	return data
}
