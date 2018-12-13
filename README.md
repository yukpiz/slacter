# Slacter

slacter is a simple slack client that implements ``io.Writer``.


# Installation

```
$ go get github.com/yukpiz/slacter
```


# Usage

```
package main

import (
	"log"

	"github.com/yukpiz/slacter"
)

func main() {
	writer := slacter.New(&slacter.Config{
		Token:    "{{YOUR_SLACK_AUTH_TOKEN}}",
		Channel:  "test-channel",
		UserName: "yukpiz",
		IconURL:  "https://i.gyazo.com/30ff8e8938efed0382400961f3c59304.jpg",
	})

	logger := log.New(writer, "", log.Ldate|log.Ltime)

	logger.Println("Hello!")
	logger.Println("こんにちは!")
	logger.Println("안녕하세요!")
	logger.Println("您好!")
}
```

Result

<img src="https://i.gyazo.com/31009ea9e0cbacfb63a5fdc95500f2aa.png"/>

How to generate slack's token.  
https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens  
