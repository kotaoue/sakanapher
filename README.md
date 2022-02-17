# sakanapher
If Mr.Sakanakun is Gopher.

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

## Flowchart
flowchart TD
  START-->A{Are you Mr.Sakanakun?}
  A-->|Yes|B{Are you Gopher?}
    B-->|Yes|replace to ギョ or Go
    B-->|No|replace to ギョ
  A-->|No|C{Are you Gopher?}
    B-->|Yes|replace to Go
    B-->|No|not replace
