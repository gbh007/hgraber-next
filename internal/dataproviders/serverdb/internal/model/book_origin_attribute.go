package model

var BookOriginAttributeTable = BookOriginAttribute{baseTable: baseTable{name: "book_origin_attributes"}}

type BookOriginAttribute struct {
	baseTable
}

func (ba BookOriginAttribute) WithPrefix(pf string) BookOriginAttribute {
	return BookOriginAttribute{
		baseTable: ba.withPrefix(pf),
	}
}

func (ba BookOriginAttribute) ColumnBookID() string { return ba.prefix + "book_id" }
func (ba BookOriginAttribute) ColumnAttr() string   { return ba.prefix + "attr" }
func (ba BookOriginAttribute) ColumnValues() string { return ba.prefix + "values" }
