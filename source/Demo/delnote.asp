<!--#include file="db.asp"-->
<%
id=request.querystring("id")
   strsql="delete * from [note] where noteid=" & id & " or root=" & id

   
   conn.execute strsql
   response.redirect("note.asp")
%>