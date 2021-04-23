# string example

## StringBuilder

~~~go
package main

import (
	"fmt"

	str "github.com/gakkiyomi/galang/string"
)

func main() {
	builder := str.String.NewStringBuilder("")
	builder.Append("x").Append("\n")
	builder.Append("xx").Append("\n")
	builder.Append("xxx").Append("\n")
	builder.Append("xxxx").Append("\n")
	builder.Append("xxxxx").Append("\n")
	fmt.Println(builder.ToString())
}
~~~