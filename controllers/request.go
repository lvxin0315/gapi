package controllers

type categoryRequest struct {
	Pid   int    `json:"pid"  form:"pid"`
	Name  string `json:"name"  form:"name"`
	Limit int    `json:"limit"  form:"limit"`
}

func (req *categoryRequest) Default() {
	req.Pid = -1
}
