package command

import (
	"boilerplate/lib"
	"boilerplate/lib/dto"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

// WebrtcService service-interface
type WebrtcService interface {
	Execute(info dto.CallInfo) error
}

type webrtcService struct {
	// TODO: repository import here when need validate..
}

func (s *webrtcService) Execute(info dto.CallInfo) error {
	log.Info(fmt.Sprintf("Calling info: %s", lib.ToJsonStr(info)))
	return nil
}

func NewCreateUserService() WebrtcService {
	return &webrtcService{}
}
