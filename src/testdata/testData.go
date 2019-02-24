package testdata

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/HRrepository"
	"time"
)

var Employees = []HRrepository.Employee{
	{
		EmployeeID:        1,
		Name:              "petro",
		LastName:          "1stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
	{
		EmployeeID:        2,
		Name:              "ivan",
		LastName:          "2stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
	{
		EmployeeID:        3,
		Name:              "3st",
		LastName:          "3stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
	{
		EmployeeID:        4,
		Name:              "ivan",
		LastName:          "3stLn",
		MiddleName:        "",
		PassSeriesNum:     "",
		IdentificationNum: "",
		BirthDate:         time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		Position:          "",
		JoinDate:          time.Date(1989, time.April, 0, 0, 0, 0, 0, time.Local),
		QuitDate:          time.Now(),
		IsQuit:            false},
}
