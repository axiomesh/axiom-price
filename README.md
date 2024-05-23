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
go build -o price-aggregator ./cmd/server
./price-aggregator
```

### Using the API
To query the price of specific cryptocurrencies:
```bash
GET http://localhost:23510/api/ticker/price?symbol=BTCUSDT,ETHUSDT
```

## Configuration

### Proxy Settings
use_proxy: Set to true to use a proxy for network requests.

proxy_url: The URL of the proxy server.

### Fetching Prices
fetch_seconds_duration: The interval (in seconds) for fetching the price list from third-party services.

### HTTP Service
http_port: The HTTP port on which the service will run.

### Private Key Configuration
private_key: The default private key used for updating prices on-chain.

env_private_key_name: The environment variable name that can be used to read the private key.

### Price Feed Contract

price_feeds_contract: The contract address for the price feed.

### Blockchain Settings

rpc_url: The RPC URL of the blockchain.

chain_id: The ID of the blockchain.

### Exchange Rate Configuration

axc_usd: The default exchange rate for AXC to USD.

### Token Configuration
#### Price Contract Tokens

List of tokens with their tickers, symbols, and contract addresses for querying prices.

aa_ticker_list:
- ticker: The token's ticker.
- symbol: The token's symbol in the price feed contract.
- address: The token's contract address.

#### Aggregation Tokens

List of tokens for aggregation with their tickers and aliases.

wallet_ticker_list:
- ticker: The token's ticker.
- alias: The alias used for the token.

## Contributing

Contributions of any kind are welcome! If you have suggestions for improvement or want to add new features, please share them with us through issues or pull requests.

## License

This project is licensed under the MIT License.