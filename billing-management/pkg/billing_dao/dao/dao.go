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
package dao

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/pkg/billing_dao/mysql"
)

type Dao interface {
	NewEntity() entity.Entity

	GetByID(int) (entity.Entity, error)
	Get(interface{}) (entity.Entity, error)
	GetByFilter(query interface{}, args ...interface{}) ([]entity.Entity, error)
	List() ([]entity.Entity, error)

	Create(entity.Entity) (entity.Entity, error)

	UpdateByID(int, entity.Entity) (entity.Entity, error)
	Update(interface{}) (entity.Entity, error)

	DeleteByID(int) error
	Delete(interface{}) (entity.Entity, error)

	Query(interface{}, *mysql.Query) (int, []entity.Entity, error)

	BeginTransaction() mysql.Session
}
