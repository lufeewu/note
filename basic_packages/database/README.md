# 简介
database/sql 标准库是 golang 提供的用于连接 sql 或类数据库的泛用接口. database/sql/driver 则定义被数据库驱动实现的接口，它们会被 sql 包使用. 

## 源码
总计 11200 多行代码，除去测试代码 4700 多行.
+ type DB struct
    - 数据库句柄，代表一个具有零到多个底层连接的连接池.
    - 线程安全，可以安全的被多个 goroutine 同时使用 
    - func Open(driverName, dataSourceName string) (*DB, error)
    - func (db *DB) Driver() driver.Driver
    - func (db *DB) Ping() error
    - func (db *DB) Close() error
    - func (db *DB) SetMaxOpenConns(n int)
    - func (db *DB) SetMaxIdleConns(n int)
    - func (db *DB) Exec(query string, args ...interface{}) (Result, error)
    - func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
    - func (db *DB) QueryRow(query string, args ...interface{}) *Row
    - func (db *DB) Prepare(query string) (*Stmt, error)
    - func (db *DB) Begin() (*Tx, error)
+ type Scanner interface
+ type NullBool struct
+ type NullInt64 struct
+ type Result interface
+ type Rows struct
+ database/sql/driver
    - type Value interface 它是驱动必须能处理的值. 支持一以下数据类型:
        - nil
        - int64
        - float64
        - bool
        - []byte
        - string   [*] Rows.Next不会返回该类型值
        - time.Time       
    - type Valuer interface 
    - type ValueConverter interface 
    - type ColumnConverter interface
    - type Driver interface
    - type Conn interface
    - type Execer interface
    - type Queryer interface
    - type Stmt interface 
    - type Tx interface
    - type Result interface
    - type Rows interface
+ ...

## 应用
1. github.com/influxdata/flux
2. github.com/jmoiron/sqlx
3. github.com/jinzhu/gorm

## sql driver
涵盖了绝大多数的数据库 driver，如 MySQL、DB2、ODBC、MS SQL Server、Postgres、Oracle、SQLite 等，[驱动列表](https://github.com/golang/go/wiki/SQLDrivers)如下:
+ Apache Ignite/GridGain: https://github.com/amsokol/ignite-go-client
+ Apache Impala: https://github.com/bippio/go-impala
+ Apache Avatica/Phoenix: https://github.com/apache/calcite-avatica-go
+ AWS Athena: https://github.com/segmentio/go-athena
+ ClickHouse (uses native TCP interface): https://github.com/kshvakov/clickhouse
+ ClickHouse (uses HTTP API): https://github.com/mailru/go-clickhouse
+ CockroachDB: Use any PostgreSQL driver
+ Couchbase N1QL: https://github.com/couchbase/go_n1ql
+ DB2 LUW and DB2/Z with DB2-Connect: https://bitbucket.org/phiggins/db2cli (Last updated 2015-08)
+ DB2 LUW (uses cgo): https://github.com/asifjalil/cli
+ DB2 LUW, z/OS, iSeries and Informix: https://github.com/ibmdb/go_ibm_db
+ Firebird SQL: https://github.com/nakagami/firebirdsql
+ MS ADODB: https://github.com/mattn/go-adodb
+ MS SQL Server (pure go): https://github.com/denisenkom/go-mssqldb
+ MS SQL Server (uses cgo): https://github.com/minus5/gofreetds
+ MySQL: https://github.com/ziutek/mymysql [*]
+ MySQL: https://github.com/go-sql-driver/mysql/ [*]
+ ODBC: https://bitbucket.org/miquella/mgodbc (Last updated 2016-02)
+ ODBC: https://github.com/alexbrainman/odbc
+ Oracle: https://github.com/mattn/go-oci8
+ Oracle: https://gopkg.in/rana/ora.v4
+ Oracle: https://gopkg.in/goracle.v2
+ QL: http://godoc.org/github.com/cznic/ql/driver
+ Postgres (pure Go): https://github.com/lib/pq [*]
+ Postgres (uses cgo): https://github.com/jbarham/gopgsqldriver
+ Postgres (pure Go): https://github.com/jackc/pgx [**]
+ Presto: https://github.com/prestodb/presto-go-client
+ SAP HANA (uses cgo): https://help.sap.com/viewer/0eec0d68141541d1b07893a39944924e/2.0.03/en-US/0ffbe86c9d9f44338441829c6bee15e6.html
+ SAP HANA (pure go): https://github.com/SAP/go-hdb
+ SAP ASE (uses cgo): https://github.com/SAP/go-ase - package cgo (pure go package planned)
+ Snowflake (pure Go): https://github.com/snowflakedb/gosnowflake
+ SQLite (uses cgo): https://github.com/mattn/go-sqlite3 [*]
+ SQLite (uses cgo): https://github.com/gwenn/gosqlite - Supports SQLite dynamic data typing
+ SQLite (uses cgo): https://github.com/mxk/go-sqlite
+ SQLite: (uses cgo): https://github.com/rsc/sqlite
+ SQL over REST: https://github.com/adaptant-labs/go-sql-rest-driver
+ Sybase SQL Anywhere: https://github.com/a-palchikov/sqlago
+ Sybase ASE (pure go): https://github.com/thda/tds
+ Vertica: https://github.com/vertica/vertica-sql-go
+ Vitess: https://godoc.org/vitess.io/vitess/go/vt/vitessdriver
+ YQL (Yahoo! Query Language): https://github.com/mattn/go-yql
+ Apache Hive: https://github.com/sql-machine-learning/gohive
+ MaxCompute: https://github.com/sql-machine-learning/gomaxcompute

## MySQL driver 例子
+ github.com/go-sql-driver/mysql
    + 使用 database/sql 连接 mysql 数据库时，需要引入该包
    + 它实现了 database/sql 的相关接口
    + driver.go 实现 Driver 接口

          func (d MySQLDriver) Open(dsn string) (driver.Conn, error) {
                cfg, err := ParseDSN(dsn)
                if err != nil {
                    return nil, err
                }
                c := &connector{
                    cfg: cfg,
                }
                return c.Connect(context.Background())
          }
    
    + type connector struct
        - func (c *connector) Connect(ctx context.Context) (driver.Conn, error)
    + type mysqlConn struct
        - netConn net.Conn
        - ...

## ref
1. [Golang SQL database drivers](https://github.com/golang/go/wiki/SQLDrivers)
2. [关于Golang中database/sql包的学习笔记](https://segmentfault.com/a/1190000003036452)