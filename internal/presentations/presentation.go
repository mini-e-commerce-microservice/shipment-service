package presentations

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/presentations/handler"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/address"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/courier"
	"net/http"
	"time"
)

type Presenter struct {
	Port               int
	JwtAccessTokenConf *secret_proto.JwtAccessToken
	AddressService     address.Service
	CourierService     courier.Service
}

func New(p *Presenter) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3002"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "X-User-Id", "X-Request-Id", "X-Correlation-Id", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	handler.Init(r, handler.Opt{
		JwtAccessTokenConf: p.JwtAccessTokenConf,
		AddressService:     p.AddressService,
		CourierService:     p.CourierService,
	})

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", p.Port),
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return s
}
