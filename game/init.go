package game

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	mapService = newRoomService()
)

func init() {
	Services.Register(mapService)
}

func OnSessionClosed(s *session.Session) {
	mapService.userDisconnected(s)
}
