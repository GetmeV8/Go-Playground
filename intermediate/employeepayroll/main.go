package main

import (
	"fmt"
)

type Employee interface {
	CalculateSalary() float64
}

type FullTime struct {
	Name          string
	MonthlySalary float64
}

type PartTime struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (ft *FullTime) CalculateSalary() float64 {
	return ft.MonthlySalary * 12
}

func (pt *PartTime) CalculateSalary() float64 {
	if pt.HoursWorked > 40 {
		overtime := pt.HoursWorked - 40
		regularPay := 40 * pt.HourlyRate
		overtimePay := overtime * pt.HourlyRate * 1.5
		return regularPay + overtimePay
	}
	return pt.HourlyRate * pt.HoursWorked
}

func NewFullTime(name string, monthlysalary float64) *FullTime {
	return &FullTime{
		Name:          name,
		MonthlySalary: monthlysalary,
	}
}

func NewPartTime(name string, hourlyrate float64, hoursworked float64) *PartTime {
	return &PartTime{
		Name:        name,
		HourlyRate:  hourlyrate,
		HoursWorked: hoursworked,
	}
}

func main() {
	emp1 := NewFullTime("Raamz", 35000)
	emp2 := NewPartTime("Sarah", 25.50, 45)

	employees := []Employee{emp1, emp2}
	for _, employee := range employees {
		fmt.Printf("Annual salary: $%.2f\n", employee.CalculateSalary())
	}

	if ft, ok := employees[0].(*FullTime); ok {
		fmt.Printf("%s (Full-time) earns $%.2f per year\n", ft.Name, ft.CalculateSalary())
	}
	if pt, ok := employees[1].(*PartTime); ok {
		fmt.Printf("%s (Part-time) earns $%.2f per year\n", pt.Name, pt.CalculateSalary())
	}
}
