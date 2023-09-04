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

Document
--------

-   `func Output(slice interface{})`

    formats slice of structs data and writes to standard output.(Using box drawing characters)

-   `func OutputA(slice interface{})`

    formats slice of structs data and writes to standard output.(Using standard ascii characters)

-   `func Table(slice interface{}) string`

    formats slice of structs data and returns the resulting string.(Using box drawing characters)

-   `func AsciiTable(slice interface{}) string`

    formats slice of structs data and returns the resulting string.(Using standard ascii characters)

-   compare [box drawing characters](http://unicode.org/charts/PDF/U2500.pdf) with [standard ascii characters](https://ascii.cl/)

    box drawing:
    ```
    ┌───────────┬──────────┬──────────────────┐
    │ Name      │ Sigil    │ Motto            │
    ├───────────┼──────────┼──────────────────┤
    │ Stark     │ direwolf │ Winter is coming │
    │ Targaryen │ dragon   │ Fire and Blood   │
    │ Lannister │ lion     │ Hear Me Roar     │
    └───────────┴──────────┴──────────────────┘
    ```

    standard ascii:

    ```
    +-----------+----------+------------------+
    | Name      | Sigil    | Motto            |
    +-----------+----------+------------------+
    | Stark     | direwolf | Winter is coming |
    | Targaryen | dragon   | Fire and Blood   |
    | Lannister | lion     | Hear Me Roar     |
    +-----------+----------+------------------+
    ```


Contributing
------------

1.  Fork it
2.  Create your feature branch (`git checkout -b my-new-feature`)
3.  Commit your changes (`git commit -am 'Add some feature'`)
4.  Push to the branch (`git push origin my-new-feature`)
5.  Create new Pull Request

License
-------

this repo is released under the [MIT License](http://www.opensource.org/licenses/MIT).
