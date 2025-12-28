package model

func JoinPageAndFile(pages Page, files File) string {
	return files.NameAlter() + " ON " + pages.ColumnFileID() + " = " + files.ColumnID()
}

func JoinFileAndPage(files File, pages Page) string {
	return pages.NameAlter() + " ON " + pages.ColumnFileID() + " = " + files.ColumnID()
}

func JoinBookAndPage(books Book, pages Page) string {
	return pages.NameAlter() + " ON " + books.ColumnID() + " = " + pages.ColumnBookID()
}

func JoinPageAndBookAttribute(pages Page, ba BookAttribute) string {
	return ba.NameAlter() + " ON " + pages.ColumnBookID() + " = " + ba.ColumnBookID()
}

func JoinBookAttributePage(ba BookAttribute, pages Page) string {
	return pages.NameAlter() + " ON " + ba.ColumnBookID() + " = " + pages.ColumnBookID()
}
