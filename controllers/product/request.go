package product

type indexRequest struct {
	StoreName string `json:"store_name" form:"store_name" `
	CateId    string `json:"cate_id" form:"cate_id" `
	Type      int    `json:"type" form:"type" `
	Sales     string `json:"sales" form:"sales" `
}

func (req *indexRequest) Default() *indexRequest {
	req.Type = 1
	req.Sales = "normal"
	return req
}
