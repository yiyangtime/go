package main

//解码 JSON 字符串
import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact 结构代表我们的 JSON 字符串
//type Contact struct {
//	Name    string `json:"name"`
//	Title   string `json:"title"`
//	Contact struct {
//		Home string `json:"home"`
//		Cell string `json:"cell"`
//	} `json:"contact"`
//}

// JSON 包含用于反序列化的演示字符串
var JSON = `{
  "name": "Gopher",
  "title": "programmer",
  "contact": {
  "home": "415.333.3333",
  "cell": "415.555.5555"
  }
 }`

//func main() {
//	// 将 JSON 字符串反序列化到变量
//	var c Contact
//	err := json.Unmarshal([]byte(JSON), &c)
//	if err != nil {
//		log.Println("ERROR:", err)
//		return
//	}
//	fmt.Println(c)
//}

func main() {
	// 将 JSON 字符串反序列化到 map 变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
