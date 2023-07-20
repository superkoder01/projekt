package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type invoiceDao struct {
	bd.Session
}

func NewInvoiceDao(s bd.Session) *invoiceDao {
	return &invoiceDao{s}
}

func (inv *invoiceDao) NewEntity() entity.Entity {
	return entity.NewInvoice()
}

func (inv *invoiceDao) GetByID(i int) (entity.Entity, error) {
	var dbInvoice entity.Invoice
	err := inv.Where(entity.Invoice{ID: i}).Take(&dbInvoice).Error()
	return &dbInvoice, err
}

func (inv *invoiceDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbInvoices []entity.Invoice
	err := inv.Where(query, args).Find(&dbInvoices).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbInvoices))
	for i, v := range dbInvoices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (inv *invoiceDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := inv.Table(inv.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbInvoices []entity.Invoice
	err = inv.Limit(q.Limit).Offset(q.Offset).Order(q.Order).Where(v).Where(q.Filter).Find(&dbInvoices).Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbInvoices))
	for i, v := range dbInvoices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (inv *invoiceDao) List() ([]entity.Entity, error) {
	var dbInvoices []entity.Invoice
	err := inv.Find(&dbInvoices).Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbInvoices))
	for i, v := range dbInvoices {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (inv *invoiceDao) Create(e entity.Entity) (entity.Entity, error) {
	err := inv.Session.Create(e).Error()
	return e, err
}

func (inv *invoiceDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := inv.Where(&entity.Invoice{ID: i}).Updates(e).Error()
	return e, err
}

func (inv *invoiceDao) DeleteByID(i int) error {
	return inv.Where(&entity.Invoice{ID: i}).Delete(&entity.Invoice{}).Error()
}

func (inv *invoiceDao) BeginTransaction() bd.Session {
	return inv.Begin()
}

func (inv *invoiceDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (inv *invoiceDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (inv *invoiceDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
