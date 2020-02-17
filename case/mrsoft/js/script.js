// JavaScript Document
$(document).ready(function(){
	$("#header .nav>li").each(function(i){
		$(this).addClass("m0"+(i+1));
	});
	
	$("#header .nav li").mouseover(function(){
		hidMenu();
		$(this).addClass("showMenu");
		$(this).addClass("current");
	});
	
	$(".logoList01 li").each(function(i){
		if((i+1)%4 == 0) {
			$(this).addClass("li");
		}
	});


	
	//最热活动效果
	$(".pagesZrhd .newsList li").mouseover(function(){
		$(this).addClass("showInfo")
	}).mouseout(function(){
		$(this).removeClass("showInfo")
	});
	$(".faqList dd").mouseover(function(){
		$(this).addClass("showInfo")
	}).mouseout(function(){
		$(this).removeClass("showInfo")
	});
	
	
	
	$(".lmksList a").mousemove(function(){
		$(this).addClass("aAlpha70")
	}).mouseout(function(){
		$(this).removeClass("aAlpha70")
	})
	
})

function hidMenu(){
	$("#header .nav li").each(function(i){
		$(this).removeClass("showMenu");
		$(this).removeClass("current");
	});
}