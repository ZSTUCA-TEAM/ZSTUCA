<!--#include file="db.asp"-->
<%
c=request.form("contents")
   root=request.form("root")
   rs.open "select * from [note]",conn,1,3
   rs.addnew
   rs("username")=session("username")
   rs("notetime")=now()
   rs("ip")=Request.ServerVariables("REMOTE_ADDR")
   rs("content")=c
   if root<>"" then
		rs("root")=root
	else
   		rs("root")=0
	end if
   rs.update
   response.redirect("note.asp")
%>