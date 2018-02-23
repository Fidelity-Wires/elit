# elit
[![CircleCI](https://circleci.com/gh/FidelityWires/elit.svg?style=shield&circle-token=404a3db148e2ff6d7047b60c628f69b1e97d8077)](https://circleci.com/gh/FidelityWires/elit) [![Go Report Card](https://goreportcard.com/badge/github.com/fidelitywires/elit)](https://goreportcard.com/report/github.com/fidelitywires/elit)[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FFidelityWires%2Felit.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FFidelityWires%2Felit?ref=badge_shield) [![GoDoc](https://godoc.org/github.com/FidelityWires/elit?status.svg)](https://godoc.org/github.com/FidelityWires/elit)

Generate elasticsearch index template from golang source code


# Installation

```
$ go get github.com/fidelitywires/elit
```

# Usage

```go

type Human struct {
	Name string
	Age  int
}

t := elit.Template{
	Template: "sample_template_*",
	Settings: elit.Settings{
		NumberOfShards:   5,
		NumberOfReplicas: 1,
	},
}

p, err := elit.Generate(Human{}, nil)

t.Mappings["sample"] = elit.Type{
	Properties: p,
}
```

This is sample so see also [ExampleGenerate](elit_test.go) function.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FFidelityWires%2Felit.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FFidelityWires%2Felit?ref=badge_large)
