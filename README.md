# url-builder

A fun app for url-building

```
url1 := NewUrlBuilder().Host("codility.com").Https().GetUrl().Build()
fmt.Println(url1) // matches  expected output -> https://codility.com

url2 := NewUrlBuilder().Host("codility.com").Path("/test/hello/world").QueryParams(map[string]string{"key1": "value1"}).Https().GetUrl().Build()
fmt.Println(url2) // matches expected output -> https://codility.com/test/hello/world?key1=value1

url3 := NewUrlBuilder().Host("codility.com").Port(1337).GetUrl().Build()
fmt.Println(url3) // matches expected output -> http://codility.com:1337
```

