package handler

import (
	"errors"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/address"
	"net/http"
)

func (h *handler) V1AddressPost(w http.ResponseWriter, r *http.Request) {
	userData, ok := h.getUserFromBearerAuth(w, r, false)
	if !ok {
		return
	}

	req := api.V1AddressPostJSONRequestBody{}
	if !h.httpOtel.BindBodyRequest(w, r, &req) {
		return
	}

	err := h.serv.addressService.CreateAddress(r.Context(), address.CreateAddressInput{
		ID:          req.Id,
		UserID:      userData.UserId,
		AddressNote: null.StringFromPtr(req.AddressNote),
	})
	if err != nil {
		switch {
		case errors.Is(err, address.ErrInvalidAddress):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, address.ErrInvalidAddress.Error())
		case errors.Is(err, address.ErrFromInterServerBiteshipApi):
			h.httpOtel.Err(w, r, http.StatusServiceUnavailable, err)
		default:
			h.httpOtel.Err(w, r, http.StatusInternalServerError, err)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
