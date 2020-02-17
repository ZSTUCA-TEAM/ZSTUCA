<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>2019干事申请表</title>
	<link rel="stylesheet" href="./assets/css/apply.css">
	<script src="./assets/js/jquery-1.9.1.min.js"></script>
</head>
<%@Language="vbscript" Codepage="65001"%>
<%
	response.Write "<div class='data'>"
	response.Write "提交数据..."
	

if session("apply")="" then
	response.write "<script language='javascript'>alert('申请填报超时');</script>"
	response.write "<script language='javascript'>window.location.href='application.asp';</script>"
end if
	starttime=cdate("2019-9-1 9:00:00")
	endtime=cdate("2019-9-30 23:59:59")
	nowtime = now()
	response.Write(datediff("s",starttime,nowtime))
	response.Write(datediff("s",nowtime,endtime))
if datediff("s",starttime,nowtime)<0 or datediff("s",nowtime,endtime)<0 then
	response.write "<script language='javascript'>"

	response.write "alert('不在申请时间范围内，如有疑问询问浙理计协！');"
	
   	response.write "window.history.back(-1);"
	
	response.write "</script>"

else 

%>

<!--#include file="conn.asp"-->
<%
	z="here"
	if request.form<>"" then
		FOR EACH name IN Request.Form
			response.Write(request.form(name))
		NEXT
	end if
	response.Write(z) 
	response.Write(conn.state) '显示连接状态，若为1表示连接成功，若为0连接失败。
	xm=request.Form("name")
	xh=request.Form("id")
	xb=request.Form("gender")
	yx=request.Form("email")
	dh=request.Form("tel")
	qq=request.Form("qq")
	xy=request.Form("college")
	zy=request.Form("major")
	bj=request.Form("class")
	first=request.Form("select")
	secondd=request.Form("selectt")
	text=request.Form("message")
	idea=request.Form("idea")
	z="here"
	response.Write(z) 
	'rs.open "select * from 2019xsjbxx where sfzh='" & u & "' and tzsh='" & p & "'",conn,1
	response.Write(first) 

	'str = "updata [userinfo] set realname='" & r & "',sex='" & s & "',city='" & c & "',phone='" & p & "',email='"  & e & "',memo='" & m & "' where username='" & session("username") & "'"
	str ="INSERT INTO sqb SELECT '" & xm & "','" & xh & "','" & xb & "','" & yx & "','"  & dh & "','" & qq & "' ,'" & xy & "','" & zy & "','" & bj & "','" & first & "','"  & secondd & "','" & text & "','"  & idea & "'"
	response.Write(str) 
   	conn.execute str
   	if conn.errors.count>0 then
		response.write "<script> alert('数据库错误!');</script>"
	else 
		response.write "<script> alert('提交申请表成功!');</script>"
	end if
	response.write "<script language='javascript'>$(function () {	$('.popup').addClass('hide');$('.container').html('');$('.data').html('');$('.container').addClass('background');$('.container.background').lazyload();var done = $('#done').html();$('.container').append(done);});</script>"
	'response.write "<script language='javascript'>window.location.href='/';</script>"
	'rs.close
	'set rs=nothing
	conn.close
	set conn=nothing
end if
response.Write "</div>"
session("apply")=""
%>
<body style="height: 100vh">
	<div class="container" data-original="./assets/img/ground.jpg">
		<div class='popup' style='height:200px'>
			<img src='./assets/img/loading.gif'> 
			<div>正在提交...</div> 
		</div>
	</div>
</body>
<script src="./assets/js/jquery.lazyload.js"></script>
<script type="application/xml" id="done">
    <div>
        <div class="done">
            <div class="success">
                <img src="./assets/img/done.png"/>
                <div>报名成功！</div>
            </div>
            <div>请耐心等待<br/>我们的笔试/面试通知吧~</div>
        </div>
        <div class="footer">计算机协会</div>
    </div>
</script>

<script language="javascript">
	$(function () {	
		console.log($(".popup"));
		console.log($(".container"));
		function success() {
			$(".popup").addClass('hide');
			$(".container").html("");
			//$container.addClass("background");
			//$(".container.background").lazyload();
			var done = $('#done').html();
			$(".container").append(done);
		}
	});
</script>
</html>