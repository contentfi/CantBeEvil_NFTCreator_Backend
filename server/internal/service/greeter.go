package service

import (
	"context"
	"time"

	v1 "server/api"
	"server/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) CreateCollection(ctx context.Context, in *v1.Collection) (*v1.CreateCollectionReply, error) {
	c, err := s.uc.Create(ctx, &biz.Collection{
		Name:         in.Name,
		Logo:         in.Logo,
		Desc:         in.Desc,
		License:      in.License,
		Address:      in.Address,
		Symbol:       in.Symbol,
		ChainID:      in.ChainId,
		Mtime:        time.Now(),
		OwnerAddress: in.CreatorAddress,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateCollectionReply{Success: true, Id: c.ID}, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) ListCollection(ctx context.Context, in *v1.ListCollectionRequest) (*v1.ListCollectionReply, error) {
	if in.Size == 0 {
		in.Size = 20
	}
	collections, err := s.uc.Find(ctx, in.AnchorId, int(in.Size), in.Reverse)
	if err != nil {
		return nil, err
	}
	reply := v1.ListCollectionReply{}
	for _, c := range collections {
		reply.Collections = append(reply.Collections, &v1.Collection{
			Id:             c.ID,
			Name:           c.Name,
			Logo:           c.Logo,
			Desc:           c.Desc,
			License:        c.License,
			Address:        c.Address,
			Symbol:         c.Symbol,
			ChainId:        c.ChainID,
			Mtime:          c.Mtime.Unix(),
			CreatorAddress: c.OwnerAddress,
		})
	}
	if len(reply.Collections) > 0 {
		reply.AnchorId = reply.Collections[len(reply.Collections)-1].Id
	}
	if len(reply.Collections) < int(in.Size) {
		reply.IsEnd = true
	}
	return &reply, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) UserCollection(ctx context.Context, in *v1.UserCollectionRequest) (*v1.UserCollectionReply, error) {
	collections, err := s.uc.UserCollections(ctx, in.OwnerAddress)
	if err != nil {
		return nil, err
	}
	reply := v1.UserCollectionReply{}
	for _, c := range collections {
		reply.Collections = append(reply.Collections, &v1.Collection{
			Id:             c.ID,
			Name:           c.Name,
			Logo:           c.Logo,
			Desc:           c.Desc,
			License:        c.License,
			Address:        c.Address,
			Symbol:         c.Symbol,
			ChainId:        c.ChainID,
			Mtime:          c.Mtime.Unix(),
			CreatorAddress: c.OwnerAddress,
		})
	}
	return &reply, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) GetCollection(ctx context.Context, in *v1.GetCollectionRequest) (*v1.Collection, error) {
	c, err := s.uc.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.Collection{
		Id:             c.ID,
		Name:           c.Name,
		Logo:           c.Logo,
		Desc:           c.Desc,
		License:        c.License,
		Address:        c.Address,
		Symbol:         c.Symbol,
		ChainId:        c.ChainID,
		Mtime:          c.Mtime.Unix(),
		CreatorAddress: c.OwnerAddress,
	}, nil
}
