<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" />
<title>吉林省明日科技有限公司</title>
<link href="css/style.css" rel="stylesheet" type="text/css" />
<link href="css/index.css" rel="stylesheet" type="text/css" />
<LINK rel=stylesheet type=text/css href="css/base.css">
</head>
<script language="javascript" src="js/AjaxRequest.js"></script>
<script language="javascript">
/******************错误处理的方法*******************************/
function onerror(){
	alert("您的操作有误！");
}
/******************实例化Ajax对象的方法*******************************/
function getInfo(){
	var loader=new net.AjaxRequest("check.php?nocache="+new Date().getTime(),deal_getInfo,onerror,"GET");
}
/************************回调函数**************************************/
function deal_getInfo(){
	document.getElementById("showInfo").innerHTML=this.req.responseText;
}
window.onload=function(){
	getInfo();	//调用getInfo()方法获取最新消息
	window.setInterval("getInfo()", 600000);	//每隔10分钟调用一次getInfo()方法
}
</script>
<body>
<div class="cen">
  <div class="c01">
    <div class="cleft"><img src="img/i01.gif" width="136" height="115" /></div>
    <div class="cright">
      <div>
        <table width="400" border="0" cellspacing="0" cellpadding="0">
          <tr>
            <td width="#" height="105" align="right" valign="top"><img src="img/i02.gif" width="97" height="98" /></td>
            <td width="280" align="right" valign="top"><img src="img/i03.gif" width="269" height="86" /></td>
          </tr>
        </table>
      </div>
      <div><img src="img/i04.gif" width="249" height="16" /></div>
    </div>
  </div>
  
  <div>
    <script language="JavaScript" type="text/javascript">
