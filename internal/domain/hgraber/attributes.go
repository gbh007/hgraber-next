package hgraber

type Attribute string

var AllAttributes = []Attribute{
	AttrTag,
	AttrAuthor,
	AttrCharacter,
	AttrLanguage,
	AttrCategory,
	AttrParody,
	AttrGroup,
}

const (
	AttrAuthor    Attribute = "author"
	AttrCategory  Attribute = "category"
	AttrCharacter Attribute = "character"
	AttrGroup     Attribute = "group"
	AttrLanguage  Attribute = "language"
	AttrParody    Attribute = "parody"
	AttrTag       Attribute = "tag"
)
