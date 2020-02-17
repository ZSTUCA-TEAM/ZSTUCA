<!--#include file="db.asp"-->
<%
u=request.form("username")
p=request.form("psd")
phone=request.form("phone")
'str="insert into [userinfo] content(username,password,phone) values('" & u & "','" & p & "','" & phone & "')"
'conn.execute str,1
   rs.open "select * from [userinfo]",conn,1,3
   rs.addnew
   rs("username")=u
   rs("password")=p
   rs("phone")=phone
   rs.update
   response.write("<script> alert('注册成功!');window.close();</script>")
%>