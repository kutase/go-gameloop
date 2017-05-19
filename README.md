# gameloop
> :video_game: :arrows_counterclockwise: Golang Game Loop implementation

## Install

```bash
go get github.com/kutase/go-gameloop
```
## Example

```go
package main

import (
	"github.com/kutase/go-gameloop"
	"log"
)

func main() {
	gl := gameLoop.New(10, func(delta float64) {
		log.Println("tick:", delta)
	})

	gl.Start()

	// Stop Game Loop:
	// gl.Stop()

	// Don't stop main goroutine
	for {}
}
```
