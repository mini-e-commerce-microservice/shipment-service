package handler

import (
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/address"
)

type handler struct {
	r                  *chi.Mux
	httpOtel           *whttp.Opentelemetry
	serv               serv
	jwtAccessTokenConf *secret_proto.JwtAccessToken
}

type serv struct {
	addressService address.Service
}

type Opt struct {
	JwtAccessTokenConf *secret_proto.JwtAccessToken
	AddressService     address.Service
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
			addressService: opt.AddressService,
		},
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/address", h.httpOtel.Trace(
		h.V1AddressPost,
	))
}
