package main

import (
	"bytes"
	"fmt"
	"github.com/lvxin0315/gapi/models"
	"io/ioutil"
	"reflect"
	"strings"
)

/**
 * @Author lvxin0315@163.com
 * @Description model生成对应service
 * @Date 11:29 上午 2020/12/11
 * @Param
 * @return
 **/

const ServiceDir = "services"
const DemoServiceFilePath = "services/demo_service.go"

func main() {
	autoService(models.CategoryModel{}, models.GoodsModel{})
}

func autoService(models ...interface{}) {
	for _, m := range models {
		moduleName := strings.ReplaceAll(reflect.TypeOf(m).Name(), "Model", "")
		writeServiceFile(moduleName)
	}
}

func writeServiceFile(moduleName string) {
	modelName := fmt.Sprintf("%sModel", moduleName)
	commonServiceBytes := readCommonServiceFile()
	//model name
	commonServiceBytes = bytes.ReplaceAll(commonServiceBytes, []byte("DemoModel"), []byte(modelName))
	//module name
	commonServiceBytes = bytes.ReplaceAll(commonServiceBytes, []byte("Demo"), []byte(moduleName))
	//写入文件
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s_service.go", ServiceDir, strings.ToLower(moduleName)), commonServiceBytes, 0755)
	if err != nil {
		panic(err)
	}
}

func readCommonServiceFile() []byte {
	serviceBytes, err := ioutil.ReadFile(DemoServiceFilePath)
	if err != nil {
		panic(err)
	}
	return serviceBytes
}
