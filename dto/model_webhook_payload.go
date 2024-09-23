package dto

type WebhookPayload struct {
	Name   string        `json:"name"`
	Active bool          `json:"active"`
	Events []string      `json:"events"`
	Config WebhookConfig `json:"config"`
}
