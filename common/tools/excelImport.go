package tools

import (
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"io"
	"reflect"
	"strconv"
	"time"
)

type ExcelData interface {
	CreateMap(arr []string) map[string]interface{}
	ChangeTime(source string) time.Time
}

type ExcelStruct struct {
	Temp  [][]string
	Model interface{}
	Info  []map[string]string
}

func (excel *ExcelStruct) ReadExcel(file string) *ExcelStruct {

	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println("read excel:", err)
	}
	var rows [][]string
	sheets := xlsx.GetSheetList()
	for _, s := range sheets {
		row, _ := xlsx.GetRows(s)
		rows = append(rows, row...)
	}
	excel.Temp = rows

	return excel

}

func (excel *ExcelStruct) ReadExcelIo(file io.Reader) error {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	sheets := xlsx.GetSheetList()
	rows, _ := xlsx.GetRows(sheets[0])
	//忽略标题行
	excel.Temp = rows[1:]
	log.Debugf("excel is %+v\n", excel)
	return err
}
func (excel *ExcelStruct) CreateMap() error {
	tag := GetTag(excel.Model)
	//利用反射得到字段名
	for _, v := range excel.Temp {
		//将数组  转成对应的 map
		var info = make(map[string]string)
		for i := 0; i < reflect.TypeOf(excel.Model).Elem().NumField(); i++ {
			obj := reflect.TypeOf(excel.Model).Elem().Field(i)
			for j, h := range tag.Field {
				//log.Debugf("h:%s--obj.Name：%s \n", h, obj.Name)
				if obj.Name == h {
					log.Debugf("j is %v \n", j)
					log.Debugf("tag.Keys is %v \n", tag.Keys)
					log.Debugf("v is %v \n", v)
					info[tag.Keys[j]] = v[j]
					//info[obj.Name] = v[i]
				}
			}
		}
		excel.Info = append(excel.Info, info)
	}
	log.Debugf("CreateMap excel.Info is %+v\n", excel.Info)
	return nil
}

func (excel *ExcelStruct) ChangeTime(source string) time.Time {
	ChangeAfter, err := time.Parse("2006-01-02", source)
	if err != nil {
		log.Fatalf("转换时间错误:%s", err)
	}
	return ChangeAfter
}

func (excel *ExcelStruct) SaveDb(db *gorm.DB) (err error) {
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Unscoped().Where("1=1").Delete(&excel.Model).Error
	if err != nil {
		return err
	}
	//var temp []interface{}

	for i := 0; i < len(excel.Info); i++ {
		t := reflect.ValueOf(excel.Model).Elem()
		log.Debugf("t is %+v\n", t)
		for k, v := range excel.Info[i] {
			log.Debugf(" excel.Info[i] is %+v\n", excel.Info[i])
			log.Debugf(" k is %+v\n", k)
			log.Debug(t.FieldByName(k))
			//log.Debug(t.Elem().FieldByName(k).Kind())
			log.Debugf("key:%v---val:%v", t.FieldByName(k), t.FieldByName(k).Kind())

			switch t.FieldByName(k).Kind() {
			case reflect.String:
				t.FieldByName(k).Set(reflect.ValueOf(v))
			case reflect.Float64:
				tempV, err := strconv.ParseFloat(v, 64)
				if err != nil {
					log.Errorf("string to float64 err：%v", err)
					return err
				}
				t.FieldByName(k).Set(reflect.ValueOf(tempV))
			case reflect.Uint64:
				reflect.ValueOf(v)
				tempV, err := strconv.ParseUint(v, 0, 64)
				if err != nil {
					log.Errorf("string to uint64 err：%v", err)
					return err
				}
				t.FieldByName(k).Set(reflect.ValueOf(tempV))

			case reflect.Struct:
				tempV, err := time.Parse("2006-01-02", v)
				if err != nil {
					return err
				}
				t.FieldByName(k).Set(reflect.ValueOf(tempV))
			default:
				log.Errorf("type err")
			}
		}
		t.FieldByName("Id").Set(reflect.ValueOf(0))
		log.Debugf("model is %+v\n", excel.Model)
		//temp = append(temp, excel.Model)
		err = tx.Create(excel.Model).Error
		if err != nil {
			log.Errorf("type err is %v\n", err)
			return err
		}

	}

	return nil
}
