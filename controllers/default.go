package controllers

import (
	"net/url"
	"strings"

	"hstats/models"

	"github.com/astaxie/beego"
	"github.com/qgweb/glib/timestamp"
)

type MainController struct {
	beego.Controller
}

type TjObject struct {
	AdvertId  string
	Refferer  string
	AccountID string
	Data      []string
}

func (c *MainController) Get() {
	c.Ctx.WriteString("")
}

func (c *MainController) parseUrl(name string) string {
	info, _ := url.QueryUnescape(c.GetString("d"))
	list := make(map[string]string)
	for _, v := range strings.Split(info, "; ") {
		if strings.Count(v, "=") > 1 {
			i := strings.Index(v, "=")
			list[v[:i]] = v[i+1:]
		} else {
			vv := strings.Split(v, "=")
			if len(vv) == 2 {
				list[vv[0]] = vv[1]
			}
		}
	}
	if v, ok := list[name]; ok {
		return v
	}
	return ""
}
func (c *MainController) Dispath() {
	account := c.parseUrl("txu-ac") //账号
	op := c.parseUrl("txu-op")      //统计类型
	data := c.parseUrl("txu-d")     //数据
	ref := c.parseUrl("txu-ref")    //来源
	gid := c.parseUrl("txu-gid")    //广告id

	var tinfo TjObject
	tinfo.AccountID = account
	tinfo.Refferer = ref
	tinfo.AdvertId = gid
	tinfo.Data = strings.Split(data, "/")
	beego.Info(c.Ctx.Request.Header)
	switch op {
	case "register":
		c.Op_Register(tinfo)
	case "tregister":
		c.Op_ThirdRegiser(tinfo)
	case "download":
		c.Op_Download(tinfo)
	case "consult":
		c.Op_Consult(tinfo)
	case "lmsg":
		c.Op_LeaveMessage(tinfo)
	}
	c.Ctx.WriteString("")
}

func (c *MainController) getDataFromAry(data []string, index int) string {
	if index >= len(data) {
		return ""
	} else {
		return data[index]
	}
}

// 注册
func (c *MainController) Op_Register(info TjObject) {
	models.SaveData(map[string]interface{}{
		"AdvertId":  info.AdvertId,
		"AccountID": info.AccountID,
		"Refferer":  info.Refferer,
		"Operator":  "Register",
		"UserId":    c.getDataFromAry(info.Data, 0),
		"UserName":  c.getDataFromAry(info.Data, 1),
		"Date":      timestamp.GetDayUnixStr(0),
		"Time":      timestamp.GetNow(),
	})
	beego.Warn(info)
}

// 第三方注册
func (c *MainController) Op_ThirdRegiser(info TjObject) {
	models.SaveData(map[string]interface{}{
		"AdvertId":  info.AdvertId,
		"AccountID": info.AccountID,
		"Refferer":  info.Refferer,
		"Operator":  "ThirdRegister",
		"UserId":    c.getDataFromAry(info.Data, 0),
		"UserName":  c.getDataFromAry(info.Data, 1),
		"SoftType":  c.getDataFromAry(info.Data, 2),
		"Date":      timestamp.GetDayUnixStr(0),
		"Time":      timestamp.GetNow(),
	})
	beego.Warn(info)
}

// 下载
func (c *MainController) Op_Download(info TjObject) {
	models.SaveData(map[string]interface{}{
		"AdvertId":  info.AdvertId,
		"AccountID": info.AccountID,
		"Refferer":  info.Refferer,
		"Operator":  "Download",
		"UserId":    c.getDataFromAry(info.Data, 0),
		"UserName":  c.getDataFromAry(info.Data, 1),
		"Date":      timestamp.GetDayUnixStr(0),
		"Time":      timestamp.GetNow(),
	})
	beego.Warn(info)
}

// 咨询
func (c *MainController) Op_Consult(info TjObject) {
	models.SaveData(map[string]interface{}{
		"AdvertId":  info.AdvertId,
		"AccountID": info.AccountID,
		"Refferer":  info.Refferer,
		"Operator":  "Consult",
		"UserId":    c.getDataFromAry(info.Data, 0),
		"UserName":  c.getDataFromAry(info.Data, 1),
		"Date":      timestamp.GetDayUnixStr(0),
		"Time":      timestamp.GetNow(),
	})
	beego.Warn(info)
}

// 集客
func (c *MainController) Op_LeaveMessage(info TjObject) {
	models.SaveData(map[string]interface{}{
		"AdvertId":  info.AdvertId,
		"AccountID": info.AccountID,
		"Refferer":  info.Refferer,
		"Operator":  "LeaveUser",
		"UserId":    c.getDataFromAry(info.Data, 0),
		"UserName":  c.getDataFromAry(info.Data, 1),
		"Date":      timestamp.GetDayUnixStr(0),
		"Time":      timestamp.GetNow(),
	})
	beego.Warn(info)
}
