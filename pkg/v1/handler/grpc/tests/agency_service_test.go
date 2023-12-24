package tests

import (
	pb "Agency_management/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateAgency(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgencyServiceClient(conn)

	request := &pb.CreateAgencyRequest{
		Name: "Agency 1",
	}

	res, err := client.CreateAgency(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling CreateAgency: %s", err)
	}

	if res.Name != "Agency 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestGetAgency(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgencyServiceClient(conn)

	request := &pb.GetAgencyRequest{
		Name: "Agency 1",
	}

	res, err := client.GetAgency(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetAgency: %s", err)
	}

	if res.Name != "Agency 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestGetAgencies(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgencyServiceClient(conn)

	request := &pb.GetAgenciesRequest{}

	res, err := client.GetAgencies(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetAgencies: %s", err)
	}

	if len(res.Agencies) < 1 {
		t.Fatalf("Expected at least one agency, got %s", res.Agencies)
	}
}

func TestUpdateAgency(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgencyServiceClient(conn)

	request := &pb.UpdateAgencyRequest{
		Name: "Agency 1",
	}

	res, err := client.UpdateAgency(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling UpdateAgency: %s", err)
	}

	if res.Name != "Agency 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestDeleteAgency(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAgencyServiceClient(conn)

	request := &pb.DeleteAgencyRequest{
		Id: "1",
	}

	res, err := client.DeleteAgency(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling DeleteAgency: %s", err)
	}

	if res.Id != "1" {
		t.Fatalf("Expected id to be %s, got %s", request.Id, res.Id)
	}
}
