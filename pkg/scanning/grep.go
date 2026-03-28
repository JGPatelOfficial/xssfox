package scanning

import "regexp"

// Grepping is function for checking pattern
func Grepping(data, regex string) []string {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return []string{}
	}
	return pattern.FindAllString(data, -1)
}

// builtinGrep is xssfox build-in grep pattern
func builtinGrep(data string) map[string][]string {
	return grepPatterns(data, builtinPatterns)
}

// customGrep is user custom grep pattern
func customGrep(data string, pattern map[string]string) map[string][]string {
	return grepPatterns(data, pattern)
}

// grepPatterns is a helper function to grep patterns from data
func grepPatterns(data string, patterns map[string]string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range patterns {
		resultArr := Grepping(data, v)
		if len(resultArr) > 0 {
			result[k] = resultArr
		}
	}
	return result
}

// builtinPatterns is a map of xssfox built-in grep patterns
var builtinPatterns = map[string]string{
	"xssfox-ssti":                  `2958816`,
	"xssfox-esii":                  `<esii-xssfox>`,
	"xssfox-rsa-key":               `-----BEGIN RSA PRIVATE KEY-----|-----END RSA PRIVATE KEY-----`,
	"xssfox-priv-key":              `-----BEGIN PRIVATE KEY-----|-----END PRIVATE KEY-----`,
	"xssfox-aws-s3":                `s3\.amazonaws\.com[/]+|[a-zA-Z0-9_-]*\.s3\.amazonaws\.com`,
	"xssfox-aws-appsync-graphql":   `da2-[a-z0-9]{26}`,
	"xssfox-slack-webhook1":        `^https://hooks\.slack\.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	"xssfox-slack-webhook2":        `^https://hooks\.slack\.com/services/T[a-zA-Z0-9_]{8,10}/B[a-zA-Z0-9_]{8,10}/[a-zA-Z0-9_]{24}`,
	"xssfox-slack-token":           `(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})`,
	"xssfox-facebook-oauth":        `[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].{0,30}['"\s][0-9a-f]{32}['"\s]`,
	"xssfox-twitter-oauth":         `[t|T][w|W][i|I][t|T][t|T][e|E][r|R].{0,30}['"\s][0-9a-zA-Z]{35,44}['"\s]`,
	"xssfox-heroku-api":            `[h|H][e|E][r|R][o|O][k|K][u|U].{0,30}[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`,
	"xssfox-mailgun-api":           `key-[0-9a-zA-Z]{32}`,
	"xssfox-mailchamp-api":         `[0-9a-f]{32}-us[0-9]{1,2}`,
	"xssfox-picatic-api":           `sk_live_[0-9a-z]{32}`,
	"xssfox-google-oauth-id":       `[0-9]+-[0-9A-Za-z_]{32}\.apps\.googleusercontent\.com`,
	"xssfox-google-api":            `AIza[0-9A-Za-z-_]{35}`,
	"xssfox-google-oauth":          `ya29\.[0-9A-Za-z\-_]+`,
	"xssfox-aws-access-key":        `AKIA[0-9A-Z]{16}`,
	"xssfox-amazon-mws-auth-token": `amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`,
	"xssfox-facebook-access-token": `EAACEdEose0cBA[0-9A-Za-z]+`,
	"xssfox-github-access-token":   `\b[a-zA-Z0-9_-]*:[a-zA-Z0-9_\-]+@github\.com\b`,
	"xssfox-github":                `[gG][iI][tT][hH][uU][bB].*['|"][0-9a-zA-Z]{35,40}['|"]`,
	"xssfox-gitlab-token":          `glpat-[0-9a-zA-Z-_]{20}`,
	"xssfox-azure-storage":         `[a-zA-Z0-9_-]*\.file\.core\.windows\.net`,
	"xssfox-telegram-bot-api-key":  `[0-9]+:AA[0-9A-Za-z\-_]{33}`,
	"xssfox-square-access-token":   `sq0atp-[0-9A-Za-z\-_]{22}`,
	"xssfox-square-oauth-secret":   `sq0csp-[0-9A-Za-z\-_]{43}`,
	"xssfox-twitter-access-token":  `[tT][wW][iI][tT][tT][eE][rR].*[1-9][0-9]+-[0-9a-zA-Z]{40}`,
	"xssfox-twilio-api-key":        `SK[0-9a-fA-F]{32}`,
	"xssfox-braintree-token":       `access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}`,
	"xssfox-stripe-api-key":        `sk_live_[0-9a-zA-Z]{24}`,
	"xssfox-stripe-restricted-key": `rk_live_[0-9a-zA-Z]{24}`,
	"xssfox-shopify-access-token":  `shpat_[0-9a-fA-F]{32}`,
	"xssfox-linear-api-key":        `lin_api_[a-zA-Z0-9]{40}`,
	"xssfox-digitalocean-token":    `dop_v1_[a-z0-9]{64}`,
	"xssfox-asana-access-Token":    `0/[0-9a-z]{32}`,
	"xssfox-dropbox-access-Token":  `sl\.[A-Za-z0-9_-]{20,100}`,
	"xssfox-sendGrid-api-key":      `SG\.[\w\d\-_]{22}\.[\w\d\-_]{43}`,
	"xssfox-firebase-secret":       `AAAA[A-Za-z0-9_-]{7}:[A-Za-z0-9_-]{140}`,
	"xssfox-netlify-token":         `netlifyAuthToken\s*=\s*['"][A-Za-z0-9_-]{40,64}['"]`,
	"xssfox-sentry-dsn":            `https://[a-zA-Z0-9]+@[a-z]+\.ingest\.sentry\.io/\d+`,
	"xssfox-error-mysql":           `(SQL syntax.*MySQL|Warning.*mysql_.*|MySqlException \(0x|valid MySQL result|check the manual that corresponds to your (MySQL|MariaDB) server version|MySqlClient\.|com\.mysql\.jdbc\.exceptions)`,
	"xssfox-error-postgresql":      `(PostgreSQL.*ERROR|Warning.*\Wpg_.*|valid PostgreSQL result|Npgsql\.|PG::SyntaxError:|org\.postgresql\.util\.PSQLException|ERROR:\s\syntax error at or near)`,
	"xssfox-error-mssql":           `(Driver.* SQL[\-\_\ ]*Server|OLE DB.* SQL Server|\bSQL Server.*Driver|Warning.*mssql_.*|\bSQL Server.*[0-9a-fA-F]{8}|[\s\S]Exception.*\WSystem\.Data\.SqlClient\.|[\s\S]Exception.*\WRoadhouse\.Cms\.|Microsoft SQL Native Client.*[0-9a-fA-F]{8})`,
	"xssfox-error-msaccess":        `(Microsoft Access (\d+ )?Driver|JET Database Engine|Access Database Engine|ODBC Microsoft Access)`,
	"xssfox-error-oracle":          `(\bORA-\d{5}|Oracle error|Oracle.*Driver|Warning.*\Woci_.*|Warning.*\Wora_.*)`,
	"xssfox-error-ibmdb2":          `(CLI Driver.*DB2|DB2 SQL error|\bdb2_\w+\(|SQLSTATE.+SQLCODE)`,
	"xssfox-error-informix":        `(Exception.*Informix)`,
	"xssfox-error-firebird":        `(Dynamic SQL Error|Warning.*ibase_.*)`,
	"xssfox-error-sqlite":          `(SQLite/JDBCDriver|SQLite\.Exception|System\.Data\.SQLite\.SQLiteException|Warning.*sqlite_.*|Warning.*SQLite3::|\[SQLITE_ERROR\])`,
	"xssfox-error-sapdb":           `(SQL error.*POS([0-9]+).*|Warning.*maxdb.*)`,
	"xssfox-error-sybase":          `(Warning.*sybase.*|Sybase message|Sybase.*Server message.*|SybSQLException|com\.sybase\.jdbc)`,
	"xssfox-error-ingress":         `(Warning.*ingres_|Ingres SQLSTATE|Ingres\W.*Driver)`,
	"xssfox-error-frontbase":       `(Exception (condition )?\d+. Transaction rollback.)`,
	"xssfox-error-hsqldb":          `(org\.hsqldb\.jdbc|Unexpected end of command in statement \[|Unexpected token.*in statement \[)`,

	//sqli
	/////////////////////////////////////////////////////////

	//mysql
	"xssfox-error-mysql1":  `SQL syntax.*?MySQL`,
	"xssfox-error-mysql2":  `Warning.*?mysqli?`,
	"xssfox-error-mysql3":  `MySQLSyntaxErrorException`,
	"xssfox-error-mysql4":  `valid MySQL result`,
	"xssfox-error-mysql5":  `check the manual that (corresponds to|fits) your MySQL server version`,
	"xssfox-error-mysql6":  `check the manual that (corresponds to|fits) your MariaDB server version`,
	"xssfox-error-mysql7":  `check the manual that (corresponds to|fits) your Drizzle server version`,
	"xssfox-error-mysql8":  `Unknown column '[^ ]+' in 'field list'`,
	"xssfox-error-mysql9":  `com\.mysql\.jdbc`,
	"xssfox-error-mysql10": `Zend_Db_(Adapter|Statement)_Mysqli_Exception`,
	"xssfox-error-mysql11": `MySqlException`,
	"xssfox-error-mysql12": `Syntax error or access violation`,

	//psql
	"xssfox-error-psql1":  `PostgreSQL.*?ERROR`,
	"xssfox-error-psql2":  `Warning.*?\Wpg_`,
	"xssfox-error-psql3":  `valid PostgreSQL result`,
	"xssfox-error-psql4":  `Npgsql\.`,
	"xssfox-error-psql5":  `PG::SyntaxError:`,
	"xssfox-error-psql6":  `org\.postgresql\.util\.PSQLException`,
	"xssfox-error-psql7":  `ERROR:\s\syntax error at or near`,
	"xssfox-error-psql8":  `ERROR: parser: parse error at or near`,
	"xssfox-error-psql9":  `PostgreSQL query failed`,
	"xssfox-error-psql10": `org\.postgresql\.jdbc`,
	"xssfox-error-psql11": `PSQLException`,

	//mssql
	"xssfox-error-mssql1":  `Driver.*? SQL[\-\_\ ]*Server`,
	"xssfox-error-mssql2":  `OLE DB.*? SQL Server`,
	"xssfox-error-mssql3":  `\bSQL Server[^&lt;&quot;]+Driver`,
	"xssfox-error-mssql4":  `Warning.*?\W(mssql|sqlsrv)_`,
	"xssfox-error-mssql5":  `\bSQL Server[^&lt;&quot;]+[0-9a-fA-F]{8}`,
	"xssfox-error-mssql6":  `System\.Data\.SqlClient\.SqlException`,
	"xssfox-error-mssql7":  `(?s)Exception.*?\bRoadhouse\.Cms\.`,
	"xssfox-error-mssql8":  `Microsoft SQL Native Client error '[0-9a-fA-F]{8}`,
	"xssfox-error-mssql9":  `\[SQL Server\]`,
	"xssfox-error-mssql10": `ODBC SQL Server Driver`,
	"xssfox-error-mssql11": `ODBC Driver \d+ for SQL Server`,
	"xssfox-error-mssql12": `SQLServer JDBC Driver`,
	"xssfox-error-mssql13": `com\.jnetdirect\.jsql`,
	"xssfox-error-mssql14": `macromedia\.jdbc\.sqlserver`,
	"xssfox-error-mssql15": `Zend_Db_(Adapter|Statement)_Sqlsrv_Exception`,
	"xssfox-error-mssql16": `com\.microsoft\.sqlserver\.jdbc`,
	"xssfox-error-mssql18": `SQL(Srv|Server)Exception`,
}
