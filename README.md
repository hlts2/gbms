# gbms
gbms is Allocation-free Boyer-Moore String Search Algorithm

## Require

- go(>=1.8)

## Installation

```shell
go get github.com/hlts2/gbms
```

## Example

`gbms.Search` searches a pattern in a str and returns count of pattern.

````go
str := "gogophergogophergogogo"
pattern := "gopher"

// search pattern in str
c := bms.Search(str, pattern) // c is 2
```
