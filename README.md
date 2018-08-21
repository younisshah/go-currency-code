## go-currency-code

`go-currency-code` can be used to list and get currency details, specifically `country`, `code`, `symbol`, `native symbol` etcetra
in *O(1)* time.

The Ruby script `currency.rb` in the `transformer` directory generates a SHA 256 for each currency using the currency name.

#### Installation

`go get github.com/younisshah/go-currency-code`


### Usage

#### All()

 ```go
 currency "github.com/younisshah/go-currency-code"
 
 currencyMap, err := currency.All()
 if err != nil {
 	// ...
 }
 
 // do something with currencyMap 
 
 ```
 #### FromCurrencyName
 
 ```go
currency "github.com/younisshah/go-currency-code"
 
curr, err := currency.FromCurrencyName("US Dollar")
if err != nil {
 	// ...
 }
 
 // ..
```
 
 #### WriteJSON
 
 ```go
currency "github.com/younisshah/go-currency-code"

err := currency.WriteJSON(2, os.Stdout)

/*

{
  "0042a5e8fd68482a2dbfc6bacd377f37d49713a5964fc7eac371a30f849a721e": {
    "code": "GBP",
    "decimal_digits": 2,
    "name": "British Pound Sterling",
    "name_plural": "British pounds sterling",
    "rounding": 0,
    "symbol": "£",
    "symbol_native": "£"
  },
  "00f4ece9b7f15c414bc4cdd8646af422061ae21efc8c69240570b667e59952a1": {
    "code": "GEL",
    "decimal_digits": 2,
    "name": "Georgian Lari",
    "name_plural": "Georgian laris",
    "rounding": 0,
    "symbol": "GEL",
    "symbol_native": "GEL"
  },
...
...
}
*/

```


#### Tests

```go
go test
```

#### License

MIT
