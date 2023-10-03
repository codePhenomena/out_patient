package controller

import (
	"out_patient_appointment_system/models"
	"out_patient_appointment_system/utility"
	"time"
)

type Env struct {
	getdoctor interface {
		
		AddDoctors() ( error)
		GetDoctorDetails()([]models.DoctorsDetails,error)
		SetAppointments(models.Appointment) (error)
		GetDoctorsById(id int64, date time.Time) (bool, error)
	}

	
}

var Db *Env

func init() {
	Db = &Env{
		
		getdoctor :models.DoctorsModel{DB: utility.Db},
	}
}
