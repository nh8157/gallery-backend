package metadata

type Key string

const (
	FileName          Key = "FileName"
	Model             Key = "Model"
	LensModel         Key = "LensModel"
	FocalLength       Key = "FocalLength"
	DateTime          Key = "DateTime"
	ApertureValue     Key = "ApertureValue"
	ISOSpeedRatings   Key = "ISOSpeedRatings"
	ShutterSpeedValue Key = "ShutterSpeedValue"
)

type Metadata struct {
	FileName          *string
	Model             *string
	LensModel         *string
	FocalLength       *string
	DateTime          *string
	ApertureValue     *string
	ISOSpeedRatings   *string
	ShutterSpeedValue *string
}

func (md *Metadata) ToMap() map[string]string {
	return map[string]string{
		"FileName":          *md.FileName,
		"Model":             *md.Model,
		"LensModel":         *md.LensModel,
		"FocalLength":       *md.FocalLength,
		"DateTime":          *md.DateTime,
		"ApertureValue":     *md.ApertureValue,
		"ISOSpeedRatings":   *md.ISOSpeedRatings,
		"ShutterSpeedValue": *md.ShutterSpeedValue,
	}
}

func (md *Metadata) IsNil() bool {
	return md.FileName == nil && md.Model == nil && md.LensModel == nil && md.FocalLength == nil && md.DateTime == nil && md.ApertureValue == nil && md.ISOSpeedRatings == nil && md.ShutterSpeedValue == nil
}
