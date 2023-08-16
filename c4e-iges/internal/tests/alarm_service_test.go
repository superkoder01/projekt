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
package tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	require2 "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain/alarms"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/services/alarmservice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/mocks"
	"os"
	"testing"
)

type AlarmServiceTestSuite struct {
	suite.Suite
	Cfg *config.AppConfig
}

func (suite *AlarmServiceTestSuite) SetupSuite() {
	os.Setenv("CONFIGPATH", "../../config/")
	suite.Cfg, _ = config.LoadConfig()
}

func (suite *AlarmServiceTestSuite) TearDownSuite() {
}

func (suite *AlarmServiceTestSuite) SetupTest() {
}

func (suite *AlarmServiceTestSuite) TearDownTest() {
}

func (suite *AlarmServiceTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("START: %s.%s\n", suiteName, testName)
}

func (suite *AlarmServiceTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("FINISH: %s.%s\n", suiteName, testName)
}

func (suite *AlarmServiceTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println("=== Suite Stats:")
	fmt.Printf("=> suite duration: %v [s]\n", stats.End.Sub(stats.Start).Seconds())
	for _, v := range stats.TestStats {
		fmt.Printf("=> %s : %v [s] : %v\n", v.TestName, v.End.Sub(v.Start).Seconds(), passed(v.Passed))
	}
}

//func passed(status bool) string {
//	if status {
//		return fmt.Sprint("PASSED")
//	}
//	return fmt.Sprint("FAILED")
//}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAlarmServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AlarmServiceTestSuite))
}

///////////////////////////////////////////////////////////////////////
func (suite *AlarmServiceTestSuite) TestSendAlarmOK() {
	//require := require2.New(suite.T())

	ctx := context.Background()

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Error", mock.Anything, mock.Anything, mock.Anything).Once()

	alarmServiceFactory := alarmservice.NewAlarmServiceFactory(ctx, alarmservice.AlarmServiceType(alarmservice.LOGGER), loggerMock, suite.Cfg)
	alarmService := alarmServiceFactory.MakeService()
	alarmService.SendAlarm(ctx, alarms.OnGeneralError(fmt.Errorf("%v", "error")))
}

func (suite *AlarmServiceTestSuite) TestSendAlarmPanic() {
	require := require2.New(suite.T())

	ctx := context.Background()

	loggerMock := mocks.NewLogger(suite.T())
	loggerMock.On("Error", mock.Anything, mock.Anything, mock.Anything).Maybe()

	alarmServiceFactory := alarmservice.NewAlarmServiceFactory(ctx, alarmservice.AlarmServiceType("fake"), loggerMock, suite.Cfg)
	require.Panics(func() { alarmServiceFactory.MakeService() }, "unknown service: fake")
}
