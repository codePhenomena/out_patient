# Doctors Appointment System Backend Application 
-> can take appointment of a doctor in a weekly scedule
->only on evenings
->In one location
->x number of patient
->without redunancy

# Run application using command: 
go run server.go 
->this commands runs server 

# To run and check application using api from postman on postmas -> hit localhost:4000/setAppointment using post request

{
    "DoctorID": 1,
    "PatientID": 1,
    "Date": "2006-01-02T17:04:05Z"
}

# To add doctors in application hit -> localhost:4000/addDoctor usisng post request 
-> it will add doctors to application

# To view all the doctors available in app hit -> localhost:4000/getDoctors using get request
->it will show list and details of available doctos
 

