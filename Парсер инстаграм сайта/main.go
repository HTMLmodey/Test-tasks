package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://hypeauditor.com/top-instagram-all-russia/"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Не удалось выполнить GET-запрос:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatal("Статус ответа не 200 OK")
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Не удалось загрузить HTML-контент:", err)
	}

	file, err := os.Create("list_of_tops.csv")
	if err != nil {
		log.Fatal("Не удалось создать файл:", err)
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	defer writer.Flush()

	doc.Find(".table .row").Each(func(i int, s *goquery.Selection) {
		var data []string
		s.Find(".row-cell").Each(func(j int, cell *goquery.Selection) {
			if j < 7 {
				text := strings.TrimSpace(cell.Text())
				data = append(data, text)
			}
		})
		if err := writer.Write(data); err != nil {
			log.Fatal("Ошибка записи данных в файл:", err)
		}
	})

	fmt.Println("Данные успешно записаны в файл list_of_tops.csv")
}
