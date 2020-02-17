<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3c.org/TR/1999/REC-html401-19991224/loose.dtd">
<HTML xmlns="http://www.w3.org/1999/xhtml">
<HEAD >
<TITLE>我的首页</TITLE>
<META content="text/html; charset=gb2312" http-equiv=Content-Type>
<LINK rel=stylesheet type=text/css href="style/index.css">
</HEAD>
<%
	if session("username")="" then
   		response.Redirect("Login.asp?err=会话超时!")
   	end if
%>
<%
browserstr= Request.ServerVariables("HTTP_USER_AGENT")
'response.write "<script>alert( '"+browserstr+"' );</script>"
if instr(browser,"MSIE") then
	response.write "<script>alert('您的浏览器是IE浏览器，建议更换浏览器！');</script>"
end if
%>
														
<script language="javascript">
	function exit(){
		if(confirm('您真的要退出吗?')){
			window.location.href="logout.asp"
			
		}
		
	}
	
</script>
<BODY>
<DIV id=page>
    <DIV id=header>
        <DIV id=headerimg>
        <H1>我的首页</H1>
        <DIV class=description>My Personal WebSite</DIV>
        </DIV>
    </DIV>
	<!--#include file="menu.asp"-->
	<HR>
    <DIV id=content class=widecolumn>
		<marquee><H2>您的IP地址是:<%=Request.ServerVariables("REMOTE_ADDR")%></H2></marquee>
        <DIV class="postmetadata alt">
			<P>您是本网站的第<%=application("count")%>位访问用户!</P>
            <p>当前在线人数：<%=application("onlineUser")%></p>
            <p>&nbsp;</p>
        </DIV>
    </DIV>

	<HR>
    <DIV id=footer>
    <P>Copyright&copy; 2019 My.Website<BR><a href="mailto:webmaster@xyz.com.cn">Contact Me:Administrator</A> 
    </P>
    </DIV>
</DIV>
</BODY>
</HTML>
