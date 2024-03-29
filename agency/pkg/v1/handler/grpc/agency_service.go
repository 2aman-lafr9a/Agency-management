package grpc

import (
	"agency/internal/models"
	interfaces "agency/pkg/v1"
	pb "agency/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type AgencyService struct {
	useCase interfaces.AgencyInterface
	pb.UnimplementedAgencyServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	useGrpc := &AgencyService{useCase: usecase}
	pb.RegisterAgencyServer(grpcServer, useGrpc)
}

func (srv *AgencyService) CreateAgency(ctx context.Context, req *pb.CreateAgencyRequest) (*pb.CreateAgencyResponse, error) {
	data := srv.transformAgencyRPC(req)
	if data.Name == "" {
		return &pb.CreateAgencyResponse{}, errors.New("name is required")
	}
	if data.Plan == "" {
		return &pb.CreateAgencyResponse{}, errors.New("plan is required")
	}
	if data.Description == "" {
		return &pb.CreateAgencyResponse{}, errors.New("description is required")
	}
	agency, err := srv.useCase.Create(data)
	return srv.transformAgencyModel(agency), err
}

func (srv *AgencyService) GetAgency(ctx context.Context, req *pb.GetAgencyRequest) (*pb.GetAgencyResponse, error) {
	data := srv.transformAgencyRPCGet(req)
	if data.Name == "" {
		return &pb.GetAgencyResponse{}, errors.New("name is required")
	}
	agency, err := srv.useCase.FindByName(data.Name)
	if agency == nil {
		return &pb.GetAgencyResponse{}, errors.New("agency not found")
	}
	offers, err := srv.useCase.FindOffersByAgencyID(agency.ID)
	if offers == nil {
		return &pb.GetAgencyResponse{}, errors.New("offers not found")
	}
	for _, offer := range offers {
		offerModelToOffer(*offer)
		agency.Offers = append(agency.Offers, *offer)
	}
	return srv.transformAgencyModelGet(*agency), err
}

func (srv *AgencyService) GetAgencies(ctx context.Context, req *pb.GetAgenciesRequest) (*pb.GetAgenciesResponse, error) {
	agencies, err := srv.useCase.FindAll()
	if agencies == nil {
		return &pb.GetAgenciesResponse{}, errors.New("agencies not found")
	}
	for _, agency := range agencies {
		offers, _ := srv.useCase.FindOffersByAgencyID(agency.ID)
		if offers == nil {
			return &pb.GetAgenciesResponse{}, errors.New("offers not found")
		}
		for _, offer := range offers {
			offerModelToOffer(*offer)
			agency.Offers = append(agency.Offers, *offer)
		}
	}
	return srv.transformAgenciesModel(agencies), err
}

func (srv *AgencyService) UpdateAgency(ctx context.Context, req *pb.UpdateAgencyRequest) (*pb.UpdateAgencyResponse, error) {
	data := srv.transformAgencyRPCUpdate(req)
	if data.ID == "" {
		return &pb.UpdateAgencyResponse{}, errors.New("name is required")
	}
	agency, err := srv.useCase.FindById(data.ID)
	agency.Name = data.Name
	agency.Plan = data.Plan
	agency.Description = data.Description
	agency.WalletAddress = data.WalletAddress
	_ = srv.useCase.Update(agency)
	return srv.transformAgencyModelUpdate(*agency), err
}

func (srv *AgencyService) DeleteAgency(ctx context.Context, req *pb.DeleteAgencyRequest) (*pb.DeleteAgencyResponse, error) {
	data := srv.transformAgencyRPCDelete(req)
	if data.ID == "" {
		return &pb.DeleteAgencyResponse{}, errors.New("id is required")
	}
	agency, _ := srv.useCase.FindById(data.ID)
	_ = srv.useCase.Delete(agency)
	return srv.transformAgencyModelDelete(*data), nil
}

func (srv *AgencyService) GetAgencyByWallet(ctx context.Context, req *pb.GetAgencyByWalletRequest) (*pb.GetAgencyByWalletResponse, error) {
	data := srv.transformAgencyRPCFindAgencyByWalletAddress(req)
	if data.WalletAddress == "" {
		return &pb.GetAgencyByWalletResponse{}, errors.New("wallet address is required")
	}
	agency, err := srv.useCase.FindAgencyByWalletAddress(data.WalletAddress)
	return srv.transformAgencyModelFindAgencyByWalletAddress(*agency), err
}

func (srv *AgencyService) transformAgencyRPCFindAgencyByWalletAddress(req *pb.GetAgencyByWalletRequest) *models.Agency {
	return &models.Agency{
		WalletAddress: req.Wallet,
	}
}

func (srv *AgencyService) transformAgencyModelFindAgencyByWalletAddress(agency models.Agency) *pb.GetAgencyByWalletResponse {
	offerItem := &pb.OfferItem{}
	if agency.Offers != nil {
		for _, offer := range agency.Offers {
			offerItem = offerToOfferItem(offer)
		}
	}
	if agency.Offers == nil {
		return &pb.GetAgencyByWalletResponse{
			Id:          agency.ID,
			Name:        agency.Name,
			Description: agency.Description,
			Plan:        agency.Plan,
			Offers:      nil,
			Wallet:      agency.WalletAddress,
		}
	}
	return &pb.GetAgencyByWalletResponse{
		Id:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
		Plan:        agency.Plan,
		Offers:      []*pb.OfferItem{offerItem},
		Wallet:      agency.WalletAddress,
	}
}

