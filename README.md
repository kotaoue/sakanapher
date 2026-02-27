# sakanapher

[![Go](https://github.com/kotaoue/sakanapher/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kotaoue/sakanapher/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kotaoue/sakanapher/branch/main/graph/badge.svg)](https://codecov.io/gh/kotaoue/sakanapher)
[![License](https://img.shields.io/github/license/kotaoue/sakanapher)](https://github.com/kotaoue/sakanapher/blob/main/LICENSE)

If Mr.Sakanakun is Gopher.

## Overview

```mermaid
flowchart TD
  START-->A{Mr.Sakanakun?}
  A-->|Yes|B{Gopher?}
    B -->|Yes| BY[ギョ or Go]
    B -->|No|ギョ
  A-->|No|C{Gopher?}
    C -->|Yes|Go
    C -->|No| CN[not replace]
```

## Usage

```ShellSession
$ go run main.go -message="おはようございます。ごきげんよろしゅうございますか？"
おはようございます。ごきげんよろしゅうございますか？

$ go run main.go -attribute=gopher -message="おはようございます。ごきげんよろしゅうございますか？"
おはようGoざいます。GoきげんよろしゅうGoざいますか？

$ go run main.go -name=sakanakun -message="おはようございます。ごきげんよろしゅうございますか？"
おはようギョざいます。ギョきげんよろしゅうギョざいますか？

$ go run main.go -name=sakanakun -attribute=gopher -message="おはようございます。ごきげんよろしゅうございますか？"
おはようぎょざいます。Goきげんよろしゅうぎょざいますか？
or
おはようGoざいます。ぎょきげんよろしゅうGoざいますか？
```
