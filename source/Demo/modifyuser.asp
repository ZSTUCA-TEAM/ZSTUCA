<!--#include file="db.asp"-->
<%
   on error resume next
	r=request.Form("realname")
   	s=request.Form("sex")
   	c=request.Form("city")
   	p=request.Form("phone")
   	e=request.Form("email")
   	m=request.Form("memo")
   	str = "updata [userinfo] set realname='" & r & "',sex='" & s & "',city='" & c & "',phone='" & p & "',email='"  & e & "',memo='" & m & "' where username='" & session("username") & "'"
   'response.write "<script> alert('" + str + "');</script>"
   	conn.execute str
   'response.write "<script> alert('修改成功!');window.close();</script>"
   	if conn.errors.count>0 then
		response.write "<script> alert('数据库错误!');history.go(-1)</script>"
	else 
		response.write "<script> alert('修改成功!');</script>"
	end if
%>