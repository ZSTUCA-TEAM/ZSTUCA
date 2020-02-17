<%
	session.Abandon()
	response.Cookies("LoginUser")=""
	response.Redirect("login.asp")
%>