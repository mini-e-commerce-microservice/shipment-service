package biteship_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/guregu/null/v5"
	"net/http"
)

func (r *repository) GetAddresses(ctx context.Context, input GetAddressesInput) (output GetAddressesOutput, err error) {
	req := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", r.conf.Token).
		SetQueryParam("countries", "ID")

	endpoint := "https://api.biteship.com/v1/maps/areas"
	if input.AreaID.Valid {
		endpoint += "/" + input.AreaID.String
	} else {
		if input.Search.Valid {
			req = req.SetQueryParam("input", input.Search.String)
		}
	}

	resp, err := req.Get(endpoint)
	if err != nil {
		return output, collection.Err(err)
	}

	if resp.StatusCode() >= http.StatusInternalServerError {
		return output, collection.Err(fmt.Errorf("%w: %s", ErrFromInterServerBiteshipApi, resp.Error()))
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

type GetAddressesInput struct {
	AreaID null.String
	Search null.String
}

type GetAddressesOutput struct {
	Items []GetAddressesResponseItem `json:"areas"`
}

type GetAddressesResponseItem struct {
	ID                               string `json:"id"`
	Name                             string `json:"name"`
	CountryName                      string `json:"country_name"`
	CountryCode                      string `json:"country_code"`
	AdministrativeDivisionLevel1Name string `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2Name string `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3Name string `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3Type string `json:"administrative_division_level_3_type"`
}
