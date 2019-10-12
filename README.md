# middlewares

[![Build Status](https://travis-ci.org/opentogo/middlewares.svg?branch=master)](https://travis-ci.org/opentogo/middlewares)
[![GoDoc](https://godoc.org/github.com/opentogo/middlewares?status.png)](https://godoc.org/github.com/opentogo/middlewares)
[![codecov](https://codecov.io/gh/opentogo/middlewares/branch/master/graph/badge.svg)](https://codecov.io/gh/opentogo/middlewares)
[![Go Report Card](https://goreportcard.com/badge/github.com/opentogo/middlewares)](https://goreportcard.com/report/github.com/opentogo/middlewares)
[![Open Source Helpers](https://www.codetriage.com/opentogo/middlewares/badges/users.svg)](https://www.codetriage.com/opentogo/middlewares)

## Usage

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/opentogo/middlewares"
)

func main() {
	var (
		mux    = http.NewServeMux()
		server = &http.Server{
			Addr:    ":3000",
			Handler: mux,
		}
	)

	middlewares.Use(
		middlewares.NewFrameOptions("SAMEORIGIN"),
		middlewares.NewIPSpoofing(),
		middlewares.NewOrigin([]string{"http://example.org"}),
		middlewares.NewPathTraversal(),
		middlewares.NewRemoteReferer([]string{http.MethodGet}),
		middlewares.NewStrictTransport(31536000, false, false),
		middlewares.NewXSS("block", true),
	)

	mux.HandleFunc("/hello-world", middlewares.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}))

	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
```

### FrameOptions middleware

It protects against clickjacking, setting header to tell the browser avoid embedding the page in a frame. You can define one option for this middleware.

Option | Description | Type
------ | ----------- | ----
option | Defines who should be allowed to embed the page in a frame. Use "DENY" or "SAMEORIGIN". | string

**Example:**

```go
middlewares.NewFrameOptions("SAMEORIGIN")
```

### IpSpoofing middleware

It detects IP spoofing attacks.

**Example:**

```go
middlewares.NewIPSpoofing()
```

### Origin middleware

It protects against unsafe HTTP requests when value of Origin HTTP request header doesn't match default or whitelisted URIs. You can define the whitelist of URIs.

Option | Description | Type
------ | ----------- | ----
whitelist | Array of allowed URIs | []string

**Example:**

```go
middlewares.NewOrigin([]string{"http://example.org"})
```

### PathTraversal middleware

It protects against unauthorized access to file system attacks, unescapes '/' and '.' from PATH_INFO.

**Example:**

```go
middlewares.NewPathTraversal()
```

### RemoteReferer middleware

It doesn't accept unsafe HTTP requests if the Referer header is set to a different host. You can define the HTTP methods that are allowed.

Option | Description | Type
------ | ----------- | ----
methods | Defines which HTTP method should be used. | []string

**Example:**

```go
middlewares.NewRemoteReferer([]string{
    http.MethodGet,
    http.MethodHead,
    http.MethodOptions,
    http.MethodTrace,
}),
```

### StrictTransport middleware

It protects against protocol downgrade attacks and cookie hijacking. You can define some options for this middleware.

Option | Description | Type
------ | ----------- | ----
max_age | How long future requests to the domain should go over HTTPS (in seconds). | int
include_subdomains | If all present and future subdomains will be HTTPS. | bool
preload | Allow this domain to be included in browsers HSTS preload list. | bool

**Example:**

```go
middlewares.NewStrictTransport(31536000, false, false)
```

### XSSHeader middleware

It sets X-XSS-Protection header to tell the browser to block attacks. XSS vulnerabilities enable an attacker to control the relationship between a user and a web site or web application that they trust.

You can define some options for this middleware.

Option | Description | Type
------ | ----------- | ----
xss_mode | How the browser should prevent the attack. | string
nosniff | Blocks a request if the requested type is "style" or "script". | bool

**Example:**

```go
middlewares.NewXSS("block", true)
```

## Contributors

- [rogeriozambon](https://github.com/rogeriozambon) Rogério Zambon - creator, maintainer
