package reponders

import (
	"context"
	"time"
)

const (
	TypeReview  = "review"
	TypeSupport = "support"
	TypeDefault = "default"
)

type Message struct {
	Message string
	Time    time.Time
}

type ReviewResponder struct{}

func (r *ReviewResponder) Answer(ctx context.Context, message Message, messageContext []Message) string {
	return "I'm a review responder"
}
func (r *ReviewResponder) Type() string {
	return TypeReview
}

type SupportResponder struct{}

func (r *SupportResponder) Answer(ctx context.Context, message Message, messageContext []Message) string {
	return "I'm a support responder"
}

func (r *SupportResponder) Type() string {
	return TypeSupport
}

type DefaultResponder struct{}

func (d *DefaultResponder) Answer(ctx context.Context, message Message, messageContext []Message) string {
	return "I'm a default responder. Your ticket will be reviewed soon."
}

func (d *DefaultResponder) Type() string {
	return TypeDefault
}
