package courier_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/sha_256_payload"
	"google.golang.org/protobuf/proto"
	"testing"
)

func Test_service_CourierRates(t *testing.T) {
	h := "54d5a62be0f46f595d42a71003e9b41a08b677e8fe2690cdfa5073ca72ee1d54"
	payload := &sha_256_payload.CourierRate{
		ProductItem: []*sha_256_payload.CourierRateProductItem{
			{
				Length:   10,
				Width:    10,
				Height:   10,
				Weight:   10,
				Quantity: 2,
				Price:    10000,
			},
		},
		AvailableForCashOnDelivery:   true,
		AvailableForProofOfDelivery:  false,
		AvailableForInstantWaybillId: true,
		AvailableForInsurance:        true,
		Company:                      "jne",
		CourierCode:                  "jne",
		CourierServiceCode:           "reg",
		ServiceType:                  "standard",
		CourierPrice:                 10000,
	}

	payloadShaMarshal, err := proto.Marshal(payload)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	combined := append([]byte("key-shippment-service-courier-rate"+"|"), payloadShaMarshal...)
	hash.Write(combined)
	hashedData := hash.Sum(nil)

	fmt.Println(hex.EncodeToString(hashedData) == h)

}
