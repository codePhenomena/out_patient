package models

import (
	"log"
	"out_patient_appointment_system/utility"
	"github.com/jmoiron/sqlx"
	"time"
)
type DoctorsDetails struct {
	Id       int64  `db:"id"`
    Name     string `db:"name"`
    Location string `db:"location"`
	Details  string `db:"details"`
}
type Appointment struct {
    ID        string      	`db:"id"`
    DoctorId  int64   	 	 `db:"doctor_id"`
    PatientId int64    	  	`db:"patient_id"`
	Date      time.Time		`db:"date"`
	DateOf    string
	
}
type DoctorsModel struct {
	DB *sqlx.DB
}
func (data DoctorsModel) AddDoctors() ( error) {
	query := `INSERT INTO doctors (name ,location,details)
	VALUES (:name ,:location,:details)
	`
	values := map[string]interface{}{
	    "name":      "Dr.Ashish",
	    "location": "jabalpur"   ,
	    "details": "heart specialist" ,
	    
	}

	_, err := utility.Db.NamedExec(query, values)
	if err != nil {
	    log.Println(err)
	}
		return  err
}

func (getData DoctorsModel) GetDoctorDetails()([]DoctorsDetails,error){
	var resulData []DoctorsDetails
	var err error
	rows,err:=utility.Db.Queryx("SELECT * FROM doctors")
	if err!=nil{
		log.Println(err)
	}else{
		defer rows.Close()
		for rows.Next(){
			var dataRow DoctorsDetails
			err:= rows.StructScan(&dataRow)
			if err!=nil{
				log.Println(err)
			}else{
				resulData=append(resulData,dataRow)
			}
		}
	}
	
	return resulData, err
}

func (data DoctorsModel) SetAppointments(appointDetails Appointment) (error) {

_,err := utility.Db.NamedExec("INSERT INTO `appointments` (doctor_id,patient_id,date) VALUES (:DoctorId,:PatientId,:Date)",
		map[string]interface{}{
		"DoctorId":appointDetails.DoctorId,
		"PatientId":appointDetails.PatientId,
		"Date":appointDetails.Date})

	if err != nil {
		log.Println(err)
	}
	return  err
}

func (doc DoctorsModel) GetDoctorsById(id int64, date time.Time) (bool, error) {
    var docTimes int
    query := "SELECT COUNT(*) FROM appointments WHERE doctor_id=? AND date =?"
    err := utility.Db.QueryRow(query, id, date).Scan(&docTimes)
    if err != nil {
        log.Println(err)
        return false, err
    }
    if docTimes > 0 {
        return false, nil
    }
    return true, nil
}