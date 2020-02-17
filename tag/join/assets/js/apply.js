
$(function(){
        var select1=document.getElementById('select');
        var select2=document.getElementById('selectt');
        select1.onchange=function() {
            var index=this.selectedIndex;
            if(index==0){
                select2.selectedIndex = 0;
                for( var i = 0; i <=4; i++ ){
                    select2.options[i].disabled=true;
                }
            }else{
                for( var i = 0; i <=4; i++ ){
                    select2.options[i].disabled=false;
                }
                select2.options[index].disabled=true;
            }
            $('select').formSelect();
        }
        select2.onchange=function() {
            var index=this.selectedIndex;
            if(select1.selectedIndex==0){
                alert("请选择第一志愿");
                return;
            }
            for( var i = 1; i <=4; i++ ){
                select1.options[i].disabled=false;
            }
            if(index!=0){
                select1.options[index].disabled=true;
            }
            $('select').formSelect();
        }
        $(document).ready(function () {
            $('select').formSelect();
        });
        //性别选择
        var $female = $(".genderbox.female"),
        $male = $(".genderbox.male"),
        $submit = $("#submit");
        console.log($submit);
        $female.on("click", function () {
            console.log(123);
            var $on = $(this).find(".icon.on");
            var $off = $(this).find(".icon.off");
            $on.removeClass("hide");
            $off.addClass("hide");
            $male.find(".icon.on").addClass("hide");
            $male.find(".icon.off").removeClass("hide");
        })
        $male.on("click", function () {
            var $on = $(this).find(".icon.on");
            var $off = $(this).find(".icon.off");
            $on.removeClass("hide");
            $off.addClass("hide");
            $female.find(".icon.on").addClass("hide");
            $female.find(".icon.off").removeClass("hide");
        })
        
})

 