# USQL

> `go install github.com/xo/usql@latest`

## Building
When building usql out-of-the-box with go build or go install, only the base drivers for PostgreSQL, MySQL, SQLite3, Microsoft SQL Server, Oracle, CSVQ will be included in the build:

### build/install with base drivers (PostgreSQL, MySQL, SQLite3, Microsoft SQL Server, Oracle, CSVQ)
> `go install github.com/xo/usql@master`

* `\drivers`
*  \conninfo

## connect to a mysql database
$ usql my://user:pass@host/dbname
$ usql mysql://user:pass@host:port/dbname
$ usql my://
$ usql /var/run/mysqld/mysqld.sock

> `usql my://mostain:Mateors321@localhost/master_host`

> \echo Welcome `echo $USER` -- 'currently:' "(" `date` ")"

> `usql --help`

> `usql -J my://mostain:Mateors321@localhost/master_host`

> `usql --json my://mostain:Mateors321@localhost/master_host`

> `\d login`

### Other databases can be enabled by specifying the build tag for their database driver.
>  go install -tags 'avatica odbc' github.com/xo/usql@master
