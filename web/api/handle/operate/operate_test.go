package operate

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {

	b := "https://www.codedesigner.net/exec/10.79M/1902mm"
	re, _ := regexp.Compile(`[\d\.]+[mM]+`)

	//查找符合正则的第一个
	all := re.FindAll([]byte(b), -1)
	for _, item := range all {
		fmt.Println(string(item))
	}
}
