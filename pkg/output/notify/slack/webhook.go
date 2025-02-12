package slack

type (
	Config interface {
		AppNameS() string
		WebhookURLS() string
		IconS() string
	}

	Slack struct {
		config Config
	}
)

func NewSlackClient(conf Config) Slack {
	return Slack{
		config: conf,
	}
}

func (s *Slack) PostWebhook(_ map[string]interface{}) {

}
