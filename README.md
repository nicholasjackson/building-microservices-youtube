#  Building Microservices in Go YouTube
Code repository for my Building Microservices YouTube series 
[https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)

Week by week Building Microservices builds on the previous weeks code teaching you how to build a multi-tier microservice system. The code structure for the course is one of a mono repo. To make it simple to follow along, each episode has its own branch showing progress to date.

Get $100 dollars of Digital Ocean credits, valid for 90 days with my referal link. Not needed for the YouTube tutorials, but helps me pay for my own servers.

<a href="https://www.digitalocean.com/?refcode=c6dee99fad25&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge"><img src="https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%201.svg" alt="DigitalOcean Referral Badge" /></a>

https://m.do.co/c/c6dee99fad25

## Services

### Product API [./product-api](./product-api)
RESTful Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./frontend)
ReactJS website for presenting the Product API information.

### Currency [./currency](./currency)
gRPC service supporting simple Unaray and Bidirectional streaming methods.

### Product Images [./product-images](./product-images)
Go based image service supporting Gzipped content, multi-part forms and a RESTful 
approach for uploading and downloading images.

## Series Content

Over the weeks we will look at the following topics, teaching you everything you need to know regarding building microservices with the go programming language:

- Introduction to microservices
- RESTFul microservices
- gRPC microservices
- Packaging applications with Docker
- Testing microservice
- Continuous Delivery
- Observability
- Using Kubernetes
- Debugging
- Security
- Asynchronous microservices
- Caching
- Microservice reliability using a Service Mesh


## Episode 1 - Building a simple microservice

### [https://youtu.be/VzBGi_n65iU](https://youtu.be/VzBGi_n65iU)

### Branch: [episode_1](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_1)

In this episode I show you how to build the simplest service you can using the standard packages in the Go programming language. 


## Episode 2 - Building a simple microservice, continued

### [https://youtu.be/hodOppKJm5Y](https://youtu.be/hodOppKJm5Y)

### Branch: [episode_2](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_2)

In this episode we learn more about the standard library and look at how we can refactor last episodes example into a reusable microservice pattern.


## Episode 3 - RESTFul microservices

### [https://youtu.be/eBeqtmrvVpg](https://youtu.be/eBeqtmrvVpg)

### Branch: [episode_3](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_3)

In this episode we start to learn about RESTFul services and reading and writing data using the JSON format.


## Episode 4 - RESTful microservices

### [https://youtu.be/UZbHLVsjpF0](https://youtu.be/UZbHLVsjpF0)

### Branch [episode_4](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_4)

We continue to look at how you can implement RESTFul services with the Standard API


## Episode 5 - Gorilla toolkit

### [https://youtu.be/DD3JlT_u0DM](https://youtu.be/DD3JlT_u0DM)

### Branch [episode_5](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_5)

In this epsode we refactor our Standard library RESTful service and start to implement the Gorill toolkit for routing.


## Episode 6 - JSON Validation

### [https://youtu.be/gE8_-8KoOLc](https://youtu.be/gE8_-8KoOLc)

### Branch [episode_6](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_6)

In this episode we will look at the Go Validator package and how it can be used to validate JSON.


## Episode 7 - Documenting APIs with Swagger

### [https://youtu.be/07XhTqE-j8k](https://youtu.be/07XhTqE-j8k)

### Branch [episode_7](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_7)

This epsiode shows you how you can document the API we have been building over the last couple of weeks. As a bonus we will also look at how we can embed ReDoc to build a nice documentation API direct into our service.


## Episode 8 - Auto-generating HTTP client code from Swagger documentation

### [https://youtu.be/Zn4joNjqBFc](https://youtu.be/Zn4joNjqBFc)

### Branch [episode_8](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_8)

In this episode we look at how we can use the Swagger API documentation we created in the last episode and generate a Go client SDK. As it turns out I had a little bug in my code
hope you all find the process of debugging this and finding root cause useful too.


## Episode 9 - CORS (Cross-Origin Resource Sharing)

### [https://youtu.be/RlYoy_RiYPw](https://youtu.be/RlYoy_RiYPw)

### Branch [episode_9](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_9)

In this episode we are going to take a look at CORS (Cross-Origin Resource Sharing). CORS is a security feature built into web browsers which restricts upstream requests to sites on different domains. We look at a typical example of a React website on one domain calling a back end API, see the impact of CORS and how to solve it.


## Episode 10 - Serving and uploading files

### [https://youtu.be/ctmhYJpGsgU](https://youtu.be/ctmhYJpGsgU)

### Branch [episode_10](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_10)

