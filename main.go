package main

import (
	"fmt"
	"net"
	"strconv"
)

type IUrlBuilder interface {
	// Host Set host
	Host(host string) IUrlBuilder
	// Path Set path
	Path(path string) IUrlBuilder
	// Port Set port number, or leave as 0
	Port(int) IUrlBuilder
	// QueryParams Set params as a map of strings
	QueryParams(queryParams map[string]string) IUrlBuilder
	// Https Set scheme to "https", otherwise leave http
	Https() IUrlBuilder
	// GetUrl Returns Url object with defined fields
	GetUrl() *Url
}

type Url struct {
	scheme      string
	host        string
	port        int
	path        string
	queryParams string
}

type UrlBuilder struct {
	url Url
}

func (b *UrlBuilder) Host(host string) IUrlBuilder {
	b.url.host = host
	return b
}

func (b *UrlBuilder) Path(path string) IUrlBuilder {
	b.url.path = path
	return b
}

func (b *UrlBuilder) Port(port int) IUrlBuilder {
	b.url.port = port
	return b
}

func (b *UrlBuilder) QueryParams(queryParams map[string]string) IUrlBuilder {
	// Build the query string from the map of query parameters
	queryString := "?"
	for key, value := range queryParams {
		queryString += key + "=" + value + "&"
	}
	// Remove the trailing "&"
	queryString = queryString[:len(queryString)-1]

	b.url.queryParams = queryString
	return b
}

func (b *UrlBuilder) Https() IUrlBuilder {
	b.url.scheme = "https"
	return b
}

func (b *UrlBuilder) GetUrl() *Url {
	return &b.url
}

func NewUrlBuilder() IUrlBuilder {
	return &UrlBuilder{}
}

func (u *Url) Build() string {
	// Build the URL from the components in the Url struct and return it as a string
	if u.scheme == "" {
		u.scheme = "http"
	}
	if u.port != 0 {
		u.host = net.JoinHostPort(u.host, strconv.Itoa(u.port))
	}
	return fmt.Sprintf("%s://%s%s%s", u.scheme, u.host, u.path, u.queryParams)
}

func main() {
	url1 := NewUrlBuilder().Host("codility.com").Https().GetUrl().Build()
	fmt.Println(url1) // matches  expected output -> https://codility.com

	url2 := NewUrlBuilder().Host("codility.com").Path("/test/hello/world").QueryParams(map[string]string{"key1": "value1"}).Https().GetUrl().Build()
	fmt.Println(url2) // matches expected output -> https://codility.com/test/hello/world?key1=value1

	url3 := NewUrlBuilder().Host("codility.com").Port(1337).GetUrl().Build()
	fmt.Println(url3) // matches expected output -> http://codility.com:1337
}
