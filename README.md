<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-05 23:30:03
 * @LastEditTime: 2025-04-06 01:35:25
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /readini/README.md
-->
# readini

A simple Go library to read ini files.

Current ver is 0.1.0.

100% coverage tested.

## How to use

For example:

``` go
import (
    "fmt"

    "github.com/FunctionSir/readini"
)

func main() {
    conf, err := readini.LoadFromFile("some_file.conf")
    fmt.Println(conf[""]["KeyA"], conf["SectionA"]["KeyA"])
}
```

## Bug report

Report bugs in Issue is recommended.
