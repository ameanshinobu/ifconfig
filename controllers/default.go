package controllers

import (
	// "bytes"
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"os"
	"strings"
	// "unicode/utf8"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) GetForwarded() {
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetHost() {
	ip := this.Ctx.Request.RemoteAddr

	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		this.Data["Value"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		this.Data["Value"] = value
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetIP() {
	this.Data["Value"] = getIP(this.Ctx.Request.RemoteAddr)
	this.TplName = "value.tpl"
}

func (this *MainController) GetPort() {
	this.Data["Value"] = getPort(this.Ctx.Request.RemoteAddr)
	this.TplName = "value.tpl"
}

func (this *MainController) GetVia() {
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Via"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetMime() {
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetLang() {
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetCharset() {
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Charset"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetEncoding() {
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetUserAgent() {
	this.Data["Value"] = this.Ctx.Request.UserAgent()
	this.TplName = "value.tpl"
}

func (this *MainController) GetConnection() {
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Connection"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetKeepAlive() {
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetAll() {
	this.Data["Email"] = "admin@ifconfig.tw"
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	ip := getIP(this.Ctx.Request.RemoteAddr)
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		this.Data["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		this.Data["Host"] = value
	}
	this.Data["IP"] = getIP(this.Ctx.Request.RemoteAddr)
	this.Data["Port"] = getPort(this.Ctx.Request.RemoteAddr)
	this.Data["Method"] = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		this.Data["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		this.Data["Mime"] = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		this.Data["Connection"] = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		this.Data["Via"] = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		this.Data["Charset"] = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		this.Data["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		this.Data["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		this.Data["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	this.Data["Referer"] = this.Ctx.Input.Refer()

	this.TplName = "all.tpl"
}

type ifconfig struct {
	Email      string
	UserAgent  string
	Host       string
	IP         string
	Port       string
	Method     string
	Encoding   string
	Mime       string
	Connection string
	Via        string
	Charset    string
	Keepalive  string
	Forwarded  string
	Lang       string
	Referer    string
}

func (this *MainController) GetAllXML() {
	thisData := ifconfig{}
	thisData.Email = "admin@ifconfig.tw"
	thisData.UserAgent = this.Ctx.Request.UserAgent()

	ip := getIP(this.Ctx.Request.RemoteAddr)
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		thisData.Host = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		thisData.Host = value
	}

	thisData.IP = getIP(this.Ctx.Request.RemoteAddr)
	thisData.Port = getPort(this.Ctx.Request.RemoteAddr)
	thisData.Method = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		thisData.Encoding = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		thisData.Mime = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		thisData.Connection = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		thisData.Via = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		thisData.Charset = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		thisData.Keepalive = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		thisData.Forwarded = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		thisData.Lang = this.Ctx.Request.Header["Accept-Language"][0]
	}
	thisData.Referer = this.Ctx.Input.Refer()

	this.Data["xml"] = thisData
	this.ServeXML()
}

func (this *MainController) GetAllJSON() {
	thisData := make(map[string]interface{})
	thisData["Email"] = "admin@ifconfig.tw"
	thisData["UserAgent"] = this.Ctx.Request.UserAgent()
	ip := getIP(this.Ctx.Request.RemoteAddr)
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		thisData["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		thisData["Host"] = value
	}

	thisData["IP"] = getIP(this.Ctx.Request.RemoteAddr)
	thisData["Port"] = getPort(this.Ctx.Request.RemoteAddr)
	thisData["Method"] = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		thisData["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		thisData["Mime"] = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		thisData["Connection"] = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		thisData["Via"] = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		thisData["Charset"] = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		thisData["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		thisData["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		thisData["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	thisData["Referer"] = this.Ctx.Input.Refer()

	this.Data["json"] = thisData
	this.ServeJSON()
}

func (this *MainController) Get() {
	if noweb := os.Getenv("NOWEB"); noweb == "1" {
		this.Abort("404")
		return
	}
	this.Data["Email"] = "admin@ifconfig.tw"
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	ip := getIP(this.Ctx.Request.RemoteAddr)
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		this.Data["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		this.Data["Host"] = value
	}

	this.Data["IP"] = getIP(this.Ctx.Request.RemoteAddr)
	this.Data["Port"] = getPort(this.Ctx.Request.RemoteAddr)
	this.Data["Method"] = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		this.Data["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		this.Data["Mime"] = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		this.Data["Connection"] = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		this.Data["Via"] = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		this.Data["Charset"] = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		this.Data["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		this.Data["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		this.Data["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	this.Data["Referer"] = this.Ctx.Input.Refer()

	if strings.Contains(this.Ctx.Request.UserAgent(), "curl") {
		this.TplName = "iponly.tpl"
	} else {
		this.TplName = "index.tpl"
	}
}

func getIPVersion(s string) string {
	// rs, _ := utf8.DecodeRuneInString(s)
	for _, r := range s {
		if r == ':' {
			return "ipv6"
		} else if r == '.' {
			return "ipv4"
		}
	}
	return "ipv4"
}

func getIP(s string) string {
	if getIPVersion(s) == "ipv6" {
		pos := strings.Index(s, "]")
		if pos == -1 {
			return s
		} else {
			return s[1:pos] // remove [ ]
		}
	} else {
		// ipv4
		pos := strings.Index(s, ":")
		return s[0:pos]
	}
}

func getPort(s string) string {
	if getIPVersion(s) == "ipv6" {
		pos := strings.Index(s, "]")
		if pos == -1 {
			return ""
		}
		return s[pos+2:]
	} else {
		// ipv4
		pos := strings.Index(s, ":")
		if pos == -1 {
			return ""
		}
		return s[pos+1:]

	}
}
