// Package tools
// @Description:
package tools

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

var (
	defaultSheetName = "Sheet1" //默认Sheet名称
	defaultHeight    = 25.0     //默认行高度
)

type ExcelExport struct {
	File      *excelize.File           `json:"file"`
	SheetName string                   `json:"sheetName"` //可定义默认sheet名称
	Params    []map[string]string      `json:"params"`
	Data      []map[string]interface{} `json:"data"`
	Path      string                   `json:"path"`
}

func NewMyExcel(sheetName string, tag TagBody, data interface{}) *ExcelExport {
	e := &ExcelExport{File: createFile(), SheetName: sheetName}
	for i, v := range tag.Keys {
		p := make(map[string]string)
		p["key"] = v
		p["title"] = tag.Header[i]
		p["width"] = "20"
		p["is_num"] = "1"
		e.Params = append(e.Params, p)
	}
	fmt.Printf("tag is %+v\n", tag)
	fmt.Printf("NewMyExcel is %+v\n", e)
	return e
}

// ExportToPath 导出基本的表格
func (l *ExcelExport) ExportToPath() (string, error) {
	l.export()
	name := createFileName()
	filePath := l.Path + "/" + name
	err := l.File.SaveAs(filePath)
	return filePath, err
}

// ExportToWeb 导出到浏览器。此处使用的gin框架 其他框架可自行修改ctx
func (l *ExcelExport) ExportToWeb(ctx *gin.Context) {
	l.export()
	buffer, _ := l.File.WriteToBuffer()
	//设置文件类型
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName()))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

//设置首行
func (l *ExcelExport) writeTop() {
	topStyle, _ := l.File.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	var word = 'A'
	//首行写入
	for _, conf := range l.Params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.File.SetCellValue(l.SheetName, line, title)
		//列宽
		_ = l.File.SetColWidth(l.SheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.File.SetCellStyle(l.SheetName, line, line, topStyle)
		word++
	}
}

//写入数据
func (l *ExcelExport) writeData() {
	lineStyle, _ := l.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range l.Data {
		//设置行高
		_ = l.File.SetRowHeight(l.SheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range l.Params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.File.SetCellValue(l.SheetName, line, valNum)
			} else {
				_ = l.File.SetCellValue(l.SheetName, line, val[valKey])
			}

			//设置样式
			_ = l.File.SetCellStyle(l.SheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.File.SetRowHeight(l.SheetName, len(l.Data)+1, defaultHeight)
}

func (l *ExcelExport) export() {
	l.writeTop()
	l.writeData()
}

func createFile() *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	SheetName := defaultSheetName
	index := f.NewSheet(SheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}
