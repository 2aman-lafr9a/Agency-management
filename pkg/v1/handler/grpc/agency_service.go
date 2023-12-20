package grpc

import (
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	pb "Agency_management/proto"
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

func (srv *AgencyService) Create(req *pb.CreateAgencyRequest) (*pb.CreateAgencyResponse, error) {
	data := srv.transformAgencyRPC(req)
	if data.Name == "" {
		return &pb.CreateAgencyResponse{}, errors.New("name is required")
	}

	agency, err := srv.useCase.Create(data)
	if err != nil {
		return &pb.CreateAgencyResponse{}, err
	}

	return srv.transformAgencyModel(agency), nil
}

func (srv *AgencyService) Get(req *pb.GetAgenciesRequest) (*pb.GetAgenciesResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.GetAgenciesResponse{}, errors.New("id cannot be blank")
	}
	agency, err := srv.useCase.FindById(id)
	if err != nil {
		return &pb.GetAgenciesResponse{}, err
	}
	return (*pb.GetAgenciesResponse)(srv.transformAgencyModel(*agency)), nil
}

func (srv *AgencyService) transformAgencyRPC(req *pb.CreateAgencyRequest) *models.Agency {
	return &models.Agency{
		Name: req.GetName(),
	}
}

func (srv *AgencyService) transformAgencyModel(agency models.Agency) *pb.CreateAgencyResponse {
	return &pb.CreateAgencyResponse{
		Name: agency.Name,
	}
}
