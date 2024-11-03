package main

import (
	"context"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/conf"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/infra"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/presentations"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/shipping_addresses"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/address"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var restApi = &cobra.Command{
	Use:   "restApi",
	Short: "use restApi",
	Run: func(cmd *cobra.Command, args []string) {
		otelConf := conf.LoadOtelConf()
		appConf := conf.LoadAppConf()
		jwtConf := conf.LoadJwtConf()

		closeFnOtel := infra.NewOtel(otelConf, appConf.TracerName)
		pgdb, pgdbCloseFn := infra.NewPostgresql(appConf.DatabaseDsn)
		rdbms := wsqlx.NewRdbms(pgdb)

		//productItemRepository := product_items.New(rdbms)
		//orderRepository := orders.New(rdbms)
		//orderItemRepository := order_items.New(rdbms)
		//outboxEventRepository := outbox_events.New(rdbms)
		//sagaStateRepository := saga_states.New(rdbms)

		biteshipApiRepository := biteship_api.New(appConf.BiteshipApi)
		shippingAddressRepository := shipping_addresses.New(rdbms)
		addressService := address.New(address.Opt{
			BiteshipApiRepository:     biteshipApiRepository,
			ShippingAddressRepository: shippingAddressRepository,
		})

		server := presentations.New(&presentations.Presenter{
			Port:               int(appConf.AppPort),
			JwtAccessTokenConf: jwtConf.AccessToken,
			AddressService:     addressService,
		})
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Err(err).Msg("failed start serve")
				stop()
			}
		}()

		<-ctx.Done()
		log.Info().Msg("Received shutdown signal, shutting down server gracefully...")

		//time.Sleep(40 * time.Second)
		closeFnOtel(context.TODO())
		pgdbCloseFn(context.TODO())
		log.Info().Msg("Shutdown complete. Exiting.")
	},
}
