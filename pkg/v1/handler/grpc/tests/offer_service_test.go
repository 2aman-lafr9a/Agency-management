package tests

import (
	pb "Agency_management/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateOffer(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOfferServiceClient(conn)

	request := &pb.CreateOfferRequest{
		Name: "Offer 1",
	}

	res, err := client.CreateOffer(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling CreateOffer: %s", err)
	}

	if res.Name != "Offer 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestGetOffer(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOfferServiceClient(conn)

	request := &pb.GetOfferRequest{
		Name: "Offer 1",
	}

	res, err := client.GetOffer(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetOffer: %s", err)
	}

	if res.Name != "Offer 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestGetOffers(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOfferServiceClient(conn)

	request := &pb.GetOffersRequest{}

	res, err := client.GetOffers(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetOffers: %s", err)
	}

	if res.Offers[0].Name != "Offer 1" {
		t.Fatalf("Expected name to be %s, got %s", request, res.Offers[0].Name)
	}
}

func TestUpdateOffer(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOfferServiceClient(conn)

	request := &pb.UpdateOfferRequest{
		Name: "Offer 1",
	}

	res, err := client.UpdateOffer(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling UpdateOffer: %s", err)
	}

	if res.Name != "Offer 1" {
		t.Fatalf("Expected name to be %s, got %s", request.Name, res.Name)
	}
}

func TestDeleteOffer(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOfferServiceClient(conn)

	request := &pb.DeleteOfferRequest{
		Id: "1",
	}

	res, err := client.DeleteOffer(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling DeleteOffer: %s", err)
	}

	if res.Id != "1" {
		t.Fatalf("Expected name to be %s, got %s", request.Id, res.Id)
	}
}
