module github.com/nicholasjackson/building-microservices-youtube/product-api

go 1.13

require (
	github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/loads v0.19.5 // indirect
	github.com/go-openapi/runtime v0.19.14
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.8
	github.com/go-openapi/validate v0.19.7
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/hashicorp/go-hclog v0.12.1
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/nicholasjackson/building-microservices-youtube/currency v0.0.0
	github.com/nicholasjackson/env v0.6.0
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.1 // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	google.golang.org/genproto v0.0.0-20200326112834-f447254575fd // indirect
	google.golang.org/grpc v1.28.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/nicholasjackson/building-microservices-youtube/currency v0.0.0 => ../currency
