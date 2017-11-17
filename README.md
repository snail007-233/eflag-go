# eflag-go

```golang
package main

import (
	"fmt"
  "github.com/snail007/eflag-go"
)

func main() {
	e := eflag.NewEasyFlag(5)
	fmt.Println(e.zuhe(1))
	fmt.Println(e.zuhe(2))
	fmt.Println(e.zuhe(3))
	fmt.Println(e.zuhe(4))
	fmt.Println(e.zuhe(5))
	fmt.Println(e.Search([]int{2}))
	fmt.Println(e.Encode([]int{1, 2, 4, 8}))
	fmt.Println(e.Decode(15))
}
```
