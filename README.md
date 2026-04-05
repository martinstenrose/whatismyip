# whatismyip

Simple HTTP service that returns the client's IP address as plain text.

## Usage

Supports both IPv4 and IPv6.

```
$ curl -4 http://localhost:8080/
192.0.2.1

$ curl -6 http://localhost:8080/
2001:db8::1
```

Respects the `X-Forwarded-For` header, so it works correctly behind a reverse proxy. It is recommended to run it behind a reverse proxy that handles SSL/TLS, such as [Caddy](https://caddyserver.com).

## Docker

```
docker pull ghcr.io/martinstenrose/whatismyip
docker run -p 8080:8080 ghcr.io/martinstenrose/whatismyip
```

Listens on port `8080`.

## License

GPL-3.0
