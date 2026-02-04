# sakanapher
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
