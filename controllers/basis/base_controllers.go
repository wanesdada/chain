package basis

import (
	"chain/utils"
	"chain/wanzlog"
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	"net/url"
)

var apiLog *log.Log

func init() {
	apiLog = log.Init("20060102.api")
}

type Matedata struct {
	CoinName string
	CoinKey  string
}

type BlockBaseControllers struct {
	beego.Controller
	m *Matedata
}

func (c *BlockBaseControllers) Prepare() {
	CoinName := c.Ctx.Input.Header("CoinName")

	if CoinName == "" {
		c.Abort("408")
		c.StopRun()
	}
	//CoinKey := c.Ctx.Input.Header("CoinKey")
	//if CoinKey == "" {
	//	c.Abort("408")
	//	c.StopRun()
	//}

	c.m = &Matedata{CoinName: CoinName, CoinKey: ""}
}

func (c *BlockBaseControllers) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = false
		hasEncoding = false
	)
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	c.JSON(c.Data["json"], hasIndent, hasEncoding)
}

// JSON writes json to response body.
// if coding is true, it converts utf-8 to \u0000 type.
func (c *BlockBaseControllers) JSON(data interface{}, hasIndent bool, coding bool) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(c.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	ip := c.Ctx.Input.IP()
	requestBody, _ := url.QueryUnescape(string(c.Ctx.Input.RequestBody))
	apiLog.Println("URI:", c.Ctx.Input.URI(), "RequestBody:", requestBody, "ResponseBody:", string(content), "IP:", ip, "CoinName:", c.m.CoinName)
	if coding {
		content = []byte(utils.StringsToJSON(string(content)))
	}
	return c.Ctx.Output.Body(content)
}
