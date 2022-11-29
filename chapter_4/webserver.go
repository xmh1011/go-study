package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)             // 设置访问的路由
	http.ListenAndServe("0.0.0.0:8080", nil) // 设置监听的端口
}

func handle(w http.ResponseWriter, r *http.Request) {
	keywords := []string{"goland", "java"} // 定义一个字符串切片
	result, err := SearchIssues(keywords)  // 调用 SearchIssues 函数
	if err != nil {
		log.Fatal(err) // 如果出错，打印错误信息
	}
	// 定义一个模板，用来展示结果
	// template.Must 是一个辅助函数，用来检查模板是否有错误，如果有错误，会抛出异常。
	// template.New 是一个辅助函数，用来创建一个模板。
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
	issueList.Execute(w, result) // 执行模板
}
