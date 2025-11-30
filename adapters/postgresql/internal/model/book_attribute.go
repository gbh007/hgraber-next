package model

var BookAttributeTable = BookAttribute{baseTable: baseTable{name: "book_attributes"}}

type BookAttribute struct {
	baseTable
}

func (ba BookAttribute) WithPrefix(pf string) BookAttribute {
	return BookAttribute{
		baseTable: ba.withPrefix(pf),
	}
}

func (ba BookAttribute) ColumnBookID() string { return ba.prefix + "book_id" }
func (ba BookAttribute) ColumnAttr() string   { return ba.prefix + "attr" }
func (ba BookAttribute) ColumnValue() string  { return ba.prefix + "value" }
