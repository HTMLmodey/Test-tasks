package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Coin struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
}

func main() {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

	for {
		responseData := getData(url)
		var coins []Coin
		err := json.Unmarshal(responseData, &coins)
		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", err)
			return
		}

		fmt.Println("Доступные криптовалюты:")
		for _, coin := range coins {
			fmt.Printf("%s (%s)\n", coin.Name, coin.Symbol)
		}

		var selectedCoin string
		fmt.Print("Введите символ выбранной криптовалюты (например, 'btc' для Bitcoin): ")
		fmt.Scanln(&selectedCoin)

		var selectedCoinData Coin
		for _, coin := range coins {
			if coin.Symbol == selectedCoin {
				selectedCoinData = coin
				break
			}
		}

		if selectedCoinData.Symbol != "" {
			fmt.Printf("Текущий курс %s: $%.2f\n", selectedCoinData.Name, selectedCoinData.CurrentPrice)
		} else {
			fmt.Println("Криптовалюта с таким символом не найдена.")
		}

		time.Sleep(10 * time.Minute)
	}
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return nil
	}

	return body
}
