package api

import (
	"github.com/congruity7/moonshot-go/pkg/service"
	"github.com/sirupsen/logrus"
)

type Context struct {
	ds   *service.DatabaseService
	rs   *service.RedisService
	_log *logrus.Logger
}

func NewAPIContext(ds *service.DatabaseService, rs *service.RedisService, logger *logrus.Logger) *Context {
	return &Context{
		ds:   ds,
		rs:   rs,
		_log: logger,
	}
}
