# Shopify Stock Monitor

Shopify Stock Monitor is a Go application designed to monitor the stock availability of products on a Shopify store and send notifications to Discord when products are back in stock.

## Features

- Periodically fetches product data from a Shopify store
- Compares the current stock status with the previous status
- Sends Discord webhooks when products are back in stock
- Supports customizable webhook messages with product details

## TODO
- Add database incase of crash
- Proxy support
- Discord webhook rate limit rety
- Pagnation support
- Cache avoidance
