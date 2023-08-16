/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package service

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
)

type Service interface {
	Create(model.Model) (model.Model, error)

	DeleteByID(int) error
	Delete(interface{}) (model.Model, error)

	GetByID(int) (model.Model, error)
	Get(interface{}) (model.Model, error)
	GetWithFilter(query interface{}, args ...interface{}) ([]model.Model, error)
	List() ([]model.Model, error)

	UpdateByID(int, model.Model) (model.Model, error)
	Update(interface{}) (model.Model, error)

	ModelToEntity(model.Model) (entity.Entity, error)
	EntityToModel(entity.Entity) (model.Model, error)
}
