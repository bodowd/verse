package cmd

type Book struct {
	bookNumber     string
	fullName       string
	idAbbreviation string
}

func booksInfo() map[string]Book {
	return map[string]Book{
		"Gen":    {"01", "Genesis", "Gen"},
		"Exo":    {"02", "Exodus", "Exo"},
		"Lev":    {"03", "Leviticus", "Lev"},
		"Num":    {"04", "Numbers", "Num"},
		"Deut":   {"05", "Deuteronomy", "Deu"},
		"Josh":   {"06", "Joshua", "Jos"},
		"Judg":   {"07", "Judges", "Jud"},
		"Ruth":   {"08", "Ruth", "Rut"},
		"1Sam":   {"09", "1Samuel", "FSa"},
		"2Sam":   {"10", "2Samuel", "SSa"},
		"1Kings": {"11", "1Kings", "FKi"},
		"2Kings": {"12", "2Kings", "SKi"},
		"1Chron": {"13", "1Chronicles", "FCh"},
		"2Chron": {"14", "2Chronicles", "SCh"},
		"Ezra":   {"15", "Ezra", "Ezr"},
		"Neh":    {"16", "Nehemiah", "Neh"},
		"Esth":   {"17", "Esther", "Est"},
		"Job":    {"18", "Job", "Job"},
		"Psa":    {"19", "Psalms", "Psa"},
		"Prov":   {"20", "Proverbs", "Pro"},
		"Eccl":   {"21", "Ecclesiastes", "Ecc"},
		"SS":     {"22", "SongofSongs", "Son"},
		"Isa":    {"23", "Isaiah", "Isa"},
		"Jer":    {"24", "Jeremiah", "Jer"},
		"Lam":    {"25", "Lamentations", "Lam"},
		"Ezek":   {"26", "Ezekiel", "Eze"},
		"Dan":    {"27", "Daniel", "Dan"},
		"Hosea":  {"28", "Hosea", "Hos"},
		"Joel":   {"29", "Joel", "Joe"},
		"Amos":   {"30", "Amos", "Amo"},
		"Obad":   {"31", "Obadiah", "Oba"},
		"Jonah":  {"32", "Jonah", "Jon"},
		"Micah":  {"33", "Micah", "Mic"},
		"Nahum":  {"34", "Nahum", "Nah"},
		"Hab":    {"35", "Habakkuk", "Hab"},
		"Zeph":   {"36", "Zephaniah", "Zep"},
		"Hag":    {"37", "Haggai", "Hag"},
		"Zech":   {"38", "Zechariah", "Zec"},
		"Mal":    {"39", "Malachi", "Mal"},
		"Matt":   {"40", "Matthew", "Mat"},
		"Mark":   {"41", "Mark", "Mar"},
		"Luke":   {"42", "Luke", "Luk"},
		"John":   {"43", "John", "Joh"},
		"Acts":   {"44", "Acts", "Act"},
		"Rom":    {"45", "Romans", "Rom"},
		"1Cor":   {"46", "1Corinthians", "FCo"},
		"2Cor":   {"47", "2Corinthians", "SCo"},
		"Gal":    {"48", "Galatians", "Gal"},
		"Eph":    {"49", "Ephesians", "Eph"},
		"Phil":   {"50", "Philippians", "Phi"},
		"Col":    {"51", "Colossians", "Col"},
		"1Thes":  {"52", "1Thessalonians", "FTh"},
		"2Thes":  {"53", "2Thessalonians", "STh"},
		"1Tim":   {"54", "1Timothy", "FTi"},
		"2Tim":   {"55", "2Timothy", "STi"},
		"Titus":  {"56", "Titus", "Tit"},
		"Philem": {"57", "Philemon", "Phm"},
		"Heb":    {"58", "Hebrews", "Heb"},
		"James":  {"59", "James", "Jam"},
		"1Pet":   {"60", "1Peter", "FPe"},
		"2Pet":   {"61", "2Peter", "SPe"},
		"1John":  {"62", "1John", "FJo"},
		"2John":  {"63", "2John", "SJo"},
		"3John":  {"64", "3John", "TJo"},
		"Jude":   {"65", "Jude", "Jde"},
		"Rev":    {"66", "Revelation", "Rev"},
	}
}
