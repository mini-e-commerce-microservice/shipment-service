package biteship_api

import "errors"

var ErrFromInterServerBiteshipApi = errors.New("internal server error")
var ErrBadRequestBiteshipApi = errors.New("bad request")
var ErrInvalidPostalCode = errors.New("invalid postal code")
var ErrMissingParameter = errors.New("missing parameter")
var ErrNoCourierAvailable = errors.New("no courier available")
var ErrFromBiteshipApi = errors.New("from biteship api error")
