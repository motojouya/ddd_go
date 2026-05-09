package repository

import (
	"github.com/motojouya/ddd_go/pkg/local/model"
)

type ServerGetter interface {
	GetServer() (*model.Server, error)
}

type ServerGet struct{}

func NewServerGet() *ServerGet {
	return &ServerGet{}
}

var serverConf *model.Server

func (getter *ServerGet) GetServer() (*model.Server, error) {
	if serverConf == nil {
		var serverConfObj, err = GetEnv[model.Server]()
		if err != nil {
			return nil, err
		}

		serverConf = &serverConfObj
	}

	return serverConf, nil
}
