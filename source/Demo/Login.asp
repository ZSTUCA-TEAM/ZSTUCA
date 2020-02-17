<!doctype html>
<div>
  <html>
  <head>
  <meta charset="gb18030">
  <title>Login</title>
  <style type="text/css">
  </style>
  <link href="democss.css" rel="stylesheet" type="text/css">
	  <script src="js/jquery-3.3.1.js"></script>
  </head>
	
	  
	  
<script language="javascript">
	function login(){
		if(document.getElementById('username').value==''){
		 	document.getElementsByClassName('login_lbl').item(0).style.visibility='visible';
			//document.getElementById('errinfo').style.visibility='visible';
			return;
		}else{
			document.getElementsByClassName('login_lbl').item(0).style.innerHTML='√';
			document.getElementsByClassName('login_lbl').item(0).style.color='#0F0';
			document.getElementsByClassName('login_lbl').item(0).style.visibility='visible';
		}
		if(document.getElementById('userpsd').value==''){
		 	document.getElementsByClassName('login_lbl').item(1).style.visibility='visible';
			//document.getElementById('errinfo').style.visibility='visible';
			return;
		}else{
			document.getElementsByClassName('login_lbl').item(1).style.innerHTML='×';
			document.getElementsByClassName('login_lbl').item(1).style.color='#FF';
			document.getElementsByClassName('login_lbl').item(1).style.visibility='visible';
			
		}
		document.forms(0).submit();
	}
	function register(){
		window.location.href='reg.html';
	}
	function keylogin(){
		if(window.Event.keyCode==13)
			login();
	}
	function Closeerrdiv(){
		document.getElementById('errinfo').style.visibility='hidden';
		
	}
</script>
	  
<body>
<%
   	if request.Cookies("LoginUser")<>"" then
		session("loginUser")=request.Cookies("LoginUser")
		application.lock()
   		application("count")=application("count")+1
   		application.unlock()
		response.Redirect("index.asp")
	end if
	
	er=request.QueryString("err")	
	if er<>"" then 			
%>
	<div id="errinfo"><%=er%><span id="closeerr" onClick="Closeerrdiv();">&times;</span></div>
<%
	End if					
%>	  
	  
<div id="logo"></div>
  <div id="board">
    <div id="login_frm">
      <form name="form1" method="post" action="checklogin.asp" onKeyPress="keylogin();">
        <div class="textbox">
          用户名：<input name="username" type="text" required id="username" placeholder="请输入用户名"><span class="login_lbl" >*必填</span>
        </div>
        <div class="textbox">
          密&ensp;&ensp;码：<input name="userpsd" type="password" required id="userpsd" placeholder="请输入密码"><span class="login_lbl" >*必填</span>
        </div>
        <div class="textbox">
          <input name="auto" type="checkbox" id="auto">自动登录
        </div>
        <div class="textbox">
          <input type="submit" autofocus="autofocus" class="login_btm" onClick="login();" ><input type="submit" autofocus="autofocus" class="login_btm" onClick="register();" value="注册" >
        </div>
      </form>
    </div>
  </div>
    
  <div id="caption">CopyRight Demo.All right reserve</div>
    
    
  </body>
  </html>
</div>
