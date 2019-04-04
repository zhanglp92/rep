package operate

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type param struct {
	r *http.Request

	ty  string
	op  string
	id  int32
	obj []byte
}

func newParam(r *http.Request) (*param, error) {
	a := param{r: r}
	return &a, a.init()
}

func (a *param) String() string {
	m := map[string]interface{}{
		"ty":  a.ty,
		"op":  a.op,
		"id":  a.id,
		"obj": a.obj,
	}
	body, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func (a *param) parseObj(v interface{}) error {
	return json.Unmarshal(a.obj, v)
}

func (a *param) init() (err error) {
	if err = a.r.ParseForm(); err != nil {
		return
	}

	if v := a.r.Form["obj"]; len(v) > 0 {
		a.obj = []byte(v[0])
	}

	if v := a.r.Form["type"]; len(v) > 0 {
		a.ty = v[0]
	}

	if v := a.r.Form["op"]; len(v) > 0 {
		a.op = v[0]
	}

	if v := a.r.Form["id"]; len(v) > 0 {
		n, err := strconv.ParseInt(v[0], 10, 0)
		if err != nil {
			return err
		}
		a.id = int32(n)
	}

	return nil
}
