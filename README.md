go-haikunator [![Build Status](https://travis-ci.org/yelinaung/go-haikunator.svg?branch=master)](https://travis-ci.org/yelinaung/go-haikunator)
=============

Heroku-like memorable random name generator. Golang port of [haikunator](https://github.com/usmanbashir/haikunator).

```bash
sparkling-cherry
snowy-brook
bitter-darkness
```

View the [docs](https://godoc.org/github.com/yelinaung/go-haikunator). 

Example
-------
```go
package main

import (
  "fmt"
  "github.com/yelinaung/go-haikunator"
  "time"
)

func main() {
  haikunator := haikunator.New(time.Now().UTC().UnixNano())
  fmt.Println(haikunator.Haikunate())
}

```

Other Languages
-------
Haikunator is also available in other languages. Check them out:
- Node: https://github.com/AtroxDev/haikunatorjs
- Python: https://github.com/AtroxDev/haikunator
- Ruby: https://github.com/usmanbashir/haikunator


License
-------
MIT

