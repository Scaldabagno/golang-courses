# Section 2-3

## Strings

* import strings module
```go
import (
	"fmt"
	"strings"
	)
```

* Same assignment as `int`

* Reassignment with `=` instead of `:=`
```go
m1 := "my name"
m1 := "shsh" //error
m1 = "shsh"  //ok
m2 := "name"
```

* strings methods
```go
strings.Contains(m1, m2) // m1 cointains m2
strings.ReplaceAll(m1, "m", "NO")
strings.Split(m1, " ")
// else...
```