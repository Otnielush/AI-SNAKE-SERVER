package game

import (
	"fmt"
	"log"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"AI-SNAKE-SERVER/protocol"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
)

type MapService struct {
	component.Base
	group *nano.Group
}

func newMapService() *MapService {
	return &MapService{
		group: nano.NewGroup("all-users"),
	}
}

func (rs *MapService) JoinMap(s *session.Session, msg *protocol.JoinMapRequest) error {
	if err := s.Bind(msg.MasterUid); err != nil {
		return errors.Trace(err)
	}

	broadcast := &protocol.NewUserBroadcast{
		Content: fmt.Sprintf("User user join: %v", msg.Nickname),
	}
	if err := rs.group.Broadcast("onNewUser", broadcast); err != nil {
		return errors.Trace(err)
	}
	return rs.group.Add(s)
}

type SyncMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (rs *MapService) SyncMessage(s *session.Session, msg *SyncMessage) error {
	// Send an RPC to master server to stats
	if err := s.RPC("GameService.Stats", &protocol.MasterStats{Uid: s.UID()}); err != nil {
		return errors.Trace(err)
	}

	// Sync message to all members in this room
	return rs.group.Broadcast("onMessage", msg)
}

func (rs *MapService) userDisconnected(s *session.Session) {
	if err := rs.group.Leave(s); err != nil {
		log.Println("Remove user from group failed", s.UID(), err)
		return
	}
	log.Println("User session disconnected", s.UID())
}
