package saga

import (
	"context"
)

func (s *service) OrderCourierRate(ctx context.Context) (err error) {
	//output, err := s.kafkaBroker.Subscribe(ctx, ekafka.SubInput{
	//	Config: kafka.ReaderConfig{
	//		Brokers: []string{s.kafkaConf.Host},
	//		GroupID: s.kafkaConf.Topic.OrderSaga.AggregateShipmentCourierRateRequest.ConsumerGroup.Shipmentsvc,
	//		Topic:   s.kafkaConf.Topic.OrderSaga.AggregateShipmentCourierRateRequest.Name,
	//	},
	//})
	//if err != nil {
	//	return collection.Err(err)
	//}
	//
	//for {
	//	data := models.OutboxEventCourierRate{}
	//	msg, err := output.Reader.FetchMessage(ctx, &data)
	//	if err != nil {
	//		return collection.Err(err)
	//	}
	//
	//	carrier := ekafka.NewMsgCarrier(&msg)
	//	ctxConsumer := s.propagators.Extract(context.Background(), carrier)
	//
	//	ctxConsumer, span := otel.Tracer("").Start(ctxConsumer, "order saga courier rate process.",
	//		trace.WithAttributes(
	//			attribute.Int64("cdc.debezium.payload.data.order_id", data.OrderID),
	//		))
	//
	//}
	return
}
