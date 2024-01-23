## Amazon Book Price Notifier
Amazon Book Price Notifier is a Go project developed to periodically check the price of specific books on Amazon. The application sends system notifications when the price drops below a pre-defined value.
## Features:
Periodically checks the price of specific books on Amazon
Uses web scraping to obtain current prices
Sends system notifications

## Dependencies
***go-co-op/gocron***: a Golang job scheduling package
***gocolly/colly***: a web scraping framework for Golang
***gen2brain/beeep***: a cross-platform library for sending desktop notifications and beeps

### Usage
go run main.go

## Disclaimer

This project is intended for educational purposes only. Web scraping can infringe on the terms of service of the website being scraped.