function showadv(par,par2,par3)
{
document.getElementById("a0").style.display = "none";
document.getElementById("a0color").style.color = "";
document.getElementById("a0bg").style.backgroundImage="";	
document.getElementById("a1").style.display = "none";
document.getElementById("a1color").style.color = "";
document.getElementById("a1bg").style.backgroundImage="";
document.getElementById("a2").style.display = "none";
document.getElementById("a2color").style.color = "";
document.getElementById("a2bg").style.backgroundImage="";
document.getElementById("a3").style.display = "none";
document.getElementById("a3color").style.color = "";
document.getElementById("a3bg").style.backgroundImage="";
document.getElementById("a4").style.display = "none";
document.getElementById("a4color").style.color = "";
document.getElementById("a4bg").style.backgroundImage="";
document.getElementById("a5").style.display = "none";
document.getElementById("a5color").style.color = "";
document.getElementById("a5bg").style.backgroundImage="";
document.getElementById(par).style.display = "";	
document.getElementById(par2).style.color = "#ffffff";
document.getElementById(par3).style.backgroundImage = "url(img/i13.gif)";	
}
  </script>
    <table cellspacing="0" cellpadding="0" width="100%" border="0">
      <tr>
        <td><div class="i01w">
            <table cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td width="#" height="42" align="center" id="a0bg"><span id="a0color" onmouseover='showadv(&quot;a0&quot;,&quot;a0color&quot;,&quot;a0bg&quot;)'><a href="index.php"><font color="#FA4A05">首页</font></a></span></td>
                <td width="1"><img src="img/i14.gif" width="1" height="25" /></td>
                <td id="a1bg" align="center" width="144"><span id="a1color" onmouseover='showadv(&quot;a1&quot;,&quot;a1color&quot;,&quot;a1bg&quot;)'><a href="gyld.html">关于明日</a></span></td>
                <td width="1"><img src="img/i14.gif" width="1" height="25" /></td>
                <td id="a2bg" align="center" width="144"><span id="a2color" onmouseover='showadv(&quot;a2&quot;,&quot;a2color&quot;,&quot;a2bg&quot;)'><a href="#">明日图书</a></span></td>
                <td width="1"><img src="img/i14.gif" width="1" height="25" /></td>
                <td id="a3bg" align="center" width="144"><span id="a3color" onmouseover='showadv(&quot;a3&quot;,&quot;a3color&quot;,&quot;a3bg&quot;)'><a href="#">明日编程词典</a></span></td>
                <td width="1"><img src="img/i14.gif" width="1" height="25" /></td>
                <td id="a4bg" align="center" width="144"><span id="a4color" onmouseover='showadv(&quot;a4&quot;,&quot;a4color&quot;,&quot;a4bg&quot;)'><a href="#">明日软件</a></span></td>
                <td width="1"><img src="img/i14.gif" width="1" height="25" /></td>
                <td id="a5bg" align="center" width="144"><span id="a5color" onmouseover='showadv(&quot;a5&quot;,&quot;a5color&quot;,&quot;a5bg&quot;)'><a href="gyld.html">加入我们</a></span></td>
              </tr>
            </table>
        </div></td>
      </tr>
      <tr>
        <td><table width="100%" height="41" cellpadding="0" cellspacing="0" id="a0"　border="0">
            <tr>
              <td align="left" style="padding-left:12px">欢迎来到吉林省明日科技有限公司</td>
            </tr>
          </table>
            <table id="a1" style="DISPLAY: none" height="41" cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td  style="padding-left:90px" align="left"><ul class="i02w">
                    <li>明日团队</li>
                  <li>明日历史</li>
                  <li>明日简介</li>
                </ul></td>
              </tr>
            </table>
          <table id="a2" style="DISPLAY: none" height="41" cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td style="padding-left:270px" align="left"><ul class="i02w">
                    <li><a href="#">精品图书</a></li>
                  <li><a href="#">热销图书</a></li>
                </ul></td>
              </tr>
          </table>
          <table id="a3" style="DISPLAY: none" height="41" cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td style="padding-left:12px">欢迎来到吉林省明日科技有限公司</td>
              </tr>
          </table>
          <table id="a4" style="DISPLAY: none" height="41" cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td style="padding-left:12px">欢迎来到吉林省明日科技有限公司</td>
              </tr>
          </table>
          <table id="a5" style="DISPLAY: none" height="41" cellspacing="0" cellpadding="0" width="100%" border="0">
              <tr>
                <td style="padding-right:10px"><ul class="i03w">
                  <li>联系我们</li>
                  <li>诚聘英才</li>
                  <li>网站地图</li>
                  <li>友情链接</li>
                </ul></td>
              </tr>
          </table></td>
      </tr>
    </table>
  </div>
  <div class="i02">
<DIV class=banner>
  <DIV id=ImageCyclerImage> <A href="#"><IMG alt="IDP Videos" src="img/hero6.jpg"></A> </DIV>
  <DIV id=ImageCyclerOverlay class=grey>
    <DIV id=ImageCyclerOverlayBackground></DIV>
    <P class=title>宁波展会</P>
    <P>讲解人员正在细心的为读者介绍产品<A href="#">Find out more &gt;</A></P>
  </DIV>
  <DIV id=ImageCyclerTabs>
  <div id=mg><A href="#"><img src="img/mg.png" width="95" height="43"></A></div>
  <div id=jnd><A href="#"><img src="img/jnd.png" width="95" height="43"></A></div>
  <div id=yg><A href="#"><img src="img/yg.png" width="95" height="43"></A></div>
  <div id=dg><A href="#"><img src="img/dg.png" width="95" height="43"></A></div>
  <div id=fg><A href="#"><img src="img/hg.png" width="142" height="43"></A></div>
    </DIV>
<div id="Layer1">
  <div id=hg><A href="#"><img src="img/fg.png" width="95" height="43" align="right"></A></div>
  <div id=rb><A href="#"><img src="img/rb.png" width="95" height="47" align="right"></A></div>
  <div id=xjp><A href="#"><img src="img/xjp.png" width="95" height="43" align="right"></A></div>
  <div id=odly><a href="#"><img src="img/odly.png" width="95" height="43" align="right" /></a></div>
  <div id=qt><a href="#"><img src="img/qt.png" width="95" height="43" align="right" /></a></div>
