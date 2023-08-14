文章转载自 https://www.cnblogs.com/f-ck-need-u/p/10053124.html

学习一下   内涵一些学习笔记

---

### 关于. 和作用域

在写template的时候，会经常用到"."。比如`{{.}}`、`{{len .}}`、`{{.Name}}`、`{{$x.Name}}`等等。

在template中，点"."代表**当前作用域的当前对象**。它类似于java/c++的this关键字，类似于perl/python的self。如果了解perl，它更可以简单地理解为默认变量`$_`。

例如，前面示例test.html中`{{.}}`，这个点是顶级作用域范围内的，它代表`Execute(w,"hello worold")`的第二个参数"hello world"。也就是说它代表这个字符串对象。

```go
package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`))
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

//hello longshuai!
//
//an email a1@qq.com
//an email a2@gmail.com
//
//my friend name is xiaofang
//my friend name is wugui
```
这里定义了一个Person结构，它有两个slice结构的字段。`在Parse()`方法中：

- 顶级作用域的`{{.UserName}}`、`{{.Emails}}`、`{{.Friends}}`中的点都代表Execute()的第二个参数，也就是Person对象p，它们在执行的时候会分别被替换成p.UserName、p.Emails、p.Friends。
- 因为Emails和Friend字段都是可迭代的，在`{{range .Emails}}...{{end}}`这一段结构内部an email `{{.}}`，这个"."代表的是range迭代时的每个元素对象，也就是p.Emails这个slice中的每个元素。
- 同理，with结构内部`{{range .}}`的"."代表的是p.Friends，也就是各个，再此range中又有一层迭代，此内层`{{.Fname}}`的点代表Friend结构的实例，分别是&f1和&f2，所以`{{.Fname}}`代表实例对象的Fname字段。


这里template的操作就自己看吧  我们主要聚焦于

text/template 和 html/template


