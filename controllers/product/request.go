package product

import "errors"

type indexRequest struct {
	StoreName string `json:"store_name" form:"store_name" `
	CateId    int    `json:"cate_id" form:"cate_id" binding:"required"`
	Type      int    `json:"type" form:"type" `
	Sales     string `json:"sales" form:"sales"`
}

func (req *indexRequest) Default() {
	req.Type = 1
}

func (req *indexRequest) CustomVerification() error {
	if req.StoreName == "demo" {
		return errors.New("StoreName is Demo")
	}
	return nil
}
