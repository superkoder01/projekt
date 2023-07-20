package mysql

import (
	"database/sql"
)

type Session interface {
	Where(query interface{}, args ...interface{}) Session
	Or(query interface{}, args ...interface{}) Session
	Not(query interface{}, args ...interface{}) Session
	Limit(value int) Session
	Offset(value int) Session
	Order(value string) Session
	Select(query interface{}, args ...interface{}) Session
	Omit(columns ...string) Session
	Group(query string) Session
	Having(query string, values ...interface{}) Session
	Joins(query string, args ...interface{}) Session
	Take(out interface{}, where ...interface{}) Session
	First(out interface{}, where ...interface{}) Session
	Last(out interface{}, where ...interface{}) Session
	Find(out interface{}, where ...interface{}) Session
	Scan(dest interface{}) Session
	Count(*int64) Session
	Update(column string, value interface{}) Session
	Updates(values interface{}) Session
	Save(value interface{}) Session
	Create(value interface{}) Session
	Delete(value interface{}, where ...interface{}) Session
	Raw(sql string, values ...interface{}) Session
	Exec(sql string, values ...interface{}) Session
	Model(value interface{}) Session
	Begin(opts ...*sql.TxOptions) Session
	Commit() Session
	Rollback() Session
	AutoMigrate(values ...interface{}) error
	Error() error
	RowsAffected() int64
	Preload(query string, args ...interface{}) Session
	Table(string) Session
	Debug() Session
}
