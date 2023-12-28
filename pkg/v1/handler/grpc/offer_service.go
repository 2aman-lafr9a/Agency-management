package grpc

import (
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	pb "Agency_management/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type OfferService struct {
	useCase interfaces.OfferInterface
	pb.UnimplementedOfferServiceServer
}

func NewServer2(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	useGrpc := &OfferService{useCase: usecase}
	pb.RegisterOfferServiceServer(grpcServer, useGrpc)
}

func (srv *OfferService) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest) (*pb.CreateOfferResponse, error) {
	data := srv.transformOfferRPC(req)
	if data.Name == "" {
		return &pb.CreateOfferResponse{}, errors.New("name is required")
	}
	offer, _ := srv.useCase.Create2(data)
	return (*pb.CreateOfferResponse)(srv.transformOfferModel(offer)), nil
}

func (srv *OfferService) GetOffer(ctx context.Context, req *pb.GetOfferRequest) (*pb.GetOfferResponse, error) {
	data := srv.transformOfferRPCGet(req)
	if data.Name == "" {
		return &pb.GetOfferResponse{}, errors.New("name is required")
	}
	offer, err := srv.useCase.FindByName2(data.Name)
	if offer == nil {
		return &pb.GetOfferResponse{}, errors.New("offer not found")
	}
	return (*pb.GetOfferResponse)(srv.transformOfferModel(*offer)), err
}

func (srv *OfferService) GetOffers(ctx context.Context, req *pb.GetOffersRequest) (*pb.GetOffersResponse, error) {
	offers, err := srv.useCase.FindAll2()
	if offers == nil {
		return &pb.GetOffersResponse{}, errors.New("offers not found")
	}
	return srv.transformOffersModel(offers), err
}

func (srv *OfferService) UpdateOffer(ctx context.Context, req *pb.UpdateOfferRequest) (*pb.UpdateOfferResponse, error) {
	data := srv.transformOfferRPCUpdate(req)
	if data.Name == "" {
		return &pb.UpdateOfferResponse{}, errors.New("name is required")
	}
	offer, _ := srv.useCase.FindByName2(data.Name)
	offer.Name = data.Name
	_ = srv.useCase.Update2(offer)
	return (*pb.UpdateOfferResponse)(srv.transformOfferModel(*offer)), nil
}

func (srv *OfferService) DeleteOffer(ctx context.Context, req *pb.DeleteOfferRequest) (*pb.DeleteOfferResponse, error) {
	data := srv.transformOfferRPCDelete(req)
	if data.ID == "" {
		return &pb.DeleteOfferResponse{}, errors.New("id is required")
	}
	offer, _ := srv.useCase.FindById2(data.ID)
	_ = srv.useCase.Delete2(offer)
	return srv.transformOfferModelDelete(*data), nil
}

func (srv *OfferService) transformOfferRPC(req *pb.CreateOfferRequest) *models.Offer {
	return &models.Offer{
		Name: req.GetName(),
	}
}

func (srv *OfferService) transformOfferRPCGet(req *pb.GetOfferRequest) *models.Offer {
	return &models.Offer{
		Name: req.Name,
	}
}

func (srv *OfferService) transformOfferRPCUpdate(req *pb.UpdateOfferRequest) *models.Offer {
	return &models.Offer{
		Name: req.GetName(),
	}
}

func (srv *OfferService) transformOfferRPCDelete(req *pb.DeleteOfferRequest) *models.Offer {
	return &models.Offer{
		ID: req.GetId(),
	}
}

func (srv *OfferService) transformOfferModel(offer models.Offer) *pb.CreateOfferResponse {
	return &pb.CreateOfferResponse{
		Id:   offer.ID,
		Name: offer.Name,
	}
}

func (srv *OfferService) transformOffersModel(offers []*models.Offer) *pb.GetOffersResponse {
	var offersRPC []*pb.Offer
	for _, offer := range offers {
		offersRPC = append(offersRPC, &pb.Offer{
			Id:   offer.ID,
			Name: offer.Name,
		})
	}
	return &pb.GetOffersResponse{
		Offers: offersRPC,
	}
}

func (srv *OfferService) transformOfferModelDelete(offer models.Offer) *pb.DeleteOfferResponse {
	return &pb.DeleteOfferResponse{
		Id: offer.ID,
	}
}
