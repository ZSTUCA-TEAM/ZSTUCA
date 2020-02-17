	<DIV id=navigation>
        <UL>
          <LI><A href="index.asp">Home</A> </LI>
          <LI><A href="#">我的好友</A></LI> 
          <LI><A href="#" >我的投票</A> </LI>
		  <LI><A href="note.asp"  target=_self>给我留言</A> </LI>
          <LI></LI>
          <LI><A href="#">我的收藏</A> </LI>
			<A href="#" onclick="window.open ('edituser.asp', 'newwindow', 'height=600, width=800, top=0, left=0, toolbar=no, menubar=no, scrollbars=no, resizable=no,location=no, status=no') "><%=session("username")%></A>
		  <LI><a href="#">修改密码</a></LI>
          <LI><A href="#" onClick="exit();" >退出</A> </LI>          
  </UL>        
	</DIV>