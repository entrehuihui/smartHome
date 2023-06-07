package service

import (
	"test/service/myrpc/proto"
)

// Service .
type Service struct {
	// ##继承
	proto.UnimplementedDeviceServer
	proto.UnimplementedGroupServer
	proto.UnimplementedLinkageServer
	proto.UnimplementedLoginServer
	proto.UnimplementedTypesServer
	// ##继承
}

// NewService .
func NewService() *Service {
	s := new(Service)
	return s
}
