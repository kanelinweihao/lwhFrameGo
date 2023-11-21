package emailConnector

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"gopkg.in/gomail.v2"
)

type EntityEmailConnector struct {
	EmailSender       gomail.SendCloser
	EntityConfigEmail *conf.EntityConfigEmail
	EmailDialer       *gomail.Dialer
}
