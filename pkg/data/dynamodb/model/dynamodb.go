package model

import "context"

type (
	UpdateItemInput struct {
		Key                   map[string]any
		AttributeValues       map[string]any
		AttributeNames        map[string]string
		UpdateExpression      *string
		ConditionalExpression *string
	}

	QueryCallInput struct {
		AttributeValues  map[string]any
		AttributeNames   map[string]string
		KeyExpression    *string
		FilterExpression *string
	}
)

type NopRateLimiter struct{}

func (r *NopRateLimiter) GetToken(_ context.Context, _ uint) (releaseToken func() error, err error) {
	return func() error { return nil }, nil
}

func (r *NopRateLimiter) AddTokens(uint) error {
	return nil
}
