package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type BrasilCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	cep := flag.String("cep", "", "CEP a ser consultado")
	flag.Parse()

	if cep == nil || *cep == "" || len(*cep) != 8 {
		log.Fatal("CEP não fornecido. Use a flag -cep para fornecer um CEP válido.")
	}

	ch1 := make(chan BrasilCep)
	ch2 := make(chan ViaCep)
	ch3 := make(chan string) //canal criado apra tratar erros de digitação do CEP

	brasilCepAPIUrl := "https://brasilapi.com.br/api/cep/v1/" + *cep
	viaCepAPIUrl := "http://viacep.com.br/ws/" + *cep + "/json/"

	go func(url string) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Println("Error creating request for BrasilCep API")
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error fetching BrasilCep API")
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			ch3 <- "Erro ao obter da BrasilCep API, verifique os dados digitados. Erro: " + resp.Status
			return
		}

		var brasilCep BrasilCep
		if err := json.NewDecoder(resp.Body).Decode(&brasilCep); err != nil {
			log.Println("Error decoding BrasilCep API response")
			return
		}

		ch1 <- brasilCep
	}(brasilCepAPIUrl)

	go func(url string) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			log.Println("Error creating request for ViaCep API")
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("Error fetching ViaCep API")
			return
		}

		if resp.StatusCode != http.StatusOK {
			ch3 <- "Erro ao obter da ViaCep API, verifique os dados digitados. Erro: " + resp.Status
			return
		}

		defer resp.Body.Close()

		var viaCep ViaCep
		if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
			log.Println("Error decoding ViaCep API response")
			return
		}

		ch2 <- viaCep
	}(viaCepAPIUrl)

	select {
	case brasilCep := <-ch1:
		log.Println("Reposta recebida da BrasilCep!")
		encoder := json.NewEncoder(os.Stdout)
		encoder.Encode(brasilCep)
	case viaCep := <-ch2:
		log.Println("Reposta recebida da ViaCep!")
		encoder := json.NewEncoder(os.Stdout)
		encoder.Encode(viaCep)
	case erroCep := <-ch3:
		log.Println(erroCep)
	case <-time.After(time.Second * 1):
		log.Println("Timeout ao receber respostas das APIs")
	}
}
