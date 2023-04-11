package usecase_test

import (
	"log"
	"testing"

	"github.com/luiscovelo/goexpert-multithreading/internal/usecase"
)

func TestShouldReturnAddressSuccessfulyFromViaCEPApi(t *testing.T) {
	uc := usecase.NewViaCepUsecase()
	address, err := uc.Execute("37540000")

	if err != nil {
		t.Errorf("fail to get address from api: %#v", err)
	}

	log.Printf("Api Name: %s", address.ApiName)
	log.Printf("City: %s", address.City)
	log.Printf("State: %s", address.State)
}

func TestShouldReturnAddressSuccessfulyFromApiCep(t *testing.T) {
	uc := usecase.NewApiCepUsecase()
	address, err := uc.Execute("37540000")

	if err != nil {
		t.Errorf("fail to get address from api: %#v", err)
	}

	log.Printf("Api Name: %s", address.ApiName)
	log.Printf("City: %s", address.City)
	log.Printf("State: %s", address.State)
}
