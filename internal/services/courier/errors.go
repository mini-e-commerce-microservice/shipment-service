package courier

import "errors"

var ErrFromBiteshipApi = errors.New("err biteship api")
var ErrInvalidAddress = errors.New("invalid address")
var ErrNoCourierAvailable = errors.New("no courier available")
