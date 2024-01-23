package main

import (
	"errors"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
	"time"
)

// TargetPrice is a constant representing the target price for the book
const TargetPrice = 10.0

// URL is a constant that represents the URL of a book on Amazon.ca.
const URL = "https://www.amazon.ca/Shift-Silo-Trilogy-Book-2-ebook/dp/B088TDQG67/ref=sr_1_3?crid=AE1TQ4DWJMR8&keywords=silo&qid=1705970850&s=books&sprefix=silo%2Cstripbooks%2C135&sr=1-3"

// URL2 is a constant representing the URL of the second book on Amazon.ca.
const URL2 = "https://www.amazon.ca/gp/product/B088TCNVGJ?notRedirectToSDP=1&ref_=dbs_mng_calw_2&storeType=ebooks"

func main() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(12).Hours().Do(CheckPrice)
	scheduler.StartAsync()

	select {}
}

// CheckPrice is a function that checks the prices of books on Amazon and sends system notifications.
func CheckPrice() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36"))

	var price float64
	err := errors.New("")
	//debug query
	//c.OnHTML(".kindle-price #kindle-price-column #kindle-price", func(e *colly.HTMLElement) { log.Println("Found element: ", e.Text) })
	c.OnHTML(".kindle-price #kindle-price-column #kindle-price", func(e *colly.HTMLElement) {
		priceText := e.Text
		priceText = strings.TrimLeft(strings.TrimSpace(priceText), "$")
		price, err = strconv.ParseFloat(priceText, 64)
		if err != nil {
			log.Println("error in parsing the price:", err)
			return
		}
	})
	c.IgnoreRobotsTxt = true
	c.Visit(URL)
	c.Wait()
	if price < TargetPrice {
		notify(fmt.Sprintf("Price for the first book is now $%.2f, which is lower than $10", price))
	} else {
		notify(fmt.Sprintf("Price for the first book is now $%.2f, more than $10", price))
	}

	// Reset the price variable and check the price for the second book
	price = 0
	c.Visit(URL2)
	c.Wait()
	if price < TargetPrice {
		notify(fmt.Sprintf("Price for the second book is now $%.2f, which is lower than $10", price))
	} else {
		notify(fmt.Sprintf("Price for the second book is now $%.2f, more than $10", price))
	}
}

// notify is a function that sends a system notification with a given message.
func notify(message string) {
	err := beeep.Notify("Price Notification", message, "assets/information.png")
	if err != nil {
		panic(err)
	}
}
