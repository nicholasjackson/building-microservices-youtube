#  Building Microservices in Go YouTube
Code repository for my Building Microservices YouTube series 
[https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)

Week by week Building Microservices builds on the previous weeks code teaching you how to build a multi-tier microservice system. The code structure for the course is one of a mono repo. To make it simple to follow along, each episode has its own branch showing progress to date.

## Services

### Product API [./product-api](./product-api)
Simple Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./frontend)
ReactJS website for presenting the Product API information

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


## Episode 9 - 9 CORS (Cross-Origin Resource Sharing)

### [https://youtu.be/RlYoy_RiYPw](https://youtu.be/RlYoy_RiYPw)

### Branch [episode_9](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_9)

In this episode we are going to take a look at CORS (Cross-Origin Resource Sharing). CORS is a security feature built into web browsers which restricts upstream requests to sites on different domains. We look at a typical example of a React website on one domain calling a back end API, see the impact of CORS and how to solve it.


## Episode 10 - 10 Serving and uploading files

### [https://youtu.be/ctmhYJpGsgU](https://youtu.be/ctmhYJpGsgU)

### Branch [episode_10](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_10)

In this episode you will learn how to upload and serve files using the Go standard library.


## Episode 11 - 11 Handling multi-part form uploads

### [https://youtu.be/_7-IhHMptNo](https://youtu.be/_7-IhHMptNo)

### Branch [episode_11](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_10)

In this episode you will learn how to handle multi-part form uploads. Mult-part forms used to be common place as they are the basic way that browsers would upload data to a server. This pattern has fallen out of fashion as most moder data transfer to the server is done using XHR requests. There might still be a case when you need to know this though.
