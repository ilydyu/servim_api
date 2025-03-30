// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Appointment struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Branch struct {
	ID          string        `json:"id"`
	Company     *Company      `json:"company"`
	Title       string        `json:"title"`
	Address     string        `json:"address"`
	Specialists []*Specialist `json:"specialists"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type Company struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Phone     string    `json:"phone"`
	Branches  []*Branch `json:"branches"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteAppointment struct {
	AppointmentID string `json:"appointmentId"`
}

type DeleteBranch struct {
	BranchID string `json:"branchId"`
}

type DeleteCompany struct {
	CompanyID string `json:"companyId"`
}

type DeleteService struct {
	ServiceID string `json:"serviceId"`
}

type DeleteSpecialist struct {
	SpecialistID string `json:"specialistId"`
}

type Mutation struct {
}

type NewAppointment struct {
	ServiceID   string    `json:"serviceId"`
	UserID      string    `json:"userId"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
}

type NewBranch struct {
	CompanyID string `json:"companyId"`
	Title     string `json:"title"`
	Address   string `json:"address"`
}

type NewCompany struct {
	Title string `json:"title"`
}

type NewService struct {
	SpecialistID string  `json:"specialistId"`
	Price        float64 `json:"price"`
}

type NewSpecialist struct {
	BranchID  string    `json:"branchId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Weekends  []*int32  `json:"weekends"`
	WorkFrom  time.Time `json:"work_from"`
	WorkUntil time.Time `json:"work_until"`
}

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

type Query struct {
}

type Service struct {
	ID           string         `json:"id"`
	Price        float64        `json:"price"`
	Appointments []*Appointment `json:"appointments"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

type Specialist struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Weekends  []*int32   `json:"weekends"`
	WorkFrom  time.Time  `json:"work_from"`
	WorkUntil time.Time  `json:"work_until"`
	Services  []*Service `json:"services,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type UpdateAppointment struct {
	AppointmentID string     `json:"appointmentId"`
	Description   *string    `json:"description,omitempty"`
	Datetime      *time.Time `json:"datetime,omitempty"`
}

type UpdateBranch struct {
	BranchID string  `json:"branchId"`
	Title    *string `json:"title,omitempty"`
	Address  *string `json:"address,omitempty"`
}

type UpdateCompany struct {
	CompanyID string `json:"companyId"`
	Title     string `json:"title"`
}

type UpdateService struct {
	ServiceID string   `json:"serviceId"`
	Price     *float64 `json:"price,omitempty"`
}

type UpdateSpecialist struct {
	SpecialistID string     `json:"specialistId"`
	FirstName    *string    `json:"firstName,omitempty"`
	LastName     *string    `json:"lastName,omitempty"`
	Weekends     []*int32   `json:"weekends,omitempty"`
	WorkFrom     *time.Time `json:"work_from,omitempty"`
	WorkUntil    *time.Time `json:"work_until,omitempty"`
}

type User struct {
	ID           string         `json:"id"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Appointments []*Appointment `json:"appointments"`
	Phone        string         `json:"phone"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}
