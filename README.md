# Codename

> an [RFC1178](https://tools.ietf.org/html/rfc1178) implementation to generate pronounceable, sometimes even memorable, _"superhero like"_ codenames, consisting of a random combination of adjective and noun.


## Usage

Codename it's a [package](https://golang.org/doc/code#ImportingRemote), so all you need to do is import it into your code ([Try it!](https://play.golang.org/p/TrbW97r7aAO)):

```go
package main

import (
	"fmt"
	"github.com/mindfarm/codename"
)

func main() {
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 8; i++ {
		name := codename.Generate(rng, 0)
		fmt.Println(name)
	}
}
```

This is how the output looks like (since it's random your will be different).

```txt
absolutekaratecha
movingcolleen
gamenova
finemadrox
propenguin
keenmorbius
firmiron
refinedepoch
```

You can request the addition of a token to create even more entropy ([Try it!](https://play.golang.org/p/5gZTKfLyIUN)):

```go
package main

import (
	"fmt"
	"github.com/mindfarm/codename"
)

func main() {
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 8; i++ {
		name := codename.Generate(rng, 4)
		fmt.Println(name)
	}
}
```
