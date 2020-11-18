package telegram

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Param set the parameter when calling bot's methods.
type Param func(params url.Values)

// resolveParam returns url.Values generated from params.
func resolveParam(params []Param) url.Values {
	urlVal := url.Values{}
	for _, param := range params {
		param(urlVal)
	}
	return urlVal
}

func setParamInt(field string, v int) Param {
	return func(params url.Values) {
		params.Set(field, strconv.Itoa(v))
	}
}

func setParamString(field string, v string) Param {
	return func(params url.Values) {
		params.Set(field, v)
	}
}

func setParamBool(field string, v bool) Param {
	return func(params url.Values) {
		params.Set(field, strconv.FormatBool(v))
	}
}

// SetOffset sets offset param.
func SetOffset(offset int) Param {
	return setParamInt("offset", offset)
}

// SetLimit sets limit param.
func SetLimit(limit int) Param {
	return setParamInt("limit", limit)
}

// SetTimeout set timeout param.
func SetTimeout(timeout int) Param {
	return setParamInt("timeout", timeout)
}

// SetAllowedUpdates set allowed_updates param.
func SetAllowedUpdates(allowedUpdates ...string) Param {
	return func(params url.Values) {
		if len(allowedUpdates) == 0 {
			allowedUpdates = make([]string, 0)
		}

		buf, err := json.Marshal(allowedUpdates)
		if err != nil {
			return
		}

		params.Set("allowed_updates", string(buf))
	}
}

// SetIPAddress sets ip_address param.
func SetIPAddress(address string) Param {
	return setParamString("ip_address", address)
}

// SetMaxConnections sets max_connections param.
func SetMaxConnections(count int) Param {
	return setParamInt("max_connections", count)
}

// SetDropPendingUpdates sets drop_pending_updates param.
func SetDropPendingUpdates(b bool) Param {
	return setParamBool("drop_pending_updates", b)
}
