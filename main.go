package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Acao struct {
	Results []struct {
		ShortName               string  `json:"shortName"`
		Currency                string  `json:"currency"`
		LongName                string  `json:"longName"`
		RegularMarketChange     float64 `json:"regularMarketChange"`
		RegularMarketChangePct  float64 `json:"regularMarketChangePercent"`
		RegularMarketTime       string  `json:"regularMarketTime"`
		RegularMarketPrice      float64 `json:"regularMarketPrice"`
		RegularMarketDayHigh    float64 `json:"regularMarketDayHigh"`
		RegularMarketDayRange   string  `json:"regularMarketDayRange"`
		RegularMarketDayLow     float64 `json:"regularMarketDayLow"`
		RegularMarketVolume     float64 `json:"regularMarketVolume"`
		RegularMarketPrevClose  float64 `json:"regularMarketPreviousClose"`
		RegularMarketOpen       float64 `json:"regularMarketOpen"`
		FiftyTwoWeekRange       string  `json:"fiftyTwoWeekRange"`
		FiftyTwoWeekLow         float64 `json:"fiftyTwoWeekLow"`
		FiftyTwoWeekHigh        float64 `json:"fiftyTwoWeekHigh"`
		Symbol                  string  `json:"symbol"`
		PriceEarnings           float64 `json:"priceEarnings,omitempty"`
		EarningsPerShare        float64 `json:"earningsPerShare"`
	} `json:"results"`
	RequestedAt string `json:"requestedAt"`
	Took        string `json:"took"`
}


func main() {
	var searchMode int

	if len(os.Args) >= 2 {
		searchMode, _ = strconv.Atoi(os.Args[1])
	}

	if searchMode == 0 || searchMode > 2 {
		fmt.Println("Invalid search mode. Please provide 1 for a single stock or 2 for multiple stocks.")
		return
	}

	var symbols []string

	if(searchMode == 2){
		if len(os.Args) < 3 {
				fmt.Println("Please provide at least one stock symbol.")
				return
		}
		for i := 2; i < len(os.Args); i++ {
			symbols = append(symbols, os.Args[i])
		}
	}else{
		if len(os.Args) < 3 {
			fmt.Println("Please provide a stock symbol.")
			return
		}
		symbols = append(symbols, os.Args[2])
	}
	

	for _, symbol := range symbols {
		url := fmt.Sprintf("https://brapi.dev/api/quote/%s?token=vw6agk3Cve8RTAV43n9R8C", symbol)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			fmt.Printf("API não disponível para %s\n", symbol)
			continue
		}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	var acao Acao
	err = json.Unmarshal(body, &acao)
	if err != nil {
		panic(err)
	}
	results := acao.Results[0]
	fmt.Println("")
	fmt.Println("Informações da Ação:")
	fmt.Printf("Nome Curto: %s\n", results.ShortName)
	fmt.Printf("Nome Longo: %s\n", results.LongName)
	fmt.Printf("Símbolo: %s\n", results.Symbol)
	fmt.Printf("Moeda: %s\n", results.Currency)
	fmt.Printf("Preço de Mercado Regular: %.2f\n", results.RegularMarketPrice)
	fmt.Printf("Variação do Mercado Regular: %.2f\n", results.RegularMarketChange)
	fmt.Printf("Porcentagem de Variação do Mercado Regular: %.2f%%\n", results.RegularMarketChangePct)
	fmt.Printf("Volume de Mercado Regular: %.0f\n", results.RegularMarketVolume)
	fmt.Printf("Abertura do Mercado Regular: %.2f\n", results.RegularMarketOpen)
	fmt.Printf("Fechamento do Mercado Regular Anterior: %.2f\n", results.RegularMarketPrevClose)
	fmt.Printf("Alta do Dia do Mercado Regular: %.2f\n", results.RegularMarketDayHigh)
	fmt.Printf("Baixa do Dia do Mercado Regular: %.2f\n", results.RegularMarketDayLow)
	fmt.Printf("Faixa do Dia do Mercado Regular: %s\n", results.RegularMarketDayRange)
	fmt.Printf("Variação de 52 Semanas: %s\n", results.FiftyTwoWeekRange)
	fmt.Printf("Mínimo de 52 Semanas: %.2f\n", results.FiftyTwoWeekLow)
	fmt.Printf("Máximo de 52 Semanas: %.2f\n", results.FiftyTwoWeekHigh)
	fmt.Printf("Relação P/L: %.2f\n", results.PriceEarnings)
	fmt.Printf("Lucro por Ação: %.2f\n", results.EarningsPerShare)
	fmt.Println("")
}
}