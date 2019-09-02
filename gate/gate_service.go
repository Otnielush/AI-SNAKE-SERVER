package gate

import (
	"github.com/lonng/nano/component"
	"SNAKE/protocol"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
)

type BindService struct {
	component.Base
	nextGateUid int64
}

func newBindService() *BindService {
	return &BindService{}
}

type (
	LoginRequest struct {
		Nickname string `json:"nickname"`
	}
	LoginResponse struct {
		Code int `json:"code"`
	}
)

func (bs *BindService) Login(s *session.Session, msg *LoginRequest) error {
	bs.nextGateUid++
	uid := bs.nextGateUid
	request := &protocol.NewUserRequest{
		Nickname: msg.Nickname,
		GateUid:  uid,
	}
	if err := s.RPC("GameService.NewUser", request); err != nil {
		return errors.Trace(err)
	}
	return s.Response(&LoginResponse{})
}

func (bs *BindService) BindMapServer(s *session.Session, msg []byte) error {
	return errors.Errorf("not implement")
}
