package elit

import "reflect"

// GenerateOption is elit generate options
type GenerateOption struct {
	Presets  map[string]PropertyEncoderFunc
	Encoders map[reflect.Kind]PropertyEncoderFunc
}

// NewGenerateOption .
func NewGenerateOption() *GenerateOption {
	return &GenerateOption{
		Presets: map[string]PropertyEncoderFunc{
			"geo": geoPointEncoder,
		},
		Encoders: map[reflect.Kind]PropertyEncoderFunc{},
	}
}
