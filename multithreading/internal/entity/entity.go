package entity

type CepBrasilApi struct {
	Cep string `json:"cep"`
	City string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street string `json:"street"`
	Service string `json:"service"`
}