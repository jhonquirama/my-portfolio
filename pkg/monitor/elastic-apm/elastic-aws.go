package apm

import (
	"github.com/aws/aws-sdk-go/aws/session"
	apmawssdkgo "go.elastic.co/apm/module/apmawssdkgo/v2"
)

func WrapSession(session *session.Session) *session.Session {
	return apmawssdkgo.WrapSession(session)
}
