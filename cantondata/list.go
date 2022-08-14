package cantondata

// getList returns a list of canton codes with names, according to:
// https://www.bfs.admin.ch/bfs/en/home/basics/symbols-abbreviations.html
func GetList() map[string]string {
	return map[string]string{
		"ZH": "Zurich",
		"BE": "Bern",
		"LU": "Lucerne",
		"UR": "Uri",
		"SZ": "Schwyz",
		"OW": "Obwalden",
		"NW": "Nidwalden",
		"GL": "Glarus",
		"ZG": "Zug",
		"FR": "Fribourg",
		"SO": "Solothurn",
		"BS": "Basel-Stadt",
		"BL": "Basel-Landschaft",
		"SH": "Schaffhausen",
		"AR": "Appenzell A. Rh.",
		"AI": "Appenzell I. Rh.",
		"SG": "St. Gallen",
		"GR": "Graubünden",
		"AG": "Aargau",
		"TG": "Thurgau",
		"TI": "Ticino",
		"VD": "Vaud",
		"VS": "Valais",
		"NE": "Neuchâtel",
		"GE": "Geneva",
		"JU": "Jura",
	}
}
