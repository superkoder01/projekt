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
package alarms

import "fmt"

func OnApplicationInitializationError(err error) *Alarm {
	return &Alarm{
		Id:          AlarmId_IGES001,
		Severity:    SeverityCritical,
		Description: fmt.Sprintf("unable to initialize application, reason: %v", err),
	}
}

func OnDataAccessProblem(err error) *Alarm {
	return &Alarm{
		Id:          AlarmId_IGES002,
		Severity:    SeverityCritical,
		Description: fmt.Sprintf("unable to access repository, reason: %v", err),
	}
}

func OnConfigurationMismatch(err error) *Alarm {
	return &Alarm{
		Id:          AlarmId_IGES003,
		Severity:    SeverityError,
		Description: fmt.Sprintf("configuration mismatch, reason: %v", err),
	}
}

func OnPublishingInvoiceProblem(err error) *Alarm {
	return &Alarm{
		Id:          AlarmId_IGES004,
		Severity:    SeverityCritical,
		Description: fmt.Sprintf("unable to publish invoice, reason: %v", err),
	}
}

func OnGeneralError(err error) *Alarm {
	return &Alarm{
		Id:          AlarmId_IGES005,
		Severity:    SeverityCritical,
		Description: fmt.Sprintf("general error, reason: %v", err),
	}
}

type Alarm struct {
	Id          AlarmId
	Severity    Severity
	Description string
}

func (a Alarm) String() string {
	return fmt.Sprintf("ALARM[id: %v, severity: %v, description: %v]", a.Id, a.Severity, a.Description)
}

type AlarmId int64

const (
	AlarmId_IGES001 AlarmId = iota // app initialization error
	AlarmId_IGES002                // data access problem
	AlarmId_IGES003                // configuration mismatch
	AlarmId_IGES004                // publishing invoice problem
	AlarmId_IGES005                // general error
	AlarmId_IGES006
)

func (ai AlarmId) String() string {
	switch ai {
	case AlarmId_IGES001:
		return "IGES001"
	case AlarmId_IGES002:
		return "IGES002"
	case AlarmId_IGES003:
		return "IGES003"
	case AlarmId_IGES004:
		return "IGES004"
	case AlarmId_IGES005:
		return "IGES005"
	case AlarmId_IGES006:
		return "IGES006"
	default:
		return "UnknownAlarmId"
	}
}

type Severity int64

const (
	SeverityInfo Severity = iota
	SeverityWarn
	SeverityError
	SeverityCritical
)

func (s Severity) String() string {
	switch s {
	case SeverityInfo:
		return "INFO"
	case SeverityWarn:
		return "WARN"
	case SeverityError:
		return "ERROR"
	case SeverityCritical:
		return "CRITICAL"
	default:
		return "UnknownSeverity"
	}
}
