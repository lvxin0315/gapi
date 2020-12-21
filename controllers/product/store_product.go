package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/core/request"
)

//Route::get('product', 'v1.product.StoreProduct/index');
func Index(c *gin.Context) {
	params := indexRequest{}
	err := request.ShouldBindAndResponse(c, &params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(params.CateId)
	fmt.Println(params.StoreName)
	fmt.Println(params.Type)
	fmt.Println(params.Sales)
}
