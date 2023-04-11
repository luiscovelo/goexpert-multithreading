package dto

type AddressApi struct {
	ApiName string `json:"api_name"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type AddressApiError struct {
	ApiName string `json:"api_name"`
	Message string `json:"message"`
}

func NewAddressApi(apiName, city, state string) *AddressApi {
	return &AddressApi{
		ApiName: apiName,
		City:    city,
		State:   state,
	}
}

func NewAddressApiError(apiName string, message string) *AddressApiError {
	return &AddressApiError{
		ApiName: apiName,
		Message: message,
	}
}
