package mysql

import (
	"context"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type session struct {
	DB  *gorm.DB
	Ctx context.Context
}

// NewSession - creates new database connection.
// connectionURL - example "user:pass@tcp(host:port)/database?tls=false&autocommit=true"
func NewSession(conURL string) (*session, error) {
	db, err := gorm.Open(mysql.Open(conURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return &session{DB: db}, err
}

// NewSessionWithContext - creates new database connection with context included.
// connectionURL - example "user:pass@tcp(host:port)/database?tls=false&autocommit=true"
func NewSessionWithContext(conURL string, ctx context.Context) (*session, error) {
	s, err := NewSession(conURL)
	s.Ctx = ctx
	return s, err
}

func (s *session) wrap(db *gorm.DB) Session {
	return &session{DB: db,
		Ctx: s.Ctx,
	}
}

func (s *session) Where(query interface{}, args ...interface{}) Session {
	return s.wrap(s.DB.Where(query, args...))
}

func (s *session) Or(query interface{}, args ...interface{}) Session {
	return s.wrap(s.DB.Or(query, args...))
}

func (s *session) Not(query interface{}, args ...interface{}) Session {
	return s.wrap(s.DB.Not(query, args...))
}

func (s *session) Limit(value int) Session {
	return s.wrap(s.DB.Limit(value))
}

func (s *session) Offset(value int) Session {
	return s.wrap(s.DB.Offset(value))
}

func (s *session) Order(value string) Session {
	return s.wrap(s.DB.Order(value))
}

func (s *session) Select(query interface{}, args ...interface{}) Session {
	return s.wrap(s.DB.Select(query, args...))
}

func (s *session) Omit(columns ...string) Session {
	return s.wrap(s.DB.Omit(columns...))
}

func (s *session) Group(query string) Session {
	return s.wrap(s.DB.Group(query))
}

func (s *session) Having(query string, values ...interface{}) Session {
	return s.wrap(s.DB.Having(query, values...))
}

func (s *session) Joins(query string, args ...interface{}) Session {
	return s.wrap(s.DB.Joins(query, args...))
}

func (s *session) Take(out interface{}, where ...interface{}) Session {
	return s.wrap(s.DB.Take(out, where...))
}

func (s *session) First(out interface{}, where ...interface{}) Session {
	return s.wrap(s.DB.First(out, where...))
}

func (s *session) Last(out interface{}, where ...interface{}) Session {
	return s.wrap(s.DB.Last(out, where...))
}

func (s *session) Find(out interface{}, where ...interface{}) Session {
	return s.wrap(s.DB.Find(out, where...))
}

func (s *session) Scan(dest interface{}) Session {
	return s.wrap(s.DB.Scan(dest))
}

func (s *session) Count(value *int64) Session {
	return s.wrap(s.DB.Count(value))
}

func (s *session) Update(column string, value interface{}) Session {
	return s.wrap(s.DB.Update(column, value))
}

func (s *session) Updates(values interface{}) Session {
	return s.wrap(s.DB.Updates(values))
}

func (s *session) Save(value interface{}) Session {
	return s.wrap(s.DB.Save(value))
}

func (s *session) Create(value interface{}) Session {
	return s.wrap(s.DB.Create(value))
}

func (s *session) Delete(value interface{}, where ...interface{}) Session {
	return s.wrap(s.DB.Delete(value, where...))
}

func (s *session) Raw(sql string, values ...interface{}) Session {
	return s.wrap(s.DB.Raw(sql, values...))
}

func (s *session) Exec(sql string, values ...interface{}) Session {
	return s.wrap(s.DB.Exec(sql, values...))
}

func (s *session) Model(value interface{}) Session {
	return s.wrap(s.DB.Model(value))
}

func (s *session) Begin(opts ...*sql.TxOptions) Session {
	return s.wrap(s.DB.Begin(opts...))
}

func (s *session) Commit() Session {
	return s.wrap(s.DB.Commit())
}

func (s *session) Rollback() Session {
	return s.wrap(s.DB.Rollback())
}

func (s *session) AutoMigrate(values ...interface{}) error {
	return s.DB.AutoMigrate(values...)
}

func (s *session) Error() error {
	return s.DB.Error
}

func (s *session) RowsAffected() int64 {
	return s.DB.RowsAffected
}

func (s *session) Preload(query string, args ...interface{}) Session {
	return s.wrap(s.DB.Preload(query, args))
}

func (s *session) Table(name string) Session {
	return s.wrap(s.DB.Table(name))
}

func (s *session) Debug() Session {
	return s.wrap(s.DB.Debug())
}
