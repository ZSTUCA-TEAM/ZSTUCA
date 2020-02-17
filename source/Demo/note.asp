<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<HTML xmlns="http://www.w3.org/1999/xhtml">
<HEAD >
<TITLE>????????</TITLE>
<META content="text/html; charset=GB2312" http-equiv=Content-Type>
<LINK rel=stylesheet type=text/css href="style/index.css">
</HEAD>
<script type="application/javascript">
	
	function res(){
		if(document.getElementById('pageon').value!='')
			{
				
			}else{
				
				
			}
	}
	
	function goto(){
		if(document.getElementById('contents').value!='')
			{
				
			}else{
				
				
			}
	}
	
</script>
<BODY>
<script language="javascript">
	function exit(){
		if(confirm('?ú??????????????')){
			window.location.href="Logout.asp";
			}
		
		}
</script>
<%
if session("username")="" then
	response.Redirect("Login.asp?errinfo=?á?°???±????????????!")
end if
%>
<DIV id=page>
    <DIV id=header>
        <DIV id=headerimg>
        <H1>????°?</H1>
        <DIV class=description>My Personal WebSite</DIV>
        </DIV>
    </DIV>
	<!--#include file="menu.asp"-->
	
	<HR>
		<!--#include file="db.asp"-->
    <DIV id=content class=widecolumn>
        <H2>????°?</H2>
		<%
		   strsql="select * from [note] where root=0 order by noteid ASC"
		   rs.open strsql,conn,1
		   if rs.recordcount=0 then
		   %>
        <DIV class="postmetadata alt">
        	??????????     
			<hr  style='display:block'>
        </DIV> 
    	<%
		   else
		   		i=0
		   		rs.pagesize=5
		   		no=request.querystring("page")
		   		if no="" then
		   			no=1
		   		end if
		   		rs.absolutepage=no
		   		do until rs.eof or i=rs.pagesize
		   			i=i+1
		   			response.write i & "??#" & rs("username") & "|" & rs("ip") & "|" & rs("notetime") 
		   			if session("username")="admin" or session("username")=rs("username") then
		   				response.write "??<a href=delnote.asp?id=" & rs("noteid") & ">????</a>??"
		   			end if
		   			response.write "??<a href=note.asp?resid=" & i & "&root=" & rs("noteid") & ">????</a>??"
		   			response.write	"<br>"
		   			response.write rs("content")
		   			response.write	"<br>"
		   			set rs1=server.CreateObject("ADODB.Recordset")
					strsql1="select * from [note] where root=" & rs("noteid") & " order by noteid ASC"
		   			rs1.open strsql1,conn,1
		   			if rs1.recordcount>0 then
		   				do until rs1.eof
							response.write "????????" & i & "??#" & rs1("username") & "|" & rs1("ip") & "|" & rs1("notetime") 
							if session("username")="admin" or session("username")=rs1("username") then
							response.write "??<a href=delnote.asp?id=" & rs1("noteid") & ">????</a>??"
							
		   					end if
							response.write "??<a href=note.asp?resid=" & i & "&root=" & rs1("noteid") & ">????</a>??"
							response.write	"<br>"
		   					response.write rs1("content")
							response.write	"<br>"
		   					rs1.movenext
		   				loop
		   			end if
		   			response.write "<hr style='display:block'>"
		   			rs.movenext
		   		loop
		   end if
		   %>
      <center>
      ??????  
	    <%
			 for i=1 to rs.pagecount
			 	response.write "<a href=note.asp?page=" & i & ">" & i & "</a>&nbsp;&nbsp;"
			 next
			 %>
      <input name="pageno" type="text" id="pageno" size="1">??
		  
			
      <form name="form1" id="form1" method="post" action="addnote.asp" >
		<%
			 res=request.querystring("resid")
			 rt=request.querystring("root")
			 if res<>"" then
			 %>
      	<br><label name="root" id="root" value=<%=res%> >??????<%=res%>????</label><br>
		<%
			 end if
			 %>  
        <textarea name="contents" id="contents" cols="45" rows="5" style="text-align:
        left; font-family:'???í????'"></textarea>
        <br>
        <input type="submit" name="button" id="button" value="????"  style="font-family:'???í????'">
        <input type="reset" name="button2" id="button2" value="????" style="font-family:'???í????'">
      </form>
      </center>
    </DIV>

	<HR>
    <DIV id=footer>
    <P>Copyright&copy; 2019 My.Website<BR><a href="mailto:webmaster@xyz.com.cn">Contact Me:Administrator</A> 
    </P>
    </DIV>
</DIV>
</BODY>
</HTML>
