# Currency Service
The currency service is a gRPC service which provides up to date exchange rates and currency conversion capabilities.

## Building protos
To build the gRPC client and server interfaces, first install protoc:

### Linux
```shell
sudo apt install protobuf-compiler
```

### Mac
```shell
brew install protoc
```

Then install the Go gRPC plugin:

```shell
go get google.golang.org/grpc
```

Then run the build command:

```shell
protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
```

## Testing
To test the system install `grpccurl` which is a command line tool which can interact with gRPC API's

https://github.com/fullstorydev/grpcurl

```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```


### List Services
```
grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection
```

### List Methods
```
grpcurl --plaintext localhost:9092 list Currency        
Currency.GetRate
```

### Method detail for GetRate
```
grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );
```

### RateRequest detail
```
grpcurl --plaintext localhost:9092 describe .RateRequest    
RateRequest is a message:
message RateRequest {
  string Base = 1 [json_name = "base"];
  string Destination = 2 [json_name = "destination"];
}
```

### Execute a request
```
grpcurl --plaintext -d '{"base": "GBP", "destination": "USD"}' localhost:9092 Currency/GetRate
{
  "rate": 0.5
}
```