package model

type ParseRequest struct {
	Url     string `form:"u" query:"u"`
	AsZip   bool   `form:"as-zip" query:"as-zip"`
	AsMulti bool   `form:"as-multi" query:"as-multi"`
}
