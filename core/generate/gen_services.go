package generate

import (
	"bytes"
	"fmt"
	"github.com/lvxin0315/gapi/core/generate/services"
	"io/ioutil"
	"reflect"
	"strings"
)

/**
 * @Author lvxin0315@163.com
 * @Description mysql生成对应model
 * @Date 11:29 上午 2020/12/11
 **/
type GenService struct {
	ServiceDir string
}

func (gen *GenService) AutoService(models ...interface{}) {
	for _, m := range models {
		moduleName := strings.ReplaceAll(reflect.TypeOf(m).Name(), "Model", "")
		gen.writeServiceFile(moduleName)
	}
}

func (gen *GenService) writeServiceFile(moduleName string) {
	modelName := fmt.Sprintf("%sModel", moduleName)
	commonServiceBytes := []byte(services.DemoServiceTpl)
	//model name
	commonServiceBytes = bytes.ReplaceAll(commonServiceBytes, []byte("DemoModel"), []byte(modelName))
	//module name
	commonServiceBytes = bytes.ReplaceAll(commonServiceBytes, []byte("Demo"), []byte(moduleName))
	//写入文件
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s_service.go", gen.ServiceDir, strings.ToLower(moduleName)), commonServiceBytes, 0755)
	if err != nil {
		panic(err)
	}
}
