package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"offer/internal/models"
	interfaces "offer/pkg/v1"
	pb "offer/proto"
)

type OfferService struct {
	useCase interfaces.OfferInterface
	pb.UnimplementedOfferServer
}

func NewServer2(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	useGrpc := &OfferService{useCase: usecase}
	pb.RegisterOfferServer(grpcServer, useGrpc)
}

func (srv *OfferService) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest) (*pb.CreateOfferResponse, error) {
	data := srv.transformOfferRPC(req)
	_, err := srv.useCase.FindAgencyByID(data.AgencyID)
	if err != nil {
		return &pb.CreateOfferResponse{}, errors.New("agency not found")
	}
	if data.Name == "" {
		return &pb.CreateOfferResponse{}, errors.New("name is required")
	}
	offer, _ := srv.useCase.Create2(data)
	return srv.transformOfferModel(offer), nil
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
	return srv.transformOfferModelGet(*offer), err
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
	return (*pb.UpdateOfferResponse)(srv.transformOfferModelGet(*offer)), nil
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
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
		Date:        req.GetDate(),
		AgencyID:    req.GetAgencyID(),
		OfferType:   models.OfferType(req.GetType()),
	}
}

func (srv *OfferService) transformOfferRPCGet(req *pb.GetOfferRequest) *models.Offer {
	return &models.Offer{
		Name: req.Name,
	}
}

func (srv *OfferService) transformOfferRPCUpdate(req *pb.UpdateOfferRequest) *models.Offer {
	return &models.Offer{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
		Date:        req.GetDate(),
		Rating:      req.GetRating(),
		OfferType:   models.OfferType(req.GetType()),
	}
}

func (srv *OfferService) transformOfferRPCDelete(req *pb.DeleteOfferRequest) *models.Offer {
	return &models.Offer{
		ID: req.GetId(),
	}
}

func (srv *OfferService) transformOfferModel(offer models.Offer) *pb.CreateOfferResponse {
	agency, _ := srv.useCase.FindAgencyByID(offer.AgencyID)
	agencyItem := srv.transformAgencyModel(*agency)
	return &pb.CreateOfferResponse{
		Id:          offer.ID,
		Name:        offer.Name,
		Description: offer.Description,
		Price:       int32(offer.Price),
		Date:        offer.Date,
		Type:        pb.OfferType(offer.OfferType),
		Agency:      &agencyItem,
	}
}

func (srv *OfferService) transformOfferModelGet(offer models.Offer) *pb.GetOfferResponse {
	agency, _ := srv.useCase.FindAgencyByID(offer.AgencyID)
	agencyItem := srv.transformAgencyModel(*agency)
	return &pb.GetOfferResponse{
		Id:          offer.ID,
		Name:        offer.Name,
		Description: offer.Description,
		Price:       int32(offer.Price),
		Date:        offer.Date,
		Rating:      offer.Rating,
		Type:        pb.OfferType(offer.OfferType),
		Agency:      &agencyItem,
	}
}

func (srv *OfferService) transformOffersModel(offers []*models.Offer) *pb.GetOffersResponse {
	var offersRPC []*pb.OfferItem
	for _, offer := range offers {
		agency, _ := srv.useCase.FindAgencyByID(offer.AgencyID)
		agencyItem := srv.transformAgencyModel(*agency)
		offersRPC = append(offersRPC, &pb.OfferItem{
			Id:          offer.ID,
			Name:        offer.Name,
			Description: offer.Description,
			Price:       int32(offer.Price),
			Date:        offer.Date,
			Rating:      offer.Rating,
			Type:        pb.OfferType(offer.OfferType),
			Agency:      &agencyItem,
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

func (srv *OfferService) transformAgencyModel(agency models.Agency) pb.AgencyItem {
	return pb.AgencyItem{
		Id:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
		Plan:        agency.Plan,
	}
}
