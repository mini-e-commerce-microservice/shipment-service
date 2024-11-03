package biteship_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"net/http"
)

func (r *repository) GetAddress(ctx context.Context, input GetAddressInput) (output GetAddressOutput, err error) {
	req := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", r.conf.Token).
		SetQueryParam("countries", "ID")

	endpoint := "https://api.biteship.com/v1/maps/areas/" + input.ID

	resp, err := req.Get(endpoint)
	if err != nil {
		return output, collection.Err(err)
	}

	if resp.StatusCode() >= http.StatusInternalServerError {
		return output, collection.Err(fmt.Errorf("%w: %s", ErrFromInterServerBiteshipApi, resp.Error()))
	}

	if resp.StatusCode() == http.StatusBadRequest {
		return output, collection.Err(fmt.Errorf("%w: %s", ErrBadRequestBiteshipApi, resp.Error()))
	}

	if resp.IsError() {
		return output, collection.Err(errors.New(resp.String()))
	}

	err = json.Unmarshal(resp.Body(), &output)
	if err != nil {
		return output, collection.Err(err)
	}

	return
}

type GetAddressInput struct {
	ID string
}

type GetAddressOutput struct {
	Data []GetAddressesResponseItem `json:"areas"`
}
