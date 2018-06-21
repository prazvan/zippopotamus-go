# ZipIT

Zipit is a package to get information about a place given a zipcode

## Usage

Using Zipit is easy:

```go
package main

import "github.com/bankrate/zippopotamus-go"

func main() {
    client := &http.Client{}
    resp, err := zipit.GetDetailsFor("90210", client) //Returns the full info about the zipcode, including the country and postal code
    place, err := zipit.GetPlaceFor("90210", client) //Returns the location associated with the zipcode
}
```

## Installlation

`go get github.com/bankrate/zippopotamus-go`

-- OR --

`dep ensure -add github.com/bankrate/zippopotamus-go`