</div>
</DIV>
  </div>
<SCRIPT type=text/javascript src="js/jquery.js"></SCRIPT>
<SCRIPT type=text/javascript>(function($) {
					$.slider = function(opts, data) {
						this.currentSlide = 0;
                        this.opts = opts;
                        this.ddata = data;
                        this.timeout = null;
                        
                        var src = this; 
                        
                        var srcAuto = true;
                        
                        this.initialize = function() {
                            this.attachListeners();
                            this.changeSlide(0);
                        }
                        
                        this.attachListeners = function() {
                            $('#'+this.opts.tabsContainer+' a').each(function(i,n) {
                                var el = $(n);
                                el.css('outline', 'none');
                                // Remove change of tab on click, use as a link instead
                                el.hover(function() {
                                    clearTimeout(src.timeout);
                                    srcAuto = false;
                                    src.currentSlide = i;
                                    src.changeSlide();
                                },function(){
									srcAuto = true;
									src.currentSlide = i;
                                    src.changeSlide();
								});
                            });
                        }
                        
                        this.changeSlide = function() {
                            var slide = src.ddata[src.currentSlide];
                            $('#'+src.opts.tabsContainer+' a').removeClass('active').eq(src.currentSlide).addClass('active');
                            $('#'+src.opts.textContainer+' p:eq(0)').html(slide.title);
                            var moreLink = " <a href='" + slide.overlaylink + "'>Find out more &gt;</a>";

							//if(src.currentSlide == 3){
							//	moreLink = "";
							//}
							
                            $('#'+src.opts.textContainer+' p:eq(1)').html(slide.desc + moreLink);
                            $('#'+src.opts.imageContainer+' img').attr('src', slide.image).attr('alt', slide.title);
                            $('#'+src.opts.imageContainer+' a').attr('href', slide.overlaylink);
                            
							
                            if(srcAuto){
                                src.timeout = setTimeout(src.changeSlide, src.opts.duration*1000);
                            }
                            src.currentSlide = parseInt(src.currentSlide) + 1;
                            if (src.currentSlide >= 5) src.currentSlide = 0; // only 4 items on the homepage
//                            if (src.currentSlide >= src.ddata.length) src.currentSlide = 0;
                        }
                        
                        this.initialize();
                        return this;
                    };
                })(jQuery);
                
                $(function() {
					  $(".favorite").click(function(){
					  showFavorite()
					  return false;
					  })
				  $.slider({ imageContainer: 'ImageCyclerImage', textContainer: 'ImageCyclerOverlay', tabsContainer: 'ImageCyclerTabs', duration: 5 }, 
                            [
                                
                      { image: 'img/hero6.jpg', title: '宁波展会', desc: '讲解人员正在细心的为读者介绍产品', overlaylink : '#'},
                                    
                      { image: 'img/hero10.jpg', title: '长春展会', desc: '明日公司总经理正在为读者介绍产品', overlaylink : '#'},
                                    
                      { image: 'img/hero1.jpg', title: '北京展会', desc: '讲解人员正在细心的为读者介绍产品', overlaylink : '#'},
                                    
                      { image: 'img/hero5.jpg', title: '大连展会', desc: '读者们正在仔细了解产品', overlaylink : '#'},
					  { image: 'img/hero7.jpg', title: '编程词典', desc: '编程全能词典全面上市', overlaylink : '#'}
                                    
                            ]
							);
                });
            </SCRIPT>
  <div class="i03">
    <div class="i03l">
      <div><img src="img/i05.gif" width="268" height="32" /></div>
      <div class="i04"><img src="img/pic1.jpg" width="268" height="109" /></div>
      <div class="i04"> <span class="t14">·编程加油站</span><br />
      课堂讲解、课堂训练、扩展训练和快乐编程驿站
      <br />
        <span class="t14">·编程竞技场</span><br />
        编程大赛、过关斩将、江湖会、编程故事会<br />
