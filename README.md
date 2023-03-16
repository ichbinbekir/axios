# Axios

## ⚡️ Quickstart

```go
package main

import "fmt"
import "github.com/ichbinbekir/axios"

func main() {

  response, err := axios.Get("http://example.com", axios.AxiosRequestConfig{Headers: map[string]string{"hello": "hello"}})
  if err != nil {
    panic(err)
  }
  
  fmt.Println(response.Data)
}
```
