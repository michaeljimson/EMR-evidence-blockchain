package tojson

import "time"

type EmrToJson struct {
	EmrName       string    `json:"emrname"`
	ID            string    `json:"patientID"`
	DoctorName    string    `json:"doctorname"`
	DoctorID      string    `json:"doctorid"`
	EmrHash       string    `json:"emrhash"`
	EmrSignPubKey string    `json:"emrsignpubkey"`
	EmrDetail     string    `json:"emrdetial"`
	EmrTime       time.Time `json:"LoginTime"`
}
