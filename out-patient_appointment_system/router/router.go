package router

import (
	"net/http"
	"out_patient_appointment_system/controller"
	
	"strings"
)
func Routes(w http.ResponseWriter, r *http.Request) {

	route := strings.Trim(r.URL.Path, "/")
	switch route {
	case "addDoctor":
		 if r.Method == "POST" {
			controller.AddDoctors(w, r)
			} 
		case "setAppointment":
			if r.Method=="POST"{
				controller.SetAppointments(w,r)
			}
	case "getDoctors":
		if r.Method=="GET"{
			controller.GetDoctorDetail(w,r)
		}
	}
}
