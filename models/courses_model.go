package models

type Course struct {
	Course_id     string `json:"Id"`
	Course_name   string `json:"Course_name"`
	Category      string `json:"Category"`
	Length        int    `json:"Length"`
	Instructor_id string `json:"Instructor"`
	Requirements  string `json:"Requirements"`
	Description   string `json:"Description"`
}
