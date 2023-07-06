package main

import (
	"fmt"

	"encoding/json"

	"example.com/greetings"
)

func main() {
	// 获取问候信息并打印出来.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

// 集合demo
func testMap() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "tom"

	var p2 map[int]string = map[int]string{}
	p2[1] = "jsy"

	var p3 map[int]string = make(map[int]string)
	p3[1] = ""

	p4 := map[int]string{}
	p4[1] = ""

	p5 := make(map[int]string)
	p5[1] = ""

	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "tom",
		"age":      "20",
		"hobby":    []string{"读书", "爬山"},
	}

	// 序列化
	jsons, _ := json.Marshal(res)
	fmt.Println(string(jsons))

	// 反序列化
	res2 := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsons), &res2)

	// 删除
	delete(p1, 1)
}

// 结构体demo
func testStruct() {
	type Person struct {
		Name string
	}

	var p1 Person
	p1.Name = ""

	var _ = Person{Name: ""}

	type Result struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	res := Result{Code: 200, Message: "success"}
	// 序列化
	jsons, _ := json.Marshal(res)
	fmt.Println(jsons)

	// 反序列化
	var res2 Result
	json.Unmarshal(jsons, &res2)
}

func demoSlice() {
	var sli_1 []int
	// var sli_2 = []int{}
	var sli_3 = []int{1, 2, 3}
	sli_4 := []int{1, 2, 3, 5}

	sli_1 = append(sli_1, 2)
	// 删除尾部2个元素
	sli_3 = sli_3[:len(sli_3)-2]
	// 删除开头2个元素
	sli_4 = sli_4[2:]
	// 删除中间2个元素
	sli_4 = append(sli_4[:1], sli_4[1+2:]...)

}

func demoChan() {
	// chan 可以理解为队列，遵循先进先出的规则。

	// 不带缓冲的通道(进出都会阻塞)
	ch1 := make(chan string)
	// 带10个缓冲的通道（进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。）
	ch2 := make(chan string, 10)
	// 只读通道
	// ch3 := make(<-chan string)
	// 只写通道
	ch4 := make(chan<- string)

	ch1 <- "a"
	ch2 <- "b"

	val, _ := <-ch2
	fmt.Println(val)
	//     close 以后不能再写入，写入会出现 panic
	// 重复 close 会出现 panic
	// 只读的 chan 不能 close
	// close 以后还可以读取数据
	close(ch4)
}
