<?php
if($connID=mysqli_connect("47.111.94.182","root","","stzg")){
    echo " ";
}else{
    echo "连接失败" . mysqli_connect_error();
}
mysqli_query($connID,"set names utf8");//写库
mysqli_query($connID,"set character set 'utf8'");//读库 
?>