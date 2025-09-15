package outline

import "smart-edu-api/embeded"

type Request struct {
	Outline embeded.Outline `json:"outline" bson:"outline"`
}
