package handler

import (
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/address"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/courier"
)

type handler struct {
	r                  *chi.Mux
	httpOtel           *whttp.Opentelemetry
	serv               serv
	jwtAccessTokenConf *secret_proto.JwtAccessToken
}

type serv struct {
	addressService address.Service
	courierService courier.Service
}

type Opt struct {
	JwtAccessTokenConf *secret_proto.JwtAccessToken
	AddressService     address.Service
	CourierService     courier.Service
}

func Init(r *chi.Mux, opt Opt) {
	h := &handler{
		r: r,
		httpOtel: whttp.NewOtel(
			whttp.WithRecoverMode(true),
			whttp.WithPropagator(),
			whttp.WithValidator(nil, nil),
		),
		jwtAccessTokenConf: opt.JwtAccessTokenConf,
		serv: serv{
			courierService: opt.CourierService,
			addressService: opt.AddressService,
		},
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/address", h.httpOtel.Trace(
		h.V1AddressPost,
	))

	h.r.Post("/v1/courier-rates", h.httpOtel.Trace(
		h.V1CourierRatesPost,
	))
}
