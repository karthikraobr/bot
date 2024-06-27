package reponders

import (
	"context"
	"strings"
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
	if strings.Contains(message.Message, "Hi") {
		return "Hello again! We noticed you've recently received your iPhone 13. We'd love to hear about your experience. Can you spare a few minutes to share your thoughts?"
	} else if strings.Contains(message.Message, "Sure, I can do that.") {
		return "Fantastic! On a scale of 1-5, how would you rate the iPhone 13?"
	} else if strings.Contains(message.Message, "I'd give it a 5.") {
		return "Thank you for sharing your feedback! If you have any more thoughts or need assistance with anything else, feel free to reach out!"
	}
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
