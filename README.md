# eflag-go
# Usage  

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
output:
```
[[8] [16] [1] [2] [4]]
[[2 1] [4 1] [8 1] [16 1] [4 2] [2 8] [16 2] [8 4] [4 16] [16 8]]
[[4 2 1] [8 2 1] [16 2 1] [8 4 1] [16 4 1] [16 8 1] [4 8 2] [16 8 2] [4 16 2] [16 8 4]]
[[8 4 2 1] [16 4 2 1] [16 8 2 1] [16 8 4 1] [4 2 16 8]]
[[16 8 4 2 1]]
[2 18 3 6 10 7 11 19 22 26 14 15 23 27 30 31]
15
[1 2 4 8]
```
