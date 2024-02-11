package metadataReader

import (
	"io"
	"log"

	"github.com/nh8157/gallery-backend/internal/metadata"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func ReadExif(image io.ReadCloser, fileName *string) metadata.Metadata {
	exif.RegisterParsers(mknote.All...)
	rawExif, err := exif.Decode(image)
	if err != nil {
		log.Println(err)
	}

	metadata := metadata.Metadata{
		FileName:          fileName,
		Model:             getField(rawExif, exif.Model),
		LensModel:         getField(rawExif, exif.LensModel),
		FocalLength:       getField(rawExif, exif.FocalLength),
		DateTime:          getField(rawExif, exif.DateTime),
		ApertureValue:     getField(rawExif, exif.ApertureValue),
		ISOSpeedRatings:   getField(rawExif, exif.ISOSpeedRatings),
		ShutterSpeedValue: getField(rawExif, exif.ShutterSpeedValue),
	}

	return metadata
}

func getField(rawExif *exif.Exif, field exif.FieldName) *string {
	valueRaw, err := rawExif.Get(field)
	if err != nil {
		log.Println(err)
		return nil
	}
	value := valueRaw.String()
	return &value
}