<span class="t14">·编程秀</span><br /> 
展示优秀作品
<br />
<span class="t14">·社区论坛</span><br /> 
相互交流平台</div>
      <div class="i04"><a href="gyld.html" class="line"><font color="#FA4303">查看更多&gt;&gt;</font></a></div>
    </div>
    <div class="i03c">
      <div><img src="img/i06.gif" width="268" height="32" /></div>
      <div class="i04">
        <p><a href="#" class="t14a">编程词典个人版及企业版将在宁波图书订</a><a href="#" class="t14a">货会与大家见面</a></p>
      </div>
      <div id="layout">
	<marquee direction="up" scrollamount="3" style="height:307px; ">
		<div id="showInfo"></div>
	</marquee>
	</div>
    </div>
    <div class="i03r">
      <div><img src="img/i07.gif" width="268" height="32" /></div>
      <div class="i04">吉林省明日科技有限公司<br />
        地址：吉林省长春市东盛大街89号 亚泰广场C座<br />
        电话：400-675-1066 0431-84978981 0431-84978982 <br />
        传真：0431-84939777<br />
        企业邮箱：mingrisoft@mingrisoft.com<br />
      </div>
      <div class="i04"><a href="gyld.html" class="line"><font color="#1F82B6">查看更多&gt;&gt;</font></a></div>
      <div class="i04"><img src="img/i09.gif" width="268" height="31" /></div>
      <div class="i05">
        <table width="250" height="20" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td width="22"><a href=javascript:window.external.AddFavorite('http://www.mingrisoft.com','吉林省明日科技有限公司')><img src="img/b_01.gif" width="22" height="20" border="0" onmouseover="this.src='img/c_01.gif'" onmouseout="this.src='img/b_01.gif'"/></a></td>
            <td width="18"></td>
            <td width="20"><a href="#"><img src="img/b_03.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_03.gif'" onmouseout="this.src='img/b_03.gif'" alt="MR编程词典"/></a></td>
            <td width="18"></td>
            <td width="20"><a href="#"><img src="img/b_05.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_05.gif'" onmouseout="this.src='img/b_05.gif'" alt="MR明日图书"/></a></td>
            <td width="18"></td>
            <td width="20"><a href="#"><img src="img/b_07.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_07.gif'" onmouseout="this.src='img/b_07.gif'" alt="MR淘宝店铺"/></a></td>
            <td width="18"></td>
            <td width="20"><a href="#"><img src="img/b_09.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_09.gif'" onmouseout="this.src='img/b_09.gif'" alt="MR新浪网微博"/></a></td>
            <td width="18"></td>
            <td width="20"><a href="#"><img src="img/b_11.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_11.gif'" onmouseout="this.src='img/b_11.gif'" alt="MR人人网"/></a></td>
            <td width="#">&nbsp;</td>
            <td width="20"><a href="#"><img src="img/b_13.gif" width="20" height="20" border="0" onmouseover="this.src='img/c_13.gif'" onmouseout="this.src='img/b_13.gif'" alt="MR腾迅网微博"/></a></td>
          </tr>
        </table>
      </div>
      <div class="i04">
        <table width="209" height="136" border="0" cellpadding="0" cellspacing="0">
          <tr>
            <td align="center" valign="bottom" background="img/i08.gif"><table width="90%" border="0" cellspacing="0" cellpadding="0">
              <tr>
                <td height="24" align="left" valign="top" class="bai">MR全球网络</td>
              </tr>
              <tr><form>
                <td height="35" align="left" valign="top">
                  <select name="select" style="width:190px;" onchange="javascript:window.open(this.options[this.selectedIndex].value)">
                    <option>选择国家</option>
                    <option value="">美国</option>
                    <option value="">加拿大</option>
                    <option value="">英国</option>
                    <option value="">德国</option>
                    <option value="">法国</option>
                    <option value="">日本</option>
                    <option value="">新加坡</option>
                    <option value="">澳大利亚</option>
                    <option value="">新西兰</option>
                    <option value="">意大利</option>
                    <option value="">西班牙</option>
                    <option value="">丹麦</option>
                    <option value="">马来西亚</option>
                    <option value="">荷兰</option>
                    <option value="">爱尔兰</option>
                    <option value="">瑞典</option>
                  </select>                </td>
              </form>
              </tr>
            </table></td>
          </tr>
        </table>
      </div>
    </div>
  </div>
  <div class="i06"><a href="registration.html" target="_blank"><img src="img/bm.jpg" width="868" height="42" border="0" /></a></div>
  <div>
    <table width="100%" border="0" cellspacing="0" cellpadding="0">
      <tr><form>
        <td height="44" align="right" background="img/i10.gif"><input name="textfield" type="text" style="border:solid 1px #CFCECE; width:150px; height:18px ; color:#A3A2A2" value="中文或英文"/>
        <input type="image" name="Submit" value="提交" src="img/i11.gif" width="69" height="22" align="absmiddle"/> <a href="#" target="_blank"><img src="img/ia11.gif" width="69" height="22" border="0" align="absmiddle" /></a></td>
      </form>
      </tr>
    </table>
  </div>
  <div class="i04">
  <TABLE width="868" align=center>
