package apm

import (
	"context"

	apm "go.elastic.co/apm/v2"
)

const (
	Request TransactionType = "request"
)

type (
	TransactionType string

	Transaction struct {
		*apm.Transaction
	}
)

func (t *Tracer) NewTransaction(
	ctx context.Context,
	functionName string,
	transactionType TransactionType,
) (*Transaction, context.Context) {
	transaction := &Transaction{
		Transaction: t.Tracer.StartTransaction(functionName, string(transactionType)),
	}

	return transaction, ContextWithTransaction(ctx, transaction)
}

func (t *Transaction) End() {
	t.Transaction.End()
}

func (t *Transaction) SetLabel(key string, value interface{}) {
	labels := getLabels(key, value)
	for k, v := range labels {
		t.Transaction.Context.SetLabel(k, v)
	}
}

func TransactionFromContext(ctx context.Context) *Transaction {
	return &Transaction{apm.TransactionFromContext(ctx)}
}

func ContextWithTransaction(ctx context.Context, transaction *Transaction) context.Context {
	return apm.ContextWithTransaction(ctx, transaction.Transaction)
}

func TransactionIDContext(ctx context.Context) interface{} {
	tx := TransactionFromContext(ctx)

	if tx == nil || tx.Transaction == nil || tx.Transaction.TransactionData == nil {
		return nil
	}

	return tx.Transaction.TraceContext().Span
}
