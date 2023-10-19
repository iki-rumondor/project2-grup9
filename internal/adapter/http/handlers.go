package customHTTP

import (
	"github.com/iki-rumondor/project2-grup9/internal/application"
)

type Handlers struct {
	Service *application.Service
}

func NewHandler(service *application.Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}
