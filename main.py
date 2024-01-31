import logging
import sched
import time

import requests
from bs4 import BeautifulSoup
from plyer import notification

logger = logging.getLogger(__name__)

# TargetPrice is a constant representing the target price for the book
TARGET_PRICE = 10.0

# URL is a constant that represents the URL of a book on Amazon.ca.
URL = "https://www.amazon.ca/Shift-Silo-Trilogy-Book-2-ebook/dp/B088TDQG67/ref=sr_1_3?crid=AE1TQ4DWJMR8&keywords=silo&qid=1705970850&s=books&sprefix=silo%2Cstripbooks%2C135&sr=1-3"

# URL2 is a constant representing the URL of the second book on Amazon.ca.
URL2 = "https://www.amazon.ca/gp/product/B088TCNVGJ?notRedirectToSDP=1&ref_=dbs_mng_calw_2&storeType=ebooks"

s = sched.scheduler(time.time, time.sleep)


# Notify is a function that sends a system notification with a given message.
def notify(message):
    notification.notify(
        title="Price Notification",
        message=message,
        app_icon="assets/information.png",
        timeout=10,
    )


# CheckPrice is a function that checks the prices of books on Amazon and sends system notifications.
def check_price(sc):
    headers = {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
        'accept-language': 'en-GB,en;q=0.9',
    }
    res = requests.get(URL, headers=headers)
    soup = BeautifulSoup(res.text, 'html.parser')
    # logger.error(f"Soup: {soup}")
    logger.error(f"Select: {soup.select('.kindle-price #kindle-price-column #kindle-price')[0].text.strip()[1:]}")
    price = float(soup.select('.kindle-price #kindle-price-column #kindle-price')[0].text.strip()[1:])
    if price < TARGET_PRICE:
        notify(f"Price for the first book is now ${price:.2f}, which is lower than $10")
    else:
        notify(f"Price for the first book is now ${price:.2f}, more than $10")

    res = requests.get(URL2, headers=headers)
    soup = BeautifulSoup(res.text, 'html.parser')
    # c.OnHTML(".kindle-price #kindle-price-column #kindle-price", func(e *colly.HTMLElement)
    logger.error(f"Select: {soup.select('.kindle-price #kindle-price-column #kindle-price')[0].text.strip()[1:]}")
    price = float(soup.select('.kindle-price #kindle-price-column #kindle-price')[0].text.strip()[1:])
    if price < TARGET_PRICE:
        notify(f"Price for the second book is now ${price:.2f}, which is lower than $10")
    else:
        notify(f"Price for the second book is now ${price:.2f}, more than $10")

    s.enter(7200, 1, check_price, (sc,))  # Check every 2 hours.


s.enter(0, 1, check_price, (s,))
s.run()
