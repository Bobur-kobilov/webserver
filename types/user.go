package types

type UserData struct {
	Email      string `json::"Email"`
	Pswd       string `json:"Pswd"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	OrgName    string `json:"OrgName"`
	Inst       string `json:"Inst"`
	BuildNo    string `json:"BuildNo"`
	FloorNo    string `json:"FloorNo"`
	LabHead    string `json:"LabHead"`
	LabAddress string `json:"LabAddress"`
	Tel        string `json:"Tel"`
}
