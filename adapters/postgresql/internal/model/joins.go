package model

func JoinPageAndFile(pages Page, files File) string {
	return files.NameAlter() + " ON " + pages.ColumnFileID() + " = " + files.ColumnID()
}
