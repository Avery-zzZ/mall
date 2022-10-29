package define

type Res_success struct{
	Code int `json:"code" example:"1"`
	Msg string `json:"msg" example:"success"`
}

type Res_login_success_data struct{
	Grade int `json:"grade" example:"2"`
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidmlwMiIsImdyYWRlIjoyLCJleHAiOjE2NjcwMTI3MDN9.HPfE71HBW-TVazslSxvJ87bNwp3Ra_zluJGoABQ3nPk"`
}

type Res_login_success struct{
	Code int `json:"code" example:"1"`
	Data Res_login_success_data `json:"data"`
}

type Res_get_success_data struct{
	Count int `json:"count" example:"1"`
	List interface{} `json:"list"` 
}

type Res_get_success struct{
	Code int `json:"code" example:"1"`
	Data Res_get_success_data `json:"data"`
}