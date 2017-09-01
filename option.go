package elit

// GenerateOption is elit generate options
type GenerateOption struct {
	Presets map[string]PropertyEncoderFunc
}

// NewGenerateOption .
func NewGenerateOption() *GenerateOption {
	return &GenerateOption{
		Presets: map[string]PropertyEncoderFunc{
			"geo": geoPointEncoder,
		},
	}
}
