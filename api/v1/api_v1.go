package api

import "github.com/objectbox/objectbox-go/objectbox"

type ApiV1 struct {
	obx *objectbox.ObjectBox
}

func NewApiV1(obx *objectbox.ObjectBox) *ApiV1 {
	return &ApiV1{obx: obx}
}