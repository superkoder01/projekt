package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type workerDao struct {
	bd.Session
}

func NewWorkerDao(s bd.Session) *workerDao {
	return &workerDao{s}
}

func (w *workerDao) NewEntity() entity.Entity {
	return entity.NewWorker()
}

func (w *workerDao) GetByID(i int) (entity.Entity, error) {
	var dbWorker entity.WorkerJoinUserRole
	err := w.Model(&entity.Worker{}).
		Select("WORKER.*, USER.ROLE_ID").
		Joins("left join USER on WORKER.ID = USER.WORKER_ID").
		Where(&entity.Worker{ID: i}).
		Scan(&dbWorker).
		Error()

	return &dbWorker, err
}

func (w *workerDao) GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error) {
	var dbWorkers []entity.WorkerJoinUserRole
	err := w.Model(&entity.Worker{}).
		Select("WORKER.*, USER.ROLE_ID").
		Joins("left join USER on WORKER.ID = USER.WORKER_ID").
		Where(query, args).
		Scan(&dbWorkers).
		Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbWorkers))
	for i, v := range dbWorkers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (w *workerDao) Query(v interface{}, q *bd.Query) (int, []entity.Entity, error) {
	var count int64
	err := w.Table(w.NewEntity().TableName()).Where(v).Where(q.Filter).Count(&count).Error()
	if err != nil {
		return 0, nil, err
	}

	var dbWorkers []entity.WorkerJoinUserRole
	err = w.Model(&entity.WorkerJoinUserRole{}).
		Select("WORKER.*, USER.ROLE_ID").
		Joins("left join USER on WORKER.ID = USER.WORKER_ID").
		Where(v).
		Where(q.Filter).
		Limit(q.Limit).
		Offset(q.Offset).
		Order(q.Order).
		Scan(&dbWorkers).
		Error()
	if err != nil {
		return 0, nil, err
	}

	entities := make([]entity.Entity, len(dbWorkers))
	for i, v := range dbWorkers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return int(count), entities, nil
}

func (w *workerDao) List() ([]entity.Entity, error) {
	var dbWorkers []entity.WorkerJoinUserRole
	err := w.Model(&entity.Worker{}).
		Select("WORKER.*, USER.ROLE_ID").
		Joins("left join USER on WORKER.ID = USER.WORKER_ID").
		Scan(&dbWorkers).
		Error()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Entity, len(dbWorkers))
	for i, v := range dbWorkers {
		a := v
		entities[i] = entity.Entity(&a)
	}

	return entities, nil
}

func (w *workerDao) Create(e entity.Entity) (entity.Entity, error) {
	err := w.Session.Create(e).Error()
	return e, err
}

func (w *workerDao) UpdateByID(i int, e entity.Entity) (entity.Entity, error) {
	err := w.Where(&entity.Worker{ID: i}).Updates(e).Error()
	return e, err
}

func (w *workerDao) DeleteByID(i int) error {
	return w.Where(&entity.Worker{ID: i}).Delete(&entity.Worker{}).Error()
}

func (w *workerDao) BeginTransaction() bd.Session {
	return w.Begin()
}

func (w *workerDao) Get(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (w *workerDao) Update(i interface{}) (entity.Entity, error) {
	return nil, nil
}

func (w *workerDao) Delete(i interface{}) (entity.Entity, error) {
	return nil, nil
}
