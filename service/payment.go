package service

import (
	"context"
	"grpc-unary-medium/payment/datastore"
	"grpc-unary-medium/payment/paymentpb"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PaymentService struct {
	paymentpb.UnimplementedPaymentServer
	PaymentData datastore.DataStore
}

// return interfase dari hasil code generation
func NewPaymentService() paymentpb.PaymentServer {
	return &PaymentService{
		PaymentData: datastore.NewDataStore(),
	}
}

// method implementasi dari interface hasil code generation
func (service *PaymentService) CreatePayment(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {

	/*
		Use case create payment
	*/

	log.Println("Received request ", req.GetOrderId(), req.GetPaymentMethod(), req.GetTotalAmount())

	// check jika order id ada di data store
	if service.PaymentData.Exist(req.GetOrderId()) {
		return nil, status.Error(6, "RECORD WITH ORDER ID ALREADY EXIST")
	}
	// insert ke data store
	service.PaymentData.Insert(req.GetOrderId(), req.GetPaymentMethod().String())

	generatedPaymentID, err := uuid.NewRandom()
	if err != nil {
		// jika error return 13 (grpc code error referense https://grpc.github.io/grpc/core/md_doc_statuscodes.html)
		return nil, status.Error(13, "INTERNAL ERROR")
	}
	generatedPaymentExpiredTime := time.Now().Add(5 * time.Hour)
	generatedPaymentLink := "/pay/goku/" + req.GetPaymentMethod().String() + "/" + generatedPaymentID.String()

	// return hasil use case seperti return method/fungsi biasa
	return &paymentpb.CreatePaymentResponse{
		PaymentId:   generatedPaymentID.String(),
		PaymentLink: generatedPaymentLink,
		Expired:     timestamppb.New(generatedPaymentExpiredTime),
	}, nil
}
