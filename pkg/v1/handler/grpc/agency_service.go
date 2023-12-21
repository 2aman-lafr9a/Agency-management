package grpc

import (
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	pb "Agency_management/proto"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type AgencyService struct {
	useCase interfaces.AgencyInterface
	pb.UnimplementedAgencyServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	useGrpc := &AgencyService{useCase: usecase}
	pb.RegisterAgencyServiceServer(grpcServer, useGrpc)
}

func (srv *AgencyService) CreateAgency(ctx context.Context, req *pb.CreateAgencyRequest) (*pb.CreateAgencyResponse, error) {
	data := srv.transformAgencyRPC(req)
	if data.Name == "" {
		return &pb.CreateAgencyResponse{}, errors.New("name is required")
	}
	agency, _ := srv.useCase.Create(data)
	return (*pb.CreateAgencyResponse)(srv.transformAgencyModel(agency)), nil
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
	return (*pb.GetAgencyResponse)(srv.transformAgencyModel(*agency)), err
}

func (srv *AgencyService) GetAgencies(ctx context.Context, req *pb.GetAgenciesRequest) (*pb.GetAgenciesResponse, error) {
	agencies, err := srv.useCase.FindAll()
	if agencies == nil {
		return &pb.GetAgenciesResponse{}, errors.New("agencies not found")
	}
	return srv.transformAgenciesModel(agencies), err
}

func (srv *AgencyService) UpdateAgency(ctx context.Context, req *pb.UpdateAgencyRequest) (*pb.UpdateAgencyResponse, error) {
	data := srv.transformAgencyRPCUpdate(req)
	if data.Name == "" {
		return &pb.UpdateAgencyResponse{}, errors.New("name is required")
	}
	agency, _ := srv.useCase.FindByName(data.Name)
	agency.Name = data.Name
	_ = srv.useCase.Update(agency)
	return (*pb.UpdateAgencyResponse)(srv.transformAgencyModel(*agency)), nil
}

func (srv *AgencyService) DeleteAgency(ctx context.Context, req *pb.DeleteAgencyRequest) (*pb.DeleteAgencyResponse, error) {
	data := srv.transformAgencyRPCDelete(req)
	if data.Name == "" {
		return &pb.DeleteAgencyResponse{}, errors.New("name is required")
	}
	return srv.transformAgencyModelDelete(*data), nil
}

func (srv *AgencyService) transformAgencyRPC(req *pb.CreateAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetName(),
	}
}

func (srv *AgencyService) transformAgencyRPCGet(req *pb.GetAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetName(),
	}
}

func (srv *AgencyService) transformAgencyRPCUpdate(req *pb.UpdateAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetName(),
	}
}

func (srv *AgencyService) transformAgencyRPCDelete(req *pb.DeleteAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetId(),
	}
}

func (srv *AgencyService) transformAgencyModel(agency models.Agency) *pb.CreateAgencyResponse {
	return &pb.CreateAgencyResponse{
		Id:   agency.ID,
		Name: agency.Name,
	}
}

func (srv *AgencyService) transformAgenciesModel(agencies []*models.Agency) *pb.GetAgenciesResponse {
	var agenciesRPC []*pb.Agency
	for _, agency := range agencies {
		agenciesRPC = append(agenciesRPC, &pb.Agency{
			Id:   agency.ID,
			Name: agency.Name,
		})
	}
	return &pb.GetAgenciesResponse{
		Agencies: agenciesRPC,
	}
}

func (srv *AgencyService) transformAgencyModelDelete(agency models.Agency) *pb.DeleteAgencyResponse {
	return &pb.DeleteAgencyResponse{
		Id: agency.ID,
	}
}
