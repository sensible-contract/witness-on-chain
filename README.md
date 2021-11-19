
## WitnessOnChain golang implementation

[WitnessOnChain](https://witnessonchain.com) is a witness/oracle service for blockchain smart contract.

### Deployment

need set 3 environment variables.

* listen address `LISTEN=0.0.0.0:8000`
* private key P `PINT=4500086422777284614649185698080302158559082854283623085071286437702609134945108376498504441947235804413771973268603081955797477873992855707579719874199`
* private key Q `QINT=644767523354888926443294407216811496298090339710859618436800524594486048803150814665399137334438348620190687521269075965377327693448171249204212488583`

0. docker
```
$ docker-compose build
$ docker-compose up -d

```

1. normal web API，
```
$ go build -v -o witnessonchain main.go
$ export PINT=4500086422777284614649185698080302158559082854283623085071286437702609134945108376498504441947235804413771973268603081955797477873992855707579719874199
$ export QINT=644767523354888926443294407216811496298090339710859618436800524594486048803150814665399137334438348620190687521269075965377327693448171249204212488583
$ LISTEN=:8000 ./witnessonchain
```
2. aws lambda
```
$ GOOS=linux GOARCH=amd64 go build -v -o witnessonchain aws/main.go
$ zip main_aws.zip witnessonchain
```

3. tencent serverless cloud function
```
$ GOOS=linux GOARCH=amd64 go build -v -o witnessonchain scf/main.go
$ zip main_scf.zip witnessonchain
```

### 0. Info

Returns greeting messages and public keys of WitnessOnChain service.

#### GET /info

* HTTP 200 OK

| Keys              | Type   | Description                                                                         |
|-------------------|--------|-------------------------------------------------------------------------------------|
| message           | string | Greeting message                                                                    |
| contact           | string | Contact information                                                                 |
| public\_key/rabin | string | Rabin public key of WitnessOnChain, little endian hex string, represents an integer |

```
{
    "message":"Welcome to WitnessOnChain!",
    "contact":"hi@witnessonchain.com",
    "public_key":{
        "rabin":"ad7e1e8d6d29...1a73a343bb40"
    }
}
```

* HTTP 404 Not Found
* HTTP 500 Internal Server Error

### 1. Timestamp

Returns Unix timestamp and UTC time string.

#### GET /timestamp

* HTTP 200 OK

| Keys                         | Type   | Description                                                                         |
|------------------------------|--------|-------------------------------------------------------------------------------------|
| timestamp                    | int    | Unix timestamp in 10 digits                                                         |
| utc                          | string | UTC time string in format like "2021-06-23 16:13:07"                                |
| digest                       | string | Signed message in hex string: 4 bytes timestamp little endian represents an integer |
| signatures/rabin/public\_key | string | Rabin public key of WitnessOnChain, little endian hex string, represents an integer |
| signatures/rabin/signature   | string | Rabin signature of digest, little endian hex string, represents an integer          |
| signatures/rabin/padding     | string | Rabin signature padding, little endian hex string                                   |
```
{
    "timestamp":1626073135,
    "utc":"2021-07-12 06:58:55",
    "digest":"2fe8eb60",
    "signatures":{
        "rabin":{
            "public_key":"ad7e1e8d6d29...1a73a343bb40",
            "signature":"56de85da727b...c90eae369926",
            "padding":"00"
        }
    }
}
```
* HTTP 400 Bad Request
* HTTP 404 Not Found
* HTTP 500 Internal Server Error
