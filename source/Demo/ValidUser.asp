<!--#include file="db.asp"-->

<% 
	
	dim u	
	u=request.QueryString ("username")

	
	str="select * from[userinfo] where username='" & u & "'"
	rs.open str,conn,1
	if rs.recordcount>0  then
		response.write "{""flag"":false}"
	else 
		response.write "{""flag"":true}"
	End  if
%>