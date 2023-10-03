package controller
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"out_patient_appointment_system/models"
	"out_patient_appointment_system/utility"
	// "net/smtp"
	// "os"
)

func JsonDecoder(r *http.Request) (models.Appointment, error) {
	var usersData models.Appointment
	err := json.NewDecoder(r.Body).Decode(&usersData)
	if err != nil {
		log.Println(err)
	}
	return usersData, err
}
func AddDoctors(w http.ResponseWriter,r *http.Request)utility.AjaxResponse{
	response := utility.AjaxResponse{Status: "success", Message: "Doctors Details Added successfully", Payload: []interface{}{}}

	err:=Db.getdoctor.AddDoctors()
	if err!=nil{
		log.Println(err)
		response.Status = "failure"

	}else{
		response.Status = "success"
	}
	utility.RenderTemplate(w,r,"",response)
	return response
}

func GetDoctorDetail(w http.ResponseWriter,r *http.Request)utility.AjaxResponse{
	response := utility.AjaxResponse{Status: "success", Message: "View Doctors details successfully ", Payload: []interface{}{}}

	res,err:=Db.getdoctor.GetDoctorDetails()
	log.Println("......",res)
	if err!=nil{
		log.Println(err)
		response.Status = "failure"

	}else{
		response.Status = "success"
		response.Payload=res
	}
	utility.RenderTemplate(w,r,"",response)
	return response
}
func SetAppointments(w http.ResponseWriter,r *http.Request)utility.AjaxResponse{
	response := utility.AjaxResponse{Status: "success", Message: "Appointment Scheduled successfully", Payload: []interface{}{}}
	res,err:=JsonDecoder(r)
	if err!=nil{
		log.Println(err)
	}else{
	
	log.Println(".......",res.Date)
	date:=res.Date
	//assuming working hours 
	workingHoursStart := time.Date(date.Year(), date.Month(), date.Day(), 17, 0, 0, 0, time.UTC) // Assuming doctors work from 5 PM
	workingHoursEnd := time.Date(date.Year(), date.Month(), date.Day(), 23, 0, 0, 0, time.UTC)   // Assuming doctors work until 11 PM
    
	//checking working hours
    if date.Before(workingHoursStart) || date.After(workingHoursEnd) {
        response.Status = "failure"
		response.Message="Appointment is outside working hours" // Appointment is outside working hours
    }else{
		//if everything above is  good then get data from appointment to validate by schedule 
		isValid,err:=Db.getdoctor.GetDoctorsById(res.DoctorId,date) 
		if err!=nil{
			log.Println(err)
		}else{
			if isValid!=true{
				response.Status = "failure"
				response.Message= "slot already booked"
			}else{
				//if time slot is available book appointment
			    err:=Db.getdoctor.SetAppointments(res)
				if err!=nil{
					log.Println(err)
					response.Status = "failure"
				}else{
					response.Status = "success"
					response.Payload=res
				}
			}
		}
	}
}
	utility.RenderTemplate(w,r,"",response)
	return response
}