func (srv *AgencyService) transformAgencyRPC(req *pb.CreateAgencyRequest) *models.Agency {
	return &models.Agency{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Plan:          req.GetPlan(),
		WalletAddress: req.GetWallet(),
	}
}

func (srv *AgencyService) transformAgencyRPCGet(req *pb.GetAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetName(),
	}
}

func (srv *AgencyService) transformAgencyRPCUpdate(req *pb.UpdateAgencyRequest) *models.Agency {
	return &models.Agency{
		ID:            req.GetId(),
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Plan:          req.GetPlan(),
		WalletAddress: req.GetWallet(),
	}
}

func (srv *AgencyService) transformAgencyRPCDelete(req *pb.DeleteAgencyRequest) *models.Agency {
	return &models.Agency{
		ID: req.GetId(),
	}
}

func (srv *AgencyService) transformAgencyModel(agency models.Agency) *pb.CreateAgencyResponse {
	return &pb.CreateAgencyResponse{
		Name:        agency.Name,
		Description: agency.Description,
		Plan:        agency.Plan,
		Wallet:      agency.WalletAddress,
	}
}

func (srv *AgencyService) transformAgencyModelGet(agency models.Agency) *pb.GetAgencyResponse {
	if agency.Offers == nil {
		return &pb.GetAgencyResponse{
			Id:          agency.ID,
			Name:        agency.Name,
			Description: agency.Description,
			Plan:        agency.Plan,
			Offers:      nil,
			Wallet:      agency.WalletAddress,
		}
	}
	var offers []*pb.OfferItem
	for _, offer := range agency.Offers {
		offerToOfferItem(offer)
		offers = append(offers, offerToOfferItem(offer))
	}
	return &pb.GetAgencyResponse{
		Id:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
		Plan:        agency.Plan,
		Offers:      offers,
		Wallet:      agency.WalletAddress,
	}
}

func (srv *AgencyService) transformAgencyModelUpdate(agency models.Agency) *pb.UpdateAgencyResponse {
	if agency.Offers == nil {
		return &pb.UpdateAgencyResponse{
			Id:          agency.ID,
			Name:        agency.Name,
			Description: agency.Description,
			Plan:        agency.Plan,
			Offers:      nil,
			Wallet:      agency.WalletAddress,
		}
	}
	var offers []*pb.OfferItem
	for _, offer := range agency.Offers {
		offerToOfferItem(offer)
		offers = append(offers, offerToOfferItem(offer))
	}
	return &pb.UpdateAgencyResponse{
		Id:          agency.ID,
		Name:        agency.Name,
		Description: agency.Description,
		Plan:        agency.Plan,
		Offers:      offers,
		Wallet:      agency.WalletAddress,
	}
}

func (srv *AgencyService) transformAgenciesModel(agencies []*models.Agency) *pb.GetAgenciesResponse {
	var agenciesRPC []*pb.AgencyItem
	for _, agency := range agencies {
		if agency.Offers == nil {
			agenciesRPC = append(agenciesRPC, &pb.AgencyItem{
				Id:          agency.ID,
				Name:        agency.Name,
				Description: agency.Description,
				Plan:        agency.Plan,
				Offers:      nil,
				Wallet:      agency.WalletAddress,
			})
			continue
		}
		var offers []*pb.OfferItem
		for _, offer := range agency.Offers {
			offerToOfferItem(offer)
			offers = append(offers, offerToOfferItem(offer))
		}
		agenciesRPC = append(agenciesRPC, &pb.AgencyItem{
			Id:          agency.ID,
			Name:        agency.Name,
			Description: agency.Description,
			Plan:        agency.Plan,
			Offers:      offers,
			Wallet:      agency.WalletAddress,
		})
	}
	return &pb.GetAgenciesResponse{
		Agencies: agenciesRPC,
	}
}

func (srv *AgencyService) transformAgencyModelDelete(agency models.Agency) *pb.DeleteAgencyResponse {
	if agency.Offers == nil {
		return &pb.DeleteAgencyResponse{
			Id:          agency.ID,
			Description: agency.Description,
			Plan:        agency.Plan,
			Name:        agency.Name,
			Offers:      nil,
			Wallet:      agency.WalletAddress,
		}
	}
	var offers []*pb.OfferItem
	for _, offer := range agency.Offers {
		offerToOfferItem(offer)
		offers = append(offers, offerToOfferItem(offer))
	}
	return &pb.DeleteAgencyResponse{
		Id:          agency.ID,
		Description: agency.Description,
		Plan:        agency.Plan,
		Name:        agency.Name,
		Offers:      offers,
		Wallet:      agency.WalletAddress,
	}
}

func offerToOfferItem(offer models.Offer) *pb.OfferItem {
	return &pb.OfferItem{
		Id:          offer.ID,
		Name:        offer.Name,
		Description: offer.Description,
		Price:       int32(offer.Price),
		Date:        offer.Date,
		Rating:      offer.Rating,
		Type:        pb.OfferType(offer.OfferType),
	}
}

func offerModelToOffer(offer models.Offer) *models.Offer {
	return &models.Offer{
		ID:          offer.ID,
		Name:        offer.Name,
		Description: offer.Description,
		Price:       offer.Price,
		Date:        offer.Date,
		Rating:      offer.Rating,
		OfferType:   offer.OfferType,
		AgencyID:    offer.AgencyID,
	}
}
