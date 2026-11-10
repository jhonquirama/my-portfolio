package config

type (
	service struct {
		Server   server   `json:"server" validate:"required"`
		Apm      apm      `json:"apm" validate:"required"`
		Services services `json:"services" validate:"required"`
		Aws      aws      `json:"aws" validate:"required"`
	}

	server struct {
		HTTP httpServer `json:"http" validate:"required"`
	}

	services struct {
		Portfolio Service `json:"portfolio" validate:"required"`
	}
)
