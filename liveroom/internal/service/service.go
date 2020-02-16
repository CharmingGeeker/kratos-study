package service

import (
	"context"
	"fmt"

	pb "liveroom/api"
	"liveroom/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

func (s *Service) Create(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Create " + req.Name,
	}
	fmt.Printf("Create %s", req.Name)
	return
}

func (s *Service) Delete(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Delete " + req.Name,
	}
	fmt.Printf("Delete %s", req.Name)
	return
}

func (s *Service) Get(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Get " + req.Name,
	}
	fmt.Printf("Get %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