In this episode you will learn how to upload and serve files using the Go standard library.


## Episode 11 - Handling multi-part form uploads

### [https://youtu.be/_7-IhHMptNo](https://youtu.be/_7-IhHMptNo)

### Branch [episode_11](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_11)

In this episode you will learn how to handle multi-part form uploads. Mult-part forms used to be common place as they are the basic way that browsers would upload data to a server. This pattern has fallen out of fashion as most moder data transfer to the server is done using XHR requests. There might still be a case when you need to know this though.


## Episode 12 - Using Gzip compression for HTTP responses

### [https://youtu.be/GtSg1H7SU5Y](https://youtu.be/GtSg1H7SU5Y)

### Branch [episode_12](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_12)

In this episode we walk through how to wrap a http.ResponseWriter to enable Gzip compression for HTTP responses.

HTTP Headers Accept-Encoding:
https://developer.mozilla.org/en-US/d...

HTTP ResponseWriter:
https://golang.org/pkg/net/http/#Resp...


## Episode 13 - Introduction to gRPC and Protocol Buffers

gRPC is a high performance framework for client server applications. It is designed to be cross platform and is an awesome alternative to RESTful services.

In this episode we take a quick look at gRPC and Protocol Buffers, and how you can use them to build a simple API. This is the first video in a series of content where we dig into gRPC services.

gRPC Framework:
https://grpc.io/

Protocol Buffers v3 Language Guide:
https://developers.google.com/protocol-buffers/docs/proto3

Protocol Buffers v3 Encoding format:
https://developers.google.com/protocol-buffers/docs/encoding

### [https://youtu.be/pMgty_RYIOc](https://youtu.be/pMgty_RYIOc)


### Branch [episode_13](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_13)


## Episode 14 - gRPC Client Connections

In this episode we take a quick look at how you can connect to gRPC services in Go.

Protocol Buffers Enum Specification:
https://developers.google.com/protocol-buffers/docs/reference/proto3-spec#enum_definition

gRPC Creating gRPC Clients:
https://grpc.io/docs/tutorials/basic/go/#client

### [https://youtu.be/oTBcd5J0VYU](https://youtu.be/oTBcd5J0VYU)

### Branch [episode_14](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_14)


## Episode 15 - Refactoring Part 1/3

This video is part 1 of 3 where we start to clean up the code base a little before continuing to develop our gRPC service. Refactoring is a natural part of software development, it is difficult to get things right first time all the time. Rather than spend too much time on the perfect solution I like to go with the flow and clean up at a later date.

As part of our refactoring we look at the encoding/xml and how it is very similar in use to encoding/json.

Encoding/XML:
https://golang.org/pkg/encoding/xml/

### [https://youtu.be/Vl88R9acq-Y](https://youtu.be/Vl88R9acq-Y)

### Branch [episode_15_1](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_15_1)


## Episode 15 - Refactoring Part 2/3

In this episode I continue to refactor the code base so far. These videos are really just intended to ensure that you are not completely confused when looking at the source code changes between episode 14 and episode 16.

Source Code:
https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_15_2

Encoding/XML:
https://golang.org/pkg/encoding/xml/

### [https://youtu.be/QBl8LZ0Rems](https://youtu.be/QBl8LZ0Rems)

### Branch [episode_15_2](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_15_2)

## Episode 15 - Refactoring Part 3/3

In this episode I finalize the refactoring for the code base.

### [https://youtu.be/ARvOyAsuFog](https://youtu.be/ARvOyAsuFog)

### Branch [episode_15_2](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_15_2)


## Episode 16 - gRPC Bi-directional streaming, part 1/2

In this video we start to look at gRPC bi-directional streaming

gRPC streaming allows you to independently receive streamed messages from the client and send a response to it. This episode looks at the basics of streaming API by adding an update to our currency service.

Server-side streaming:
https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc

Client-side streaming:
https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc

### [https://youtu.be/4ohwkWVgEZM](https://youtu.be/4ohwkWVgEZM)

### Branch [episode_16](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_16)


## Episode 17 - gRPC Bi-directional streaming, part 2/2

In this video we continue to look at gRPC bi-directional streaming.

You will see how to take the simple example in the first part and how it can be implemented into the Products API to support
independent client and server streams. The simple example allows a client in the Product API to subscribe for currency
rate changes in the Currency service. Whenever a rate changes the currency service broadcasts this change to all
interested subscribers. 

Server-side streaming:
https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc

Client-side streaming:
https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc

### [https://youtu.be/MT5tXSKa-KY](https://youtu.be/MT5tXSKa-KY)

### Branch [episode_17](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_17)
