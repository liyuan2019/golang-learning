//Package temp は
package temp

import (
	"fmt"
	"strings"
	"time"
)

// type Count struct {
// 	Num interface{} `json:"num"`
// }

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Post struct {
	CreateTime CustomTime `json:"create_time"`
}

// func main() {
// 	// map[string]interface{} -> json string
// 	// var m = make(map[string]interface{}, 1)
// 	// m["count"] = 1 // int
// 	// var m = Count{
// 	// 	Num: 1,
// 	// }
// 	// b, err := json.Marshal(m)
// 	// if err != nil {
// 	// 	fmt.Printf("marshal failed, err:%v\n", err)
// 	// }
// 	// fmt.Printf("str:%#v\n", string(b))
// 	// var m2 Count
// 	// err = json.Unmarshal(b, &m2)
// 	// if err != nil {
// 	// 	fmt.Printf("unmarshal failed, err:%v\n", err)
// 	// 	return
// 	// }
// 	// fmt.Printf("value:%v\n", m2.Num) // 1
// 	// fmt.Printf("type:%T\n", m2.Num)  // float64

// 	p1 := Post{CreateTime: CustomTime{time.Now()}}
// 	b, err := json.Marshal(p1)
// 	if err != nil {
// 		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("str:%s\n", b)
// 	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
// 	var p2 Post
// 	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
// 		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("p2:%#v\n", p2)
// 	fmt.Printf("p2:%s\n", p2.CreateTime.Time)
// }
