package http

import (
	"golang.org/x/net/context"
	"net/http"
)

// UI represents an http "user interface", that is, an API.
type UI struct {
	Ctx context.Context
	R   *http.Request
	W   http.ResponseWriter
}

// Webhook is the JSON representation for a webhook.
type Webhook struct {
	ID                string `json:"id"`
	URL               string `json:"url"`
	DestinationFilter string `json:"destination_filter"`
	MemoTypeFilter    string `json:"memo_type_filter"`
	MemoFilter        string `json:"memo_filter"`
}

// WebhookPage is the JSON representation for a page of webhook
type WebhookPage struct {
	Embedded struct {
		Records []Webhook `json:"records"`
	} `json:"_embedded"`
}
