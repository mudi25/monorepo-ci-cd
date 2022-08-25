package errors

import (
	"fmt"
)

type apiErrors struct {
	Status string
	Reason string
}

func (apiError *apiErrors) Error() string {
	if apiError == nil {
		return "api error is nil"
	}
	return fmt.Sprintf("%s: %s", apiError.Status, apiError.Reason)
}

// create new error
func New(status string, err error) *apiErrors {
	if err == nil {
		return &apiErrors{
			Status: status,
			Reason: "err is nil",
		}
	}
	v, ok := err.(*apiErrors)
	if ok || v != nil {
		return v
	}
	return &apiErrors{
		Status: status,
		Reason: err.Error(),
	}
}

// extract error
func FromError(err error) (string, string) {
	if err == nil {
		return "success", ""
	}
	v, ok := err.(*apiErrors)
	if !ok || v == nil {
		v = &apiErrors{Status: "unknown", Reason: err.Error()}
	}
	return v.Status, v.Reason
}
