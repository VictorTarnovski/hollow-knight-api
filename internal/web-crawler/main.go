package webcrawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Enemy struct {
	Name        string
	Description string
}

var baseURL = "https://hollowknight.fandom.com"
var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
var enemies = []Enemy{}

func SearchInfo() {
	c := colly.NewCollector()
	c.UserAgent = UserAgent

	c.OnHTML("td[width='20%'] > div[class='center'] > div[class='floatnone'] > a", FilterEnemies)
	c.Visit(fmt.Sprint(baseURL, "/wiki/Category:Enemies_(Hollow_Knight)"))

	fmt.Println(enemies)
}

func FilterEnemies(e *colly.HTMLElement) {
	c := colly.NewCollector()
	c.UserAgent = UserAgent

	enemy := Enemy{Name: e.Attr("title")}

	c.OnHTML("body", func(h *colly.HTMLElement) {
		h.ForEachWithBreak("td[class='quote-text']", func(i int, h *colly.HTMLElement) bool {
			textContent := strings.Split(h.Text, "\n")
			enemy.Description = fmt.Sprint(textContent[1], " ", textContent[2])
			return false
		})
	})

	c.Visit(fmt.Sprint(baseURL, e.Attr("href")))

	enemies = append(enemies, enemy)
}
