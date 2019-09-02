package master

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	// Game service
	gameService = newGameService()
	// ... other services
)

func init() {
	Services.Register(gameService)
}

func OnSessionClosed(s *session.Session) {
	gameService.userDisconnected(s)
}
