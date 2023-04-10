# GoLang MySQL Error

> Library to handle MySQL errors

![Tests](https://github.com/verify-lab/mysqlerr/actions/workflows/tests.yml/badge.svg)

## Install

```sh
go get github.com/verify-lab/mysqlerr
```

## Example

```go
package example

import "github.com/verify-lab/mysqlerr"

...

fmt.Println(mysqlerr.IsMySQLError(err))
fmt.Println(mysqlerr.CheckCode(err, mysqlerr.ErrDupEntry))
fmt.Println(mysqlerr.CastToMySQLError(err))
fmt.Println(mysqlerr.CastToDuplicateEntryError(err))

```
