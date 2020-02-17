<!--#include file="db.asp"-->
<%
u=request.Form("username")
p=request.Form("userpsd")
   str="select * from[userinfo] where username='" & u & "'"
	rs.open str,conn,1
	if rs.recordcount=0  then
		'response.write("<script> alert('用户名不存在!');window.close();</script>") 
	response.Redirect("Login.asp?err=用户名不存在!")
	else 
	if rs("password")<>p  then
			'response.write("<script> alert('密码不正确!');window.close();</script>") 
		response.Redirect("Login.asp?err=密码不正确!")
		else 
			session("username")=u
   			application.lock()
   			application("count")=application("count")+1
   			application.unlock()
			if(request.Form("auto")="on")  then
				response.Cookies("LoginUser")=u
				response.Cookies("LoginUser").expires=date()+365
			End if
		response.Redirect("index.asp")
		End  if
	End  if
   
%>