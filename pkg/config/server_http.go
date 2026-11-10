package config

type httpServer struct {
	Address string `json:"address" validate:"required,gt=0"`
}
