package models

type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type CustomerResponse struct {
	Status  Status     `json:"status"`
	Records []Customer `json:"records"`
}

type User struct {
	ClientCode string `json:"client_code"`
	Username   string `json:"username"`
}
type Session struct {
	User       User   `json:"user"`
	SessionKey string `json:"session_key"`
}

type Record struct {
	UserID        string `json:"userID"`
	UserName      string `json:"userName"`
	EmployeeID    string `json:"employeeID"`
	EmployeeName  string `json:"employeeName"`
	GroupID       string `json:"groupID"`
	SessionKey    string `json:"sessionKey"`
	IdentityToken string `json:"identityToken"`
	Token         string `json:"token"`
}

type Status struct {
	Request         string  `json:"request"`
	RequestUnixTime int64   `json:"requestUnixTime"`
	ResponseStatus  string  `json:"responseStatus"`
	ErrorCode       int     `json:"errorCode"`
	GenerationTime  float64 `json:"generationTime"`
}

type Response struct {
	Status  Status   `json:"status"`
	Records []Record `json:"records"`
}

type TemplateData struct {
	Message string
	Error   string
	Data    map[string]interface{}
}

type GetSessionKeyInfoResponse struct {
	Status  Status             `json:"status"`
	Records []SessionKeyRecord `json:"records"`
}

type SessionKeyRecord struct {
	CreationUnixTime string `json:"creationUnixTime"`
	ExpireUnixTime   string `json:"expireUnixTime"`
}

type GetSessionKeyUserResponse struct {
	Status  Status              `json:"status"`
	Records []SessionUserRecord `json:"records"`
}

type SessionUserRecord struct {
	UserName string `json:"userName"`
}

type SaveCustomerResponse struct {
	Status  Status               `json:"status"`
	Records []SaveCustomerRecord `json:"records"`
}

type SaveCustomerRecord struct {
	ClientID      int  `json:"clientID"`
	CustomerID    int  `json:"customerID"`
	AlreadyExists bool `json:"alreadyExists"`
}
