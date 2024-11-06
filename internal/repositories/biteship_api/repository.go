package biteship_api

import (
	resty_otel "github.com/SyaibanAhmadRamadhan/resty-otel"
	"github.com/go-resty/resty/v2"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
)

type repository struct {
	client *resty.Client
	conf   *secret_proto.ShipmentServiceBiteshipApi
}

func New(conf *secret_proto.ShipmentServiceBiteshipApi) *repository {
	c := resty.New()
	resty_otel.New(c, resty_otel.WithTraceResponseBody())

	return &repository{
		client: c,
		conf:   conf,
	}
}
