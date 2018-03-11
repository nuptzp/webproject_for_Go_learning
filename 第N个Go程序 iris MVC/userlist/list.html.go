
// Code generated by hero.
// source: C:\Users\hduoct\Desktop\GoWeb\userlist\list.html
// DO NOT EDIT!
package userlist

import (
	"github.com/shiyanhui/hero"
	"io"
	"strconv"
	"../Entity"
)
func UserListToWriter(userList []Entity.UserData, w io.Writer) (int, error){
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>index</title>
    <!-- Bootstrap -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style type="text/css">
        body{ font-family: Microsoft YaHei,'宋体' , Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;}
    </style>
    <script src="https://cdn.bootcss.com/markdown.js/0.6.0-beta1/markdown.min.js"></script>
</head>
<body>
    `)
	_buffer.WriteString(`
    <table class="table table-bordered">
        <tr>
            <td>编号</td><td>帐号</td><td>密码</td>
        </tr>
    `)
	for _, user := range userList {
		_buffer.WriteString(`<tr>
    <td>`)
		hero.EscapeHTML(strconv.FormatInt(user.Id,10), _buffer)
		_buffer.WriteString(`</td>
    <td>`)
		hero.EscapeHTML(user.Username, _buffer)
		_buffer.WriteString(`</td>
    <td>`)
		hero.EscapeHTML(user.Password, _buffer)
		_buffer.WriteString(`</td>
</tr>`)
	}
	_buffer.WriteString(`
    </table>
`)

	_buffer.WriteString(`
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}