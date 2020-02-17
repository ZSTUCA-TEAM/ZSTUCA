/*
    ① 用setInterval(draw, 33)设定刷新间隔

    ② 用String.fromCharCode(1e2+Math.random()*33)随机生成字母

    ③ 用ctx.fillStyle=’rgba(0,0,0,.05)’; ctx.fillRect(0,0,width,height); ctx.fillStyle=’#0F0′; 反复生成opacity为0.5的半透明黑色背景

    ④ 用x = (index * 10)+10;和yPositions[index] = y + 10;顺序确定显示字母的位置

    ⑤ 用fillText(text, x, y); 在指定位置显示一个字母 以上步骤循环进行，就会产生《黑客帝国》的片头效果。
*/
$(document).ready(function () {
    var s = window.screen;
    var width = q.width = s.width;
    var height = q.height=s.height;
    var yPositions = Array(500).join(0).split('');
    var ctx = q.getContext('2d');
    var draw = function () {
        ctx.fillStyle = 'rgba(0,0,0,.05)';
        ctx.fillRect(0, 0, width, height);
        ctx.fillStyle = 'green';
        ctx.font = '10pt Georgia';
        yPositions.map(function (y, index) {
            text = String.fromCharCode(1e2 + Math.random() * 33);
            x = (index * 10) + 10;
            q.getContext('2d').fillText(text, x, y);
            if (y > Math.random() * 1e4) {
                yPositions[index] = 0;
            } else {
                yPositions[index] = y + 10;
            }
        });
    };
    RunMatrix();
    function RunMatrix() {
        Game_Interval = setInterval(draw, 30);
    }
    $.fn.extend({
        obj:{},
        lu_word:function(obj={}){
            var id = $(this).attr('id');
            this.obj[id] = { text:["Hello World!"],width:"200", height:"300",background:"#000",color:"#fff",speed:'300',sleep:'2000',sign:true,type:true,end:0};
            $.extend(this.obj[id],obj);
            var t = this.obj;
            var that = t[id];
    
            var word = $(this);
            word.css({"width":that.width,"height":that.height,"word-wrap":"break-word","background":that.background});
            word.append("<span class='lu_word_title'></span><span class='lu_word_line' style='width: 2px;background:"+that.color+";height: 20px;border: 1px solid "+that.color+";'></span>");
    
            var title = $(this).find('.lu_word_title');
            var line = $(this).find('.lu_word_line');
            title.css('color',that.color);
    
            var lineSign = true;
            var i = 0;
    
            setInterval(function(){ that = t[id];  go();},that.speed);
            setInterval(cursor,1000);
            function go(){
                if(!that.sign){
                    return;
                }
                if(i >= that.text.length){
                    i=0;
                }
                var text = that.text[i];
                var w = text.substr(0, that.end);
                title.text(w);
                if(that.type) {
                    if(that.end >= text.length){
                        t[id].type = false;
                        that.sign = false;
                        setTimeout(()=>{that.sign = true}, that.sleep);
                    }
                    t[id].end++;
                }else{
                    if(that.end <= 0){
                        t[id].type = true;
                        i++
                    }
                    t[id].end--;
                }
            }
            function cursor(){
                if(lineSign){
                    line.hide();
                    lineSign = false;
                }else{
                    line.show();
                    lineSign = true;
                }
            }
        },
        setColor:function(color){
            var title = $(this).find('.lu_word_title');
            var line = $(this).find('.lu_word_line');
            title.css('color',color);
            line.css('border',"1px solid "+color+"");
        },
        setBg:function(color){
            $(this).css('background',color);
        },
        setSpeed:function(speed){
            that.speed = speed;
        },
        start:function(){
            var id = $(this).attr('id');
            this.obj[id].sign=true;
        },
        stop:function(){
            var id = $(this).attr('id');
            this.obj[id].sign=false;
        },
        write:function(){
            var id = $(this).attr('id');
            this.obj[id].type = true;
        },
        del:function(){
            var id = $(this).attr('id');
            this.obj[id].type=false;
        },
        setWord:function(word){
            var id = $(this).attr('id');
            var that = this.obj[id];
            that.text=[word];
            that.end=0;
            that.sign=true;
            that.type=true;
        }
    });
    //$('#lu_word_box').lu_word();
    var obj1 = {
        text:["A Student Association @ ZheJiang Sci-Tech University"],//文本
        color:"white",//字体和光标颜色
        speed:100,//打字速度
        sleep:3000,//回退停顿时间
        width:"400",//宽度
        height:"300",//高度
        background:"#000",//背景颜色
        sign:true//启动或停止
    };
    var box = $('#lu_word_box');
    box.lu_word(obj1);
});