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
package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Code    int
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Message
}

func New(msg string, code int) error {
	return &Error{
		Message: msg,
		Code:    code,
	}
}

func Wrap(e error, code int) error {
	return &Error{
		Cause: e,
		Code:  code,
	}
}

func HandleError(e error, ctx *gin.Context) {
	err, ok := e.(*Error)
	if ok {
		ctx.String(err.Code, err.Error())
	} else {
		ctx.String(http.StatusInternalServerError, e.Error())
	}
	ctx.Abort()
}
