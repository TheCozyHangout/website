package application

import (
	"github.com/TheCozyHangout/website/backend/internal/core/ports"
)

type Application struct {
	Db ports.MySQLPort
}

func NewApplication(db ports.MySQLPort) *Application {
	return &Application{Db: db}
}
