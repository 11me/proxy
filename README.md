# Super Simple Proxy Server
This is a simple application layer proxy server written in Go.
This proxy useful to use when dealing with CORS.

## Yet Another Proxy?
I wrote this proxy server when I've been developing an app with ionic
framework. The app has been making http requests to remote servers, but
browser was blocking those requests.

## Good To Know
The proxy DOES NOT DO any caching and It has been developed to test apps while
developing.

## Running Locally
Locally server starts on `localhost` by default. Send all requests to
`http://localhost:8888/example.com`

## Running on a server
If you want run proxy on a server it implies that server has public IP.
It picks IP env variable.
Just run:
```
IP=$SOME_IP go run main.go
```
