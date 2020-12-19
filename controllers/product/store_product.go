package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/controllers"
)

//Route::get('product', 'v1.product.StoreProduct/index');
func Index(c *gin.Context) {
	params := indexRequest{}
	err := c.ShouldBind(params.Default())
	if err != nil {
		controllers.Fail(c, err)
		return
	}
	fmt.Println(params)
}
