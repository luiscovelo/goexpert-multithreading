package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/luiscovelo/goexpert-multithreading/internal/dto"
	"github.com/luiscovelo/goexpert-multithreading/pkg/formatter"
	"github.com/luiscovelo/goexpert-multithreading/pkg/validator"
)

const VIA_CEP_SERVICE_NAME string = "Via Cep Service"

type viaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

type viaCepUsecase struct {
	url string
}

func NewViaCepUsecase() *viaCepUsecase {
	return &viaCepUsecase{
		url: "https://viacep.com.br",
	}
}

func (uc *viaCepUsecase) Execute(rawCep string) (*dto.AddressApi, *dto.AddressApiError) {
	cep := formatter.SanitalizeCep(rawCep)

	if !validator.Cep(cep) {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "CEP is invalid")
	}

	cep = formatter.Cep(cep)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	endpoint := fmt.Sprintf("%s/ws/%s/json", uc.url, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "error to get data from CEP")
	}

	var response viaCepOutput

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}

	if response.Erro {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "not found CEP")
	}

	address := dto.NewAddressApi(VIA_CEP_SERVICE_NAME, response.Localidade, response.Uf)
	return address, nil
}
