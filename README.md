# Tuple

# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-tuple
```

# 二、示例代码

```go
package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-tuple"
)

func main() {

	t2 := tuple.New2(1, "test")
	fmt.Println(t2)

}
```



