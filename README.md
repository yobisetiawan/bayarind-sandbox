# Bayarind Sandbox

This project is a simple Go (Golang) API for handling Virtual Account (VA) creation, using the Echo web framework.

## Requirements

- Go 1.23.2 or higher
- [Echo framework](https://echo.labstack.com/)

## Getting Started

### 1. Clone the repository

```sh
git clone <your-repo-url>
cd bayarind-sandbox
```

### 2. Install dependencies

```sh
go mod tidy
```

### 3. Run in Development

You can run the server locally with:

```sh
go run main.go
```

The server will start and listen for requests.

### 4. Run in Production

Build the binary and run it:

```sh
go build -o bayarind-sandbox
./bayarind-sandbox
```

Or use a process manager (e.g., systemd, supervisor, pm2) for better reliability in production.

## API Endpoint

- `POST /va` - Create a Virtual Account

### Request Headers
- `X-TIMESTAMP`
- `X-SIGNATURE`
- `X-PARTNER-ID`
- `X-EXTERNAL-ID`
- `CHANNEL-ID`

### Request Body Example
```json
{
  "partnerServiceId": "string",
  "customerNo": "string",
  "virtualAccountNo": "string",
  "virtualAccountName": "string",
  "virtualAccountEmail": "string",
  "trxId": "string",
  "totalAmount": { "value": "string", "currency": "string" },
  "billDetails": [
    { "billDescription": { "english": "string", "indonesia": "string" } }
  ],
  "expiredDate": "2025-08-04T15:04:05Z",
  "additionalInfo": { "insertId": 123 }
}
```

### Response Example
```json
{
  "responseCode": "2002700",
  "responseMessage": "Successful",
  "virtualAccountData": {
    "partnerServiceId": "string",
    "customerNo": "string",
    "virtualAccountNo": "string",
    "virtualAccountName": "string",
    "virtualAccountEmail": "string",
    "trxId": "string",
    "totalAmount": { "value": "string", "currency": "string" },
    "billDetails": [
      { "billDescription": { "english": "string", "indonesia": "string" } }
    ],
    "expiredDate": "2025-08-05T15:04:05Z",
    "additionalInfo": { "insertId": 123 }
  }
}
```
