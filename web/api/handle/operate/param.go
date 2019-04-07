package operate

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	pb_timestamp "github.com/golang/protobuf/ptypes/timestamp"
	pb_item "github.com/zhanglp92/rep/api/pb/item"
	pb_user "github.com/zhanglp92/rep/api/pb/user"
)

type param struct {
	r *http.Request

	ty     string
	op     string
	phone  string
	id     int32
	isover bool
}

func newParam(r *http.Request) (*param, error) {
	a := param{r: r}
	return &a, a.init()
}

func (a *param) String() string {
	m := map[string]interface{}{
		"ty":    a.ty,
		"op":    a.op,
		"phone": a.phone,
		"id":    a.id,
	}
	body, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func (a *param) init() (err error) {
	if err = a.r.ParseForm(); err != nil {
		return
	}

	a.ty = a.getStr("type")
	a.op = a.getStr("op")
	a.isover = a.getBool("isover")
	a.phone = a.getStr("phone")
	a.id = int32(a.getInt("id"))

	return nil
}

func (a *param) toUser() *pb_user.Item {
	now := time.Now()
	return &pb_user.Item{
		Name:   a.getStr("u-name"),
		Phone:  a.getStr("u-phone"),
		Sex:    a.getStr("u-sex"),
		Addr:   a.getStr("u-addr"),
		Remark: a.getStr("u-remark"),
		CreateTime: &pb_timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}
}

func (a *param) toForm() *pb_item.Item {
	item := &pb_item.Item{}
	if v := a.r.Form["obj"]; len(v) > 0 {
		json.Unmarshal([]byte(v[0]), item)
	}

	now := time.Now()
	item.Id = int32(a.getInt("f-id", int64(item.Id)))
	item.Phone = a.getStr("f-phone", item.Phone)
	item.Type = a.getStr("f-type", item.Type)
	item.Colour = a.getStr("f-colour", item.Colour)
	item.Opendirection = a.getStr("f-opendirection", item.Opendirection)
	item.Size = a.getStr("f-size", item.Size)
	item.Shape = a.getStr("f-shape", item.Shape)
	item.Glass = a.getStr("f-glass", item.Glass)
	item.Squaremeter = float32(a.getFloat("f-squaremeter", float64(item.Squaremeter)))
	item.Priceunit = float32(a.getFloat("f-priceunit", float64(item.Priceunit)))
	item.Priceset = float32(a.getFloat("f-priceset", float64(item.Priceset)))
	item.Price = float32(a.getFloat("f-price", float64(item.Price)))
	item.Remarks = a.getStr("f-remarks", item.Remarks)
	item.UpdateTime = &pb_timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	return item
}

func (a *param) getBool(key string, def ...bool) bool {
	if v := a.r.Form[key]; len(v) > 0 && v[0] == "1" {
		return true
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

func (a *param) getInt(key string, def ...int64) int64 {
	if v := a.r.Form[key]; len(v) > 0 {
		if n, err := strconv.ParseInt(v[0], 10, 0); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (a *param) getFloat(key string, def ...float64) float64 {
	if v := a.r.Form[key]; len(v) > 0 {
		if n, err := strconv.ParseFloat(v[0], 10); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (a *param) getStr(key string, def ...string) string {
	if v := a.r.Form[key]; len(v) > 0 {
		return v[0]
	}

	if len(def) > 0 {
		return def[0]
	}
	return ""
}
