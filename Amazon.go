package main

import (
	"fmt"                         //formatted I/O
	"github.com/gocolly/colly"    //scraping framework
	"github.com/gofiber/fiber/v2" // Fiber Framework
)

// Create Struct to save the extract data
type item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {

	// Initiate New Fiber
	app := fiber.New()
	app.Get("/scrape", func(c *fiber.Ctx) error {
		var items []item

		// initialize a Collector

		collector := colly.NewCollector()
		collector.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		// C.Visit Links
		urlScrap := []string{
			"https://www.amazon.ae/deal/f84ab57e?moreDeals=a9ef4295&restrictions=n%3A11995876031&showVariations=false&pf_rd_r=3RY1GHSYKECEWCEY4PTV&pf_rd_t=Events&pf_rd_i=deals&pf_rd_p=96c2f2b9-16f1-406b-850d-87c5a87bcb79&pf_rd_s=slot-14&ref=dlx_deals_gd_dcl_img_1_f84ab57e_dt_sl14_79",
			"https://www.amazon.ae/deal/015ef9d8?moreDeals=a9ef4295&restrictions=n%3A11995893031&showVariations=false&pf_rd_r=3RY1GHSYKECEWCEY4PTV&pf_rd_t=Events&pf_rd_i=deals&pf_rd_p=96c2f2b9-16f1-406b-850d-87c5a87bcb79&pf_rd_s=slot-14&ref=dlx_deals_gd_dcl_img_2_015ef9d8_dt_sl14_79",
		}

		// Colly selector

		collector.OnHTML("div.a-section", func(h *colly.HTMLElement) {
			h.ForEach(".octopus-dlp-asin-info-section.octopus-softline-asin-info-section", func(_ int, h *colly.HTMLElement) {

				item := item{
					Name:  h.ChildAttr("a", "title"),
					Price: h.DOM.Find("div.a-row.octopus-dlp-price > span.a-price.octopus-widget-price > span:nth-child(2) > span.a-price-whole").Text(),
				}

				items = append(items, item)
				// Printout Items Name
				fmt.Println(items)
				fmt.Println("ProductName: ", item.Name)
				//	fmt.Println("Price:", price, "AED")

			})

		})
		// For loop C.Visit Links
		for _, urlScrap := range urlScrap {
			collector.Visit(urlScrap)
		}
		//we return the extracted data to the client by calling the c.JSON(...) method.
		return c.JSON(items)
	})

	// Listener Port 9000
	app.Listen(":9000")

}