<TR>
<TD><div id=demo style="BACKGROUND: #ffffff; OVERFLOW: hidden; WIDTH: 868px; HEIGHT: 264px">
  <table width="100%" cellpadding="0" cellspacing="0">
    <tr>
      <td id=demo1><table width=100% align=center cellpadding="0" cellspacing="0">
        <tr>
          <td width=160 height=132 align=middle valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td width=160 height=132 align=middle valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td width=160 height=132 align=middle valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td width=160 height=132 align=middle valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td width=160 height=132 align=middle valign="top" style="padding-right:10px"><div class="i07"><a href="#"><img src="img/biao2.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="#">编程词典</a><br />
                <a href="#">MRBCCD</a></div></td>
        </tr>
        <tr>
          <td height="132" align="middle" valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td height="132" align="middle" valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td height="132" align="middle" valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td height="132" align="middle" valign="top" style="padding-right:10px"><div class="i07"><a href="http://www.mrbccd.com" target="_blank"><img src="img/biao.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="http://www.mrbccd.com" target="_blank">编程词典</a><br />
                <a href="http://www.mrbccd.com" target="_blank">MRBCCD</a></div></td>
          <td height="132" align="middle" valign="top" style="padding-right:10px"><div class="i07"><a href="#"><img src="img/biao2.gif" width="160" height="70" border="0" /></a></div>
                <div class="i08"><a href="#">编程词典</a><br />
                <a href="#">MRBCCD</a></div></td>
        </tr>
      </table></td>
      <td id=demo2></td>
    </tr>
  </table>
</div>
<SCRIPT>
var speed=30
demo2.innerHTML=demo1.innerHTML
function Marquee(){
if(demo2.offsetWidth-demo.scrollLeft<=0)
demo.scrollLeft-=demo1.offsetWidth
else{
demo.scrollLeft++
}
}
var MyMar=setInterval(Marquee,speed)
demo.onmouseover=function() {clearInterval(MyMar)}
demo.onmouseout=function() {MyMar=setInterval(Marquee,speed)}
</SCRIPT></TD>
</TR>
</TABLE>
  </div>
    <div style="padding-top:15px">
    <iframe frameborder="0" height="95" id="top" name="top" scrolling="No" src="inc/lxbom.html" width="872"></iframe>
  </div>
</div>
  <SCRIPT type=text/javascript src="js/floatdiv.js"></SCRIPT>
</body>
</html>
