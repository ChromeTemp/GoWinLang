<!-- cspell:word Println -->

# GoWinLang [![CI](https://github.com/ChromeTemp/GoWinLang/actions/workflows/CI.yml/badge.svg)](https://github.com/ChromeTemp/GoWinLang/actions/workflows/CI.yml)
> Detect Windows user language in Go

## Installation

```bash
go get github.com/ChromeTemp/GoWinLang
```

## Usage

```go
package main

import "github.com/ChromeTemp/GoWinLang"

main() {
    language := GoWinLang.DetectLanguage()
    fmt.Println("Your main language: " + language)
    // Output: Your main language: en
}
```
