# elit
[![CircleCI](https://circleci.com/gh/fidelitywires/elit.svg?style=shield&circle-token=404a3db148e2ff6d7047b60c628f69b1e97d8077)](https://circleci.com/gh/fidelitywires/elit)

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

p, err := elit.Generate(Sample{}, nil)

t.Mappings["sample"] = elit.Type{
	Properties: p,
}
```
