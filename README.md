table
=====

Modified from github.com/modood/table.

Installation
------------

```
$ go get github.com/404tk/table
```

Quick start
-----------

```go
package main

import (
	"github.com/404tk/table"
)

type House struct {
	Name  string `table:"Name"`
	Sigil string
	Motto string
}

func main() {
	hs := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	// Output to stdout
	table.Output(hs)

	// Or just return table string and then do something
	s := table.Table(hs)
	_ = s
}
```

output:
```
┌───────────┬──────────┬──────────────────┐
│ Name      │ Sigil    │ Motto            │
├───────────┼──────────┼──────────────────┤
│ Stark     │ direwolf │ Winter is coming │
│ Targaryen │ dragon   │ Fire and Blood   │
│ Lannister │ lion     │ Hear Me Roar     │
└───────────┴──────────┴──────────────────┘
```