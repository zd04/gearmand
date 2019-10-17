module github.com/zd04/gearmand

go 1.12

require (
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab
	github.com/go-redis/redis/v7 v7.0.0-beta.4
	github.com/go-sql-driver/mysql v1.4.1
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/ngaut/gearmand v0.0.0-00010101000000-000000000000
	github.com/ngaut/logging v0.0.0-20150203141111-f98f5f4cd523
	github.com/ngaut/stats v0.0.0-20140515082619-8564f986a400
)

replace github.com/ngaut/gearmand => ./
