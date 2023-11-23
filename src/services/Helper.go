package services

import (
	"strconv"
)

/**
 * @author  巨昊
 * @date  2021/7/19 20:52
 * @version 1.15.3
 */

type ItemsPage struct {
	Total   int           //一共多少条
	Current int           //当前页
	Size    int           // 页尺寸
	PageNum int           //一共多少页
	Data    []interface{} //数据
	Ext     interface{}   //扩展信息，方便插入值 给前端用
}

func (this *ItemsPage) SetExt(ext interface{}) *ItemsPage {
	this.Ext = ext
	return this
}

//@Component
type Helper struct{}

func NewHelper() *Helper {
	return &Helper{}
}

//字符串转int
func (*Helper) StrToInt(str string, def int) int {
	ret, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return ret
}

//分页 资源
func (*Helper) PageResource(current, size int, list []interface{}) *ItemsPage {
	total := len(list)
	if size == 0 || size > total {
		size = 5 //默认 每页5个
	}
	if current <= 0 {
		current = 1
	}
	pageInfo := &ItemsPage{Total: total, Size: size}
	//计算总页数
	pageNum := 1
	if pageInfo.Total > size {
		pageNum = pageInfo.Total / size
		if pageInfo.Total%size != 0 {
			pageNum++
		}
	}
	if current > pageNum {
		current = 1
	}
	pageInfo.Current = current       //重新赋值Current ----当前页
	newSet := make([]interface{}, 0) //构建一个新的 切片

	if current*size > pageInfo.Total {
		newSet = append(newSet, list[(current-1)*size:]...)
	} else {
		newSet = append(newSet, list[(current-1)*size:(current-1)*size+size]...)
	}
	//重新整理赋值
	pageInfo.Data = newSet
	pageInfo.PageNum = pageNum
	return pageInfo
}
