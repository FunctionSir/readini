<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-05 23:30:03
 * @LastEditTime: 2025-06-20 09:14:48
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /readini/README.md
-->
# readini

A simple Go library to read ini files.

Current ver is 0.3.1.

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
    sec := conf["SectionA"]
    fmt.Println(sec["KeyA"], sec.HasKey("KeyA"))
}
```

## Bug report

Report bugs in Issue is recommended.
