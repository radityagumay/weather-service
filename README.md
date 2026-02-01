# Weather Service for Cloudflare Workers

This project implements a simple weather service using static JSON data. It is designed to be deployed on Cloudflare Workers.

## Project Structure

```
weather-service
├── cmd
│   └── worker
│       └── main.go          # Entry point for the Cloudflare worker
├── internal
│   ├── handlers
│   │   └── weather.go       # Handler for weather requests
│   ├── models
│   │   └── weather.go       # Weather data model
│   └── data
│       └── static.go        # Loads static weather data
├── static
│   └── weather.json         # Static JSON data for weather information
├── wrangler.toml            # Configuration for Cloudflare worker deployment
├── go.mod                   # Module definition and dependencies
├── go.sum                   # Dependency checksums
└── README.md                # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd weather-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Configure your Cloudflare account in `wrangler.toml`.

4. Deploy the worker:
   ```
   wrangler publish
   ```

## How to run the service locally?

```
go run cmd/worker/main.go
```

## Usage

Once deployed, the weather service can be accessed via the Cloudflare worker URL. It will return weather information based on the static JSON data provided in `static/weather.json`.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.