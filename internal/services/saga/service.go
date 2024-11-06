package saga

import (
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/shipping_addresses"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type service struct {
	kafkaConf                 *secret_proto.Kafka
	biteshipApiRepository     biteship_api.Repository
	shippingAddressRepository shipping_addresses.Repository
	kafkaBroker               ekafka.KafkaPubSub
	propagators               propagation.TextMapPropagator
	dbTransaction             wsqlx.Tx
}

func New(kafkaBroker ekafka.KafkaPubSub, kafkaConf *secret_proto.Kafka, tx wsqlx.Tx, biteshipApiRepository biteship_api.Repository, shippingAddressRepository shipping_addresses.Repository) *service {
	return &service{
		kafkaConf:                 kafkaConf,
		kafkaBroker:               kafkaBroker,
		propagators:               otel.GetTextMapPropagator(),
		dbTransaction:             tx,
		biteshipApiRepository:     biteshipApiRepository,
		shippingAddressRepository: shippingAddressRepository,
	}
}
