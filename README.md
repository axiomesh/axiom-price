Price-Aggregator
-

Price-Aggregator is a Go project designed to fetch and aggregate real-time price information from multiple cryptocurrency exchanges. It allows users to query the latest prices of specific cryptocurrencies, including but not limited to BTC and ETH, through a simple API.

## Features

- Support for fetching real-time price information from multiple cryptocurrency exchanges like Binance and Bybit.
- Provides a RESTful API for easy querying of specific cryptocurrency prices.
- Proxy support for accessing exchange APIs.
- Flexible configuration management for quick switching between local and production environments.
- Automated processing and aggregation of price information from different data sources.

## Quick Start

### Prerequisites
Go 1.21 or higher

### Installation and Running
1. Clone the project locally:
```shell
git clone https://github.com/axiomesh/price-aggregator.git
cd price-aggregator
```

2. Edit the config.yml file to set the proxy and other configurations as needed.
3. Build and run the project:
```shell
go build -o price-aggregator .
./price-aggregator
```

### Using the API
To query the price of specific cryptocurrencies:
```bash
GET http://localhost:23510/api/ticker/price?symbol=BTCUSDT,ETHUSDT
```

## Configuration

The config.yml configuration file supports the following settings:

- useProxy: Whether to access the API through a proxy (true/false).
- proxyURL: The URL of the proxy server.
- httpPort: The port number on which the HTTP server listens.

## Contributing

Contributions of any kind are welcome! If you have suggestions for improvement or want to add new features, please share them with us through issues or pull requests.

## License

This project is licensed under the MIT License.