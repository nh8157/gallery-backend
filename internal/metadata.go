package metadata

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
		"Model":             *md.FileName,
		"LensModel":         *md.LensModel,
		"FocalLength":       *md.FocalLength,
		"DateTime":          *md.DateTime,
		"ApertureValue":     *md.ApertureValue,
		"ISOSpeedRatings":   *md.ISOSpeedRatings,
		"ShutterSpeedValue": *md.ShutterSpeedValue,
	}
}
