package dto

type WebHookRepo struct {
	Owner      string `json:"owner"`
	Repo       string `json:"repo"`
	Token      string `json:"token"`
	WebHookUrl string `json:"webhookUrl"`
	Secret     string `json:"secret"`
}
