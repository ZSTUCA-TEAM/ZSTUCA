<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    

<?php 
    session_start();
    if(extension_loaded('gd')){
        echo '可以使用gd<br>';
        foreach(gd_info() as $cate=>$value){
            echo "$cate:$value<br>";
        }
    }
    else{
        echo '没有安装gd扩展';
    }
    
    echo '这是#2222';
    $i='hell0';
    echo $i;
    echo "<p>";
    echo "$i";
    echo '$i';
    echo "<p>";
    echo "222".$i;
    echo "<p>";
    echo __FILE__."<p>".__LINE__."<p>".PHP_OS."<p>".PHP_VERSION;
    echo "<p>";
    echo $_SERVER['REMOTE_ADDR'];
    echo "<p>";
    echo $_SERVER['REMOTE_HOST'];
    echo "<p>";
    echo $_SERVER['REQUEST_METHOD'];
    echo "<p>";
    echo $_SERVER['REMOTE_PORT'];
    echo "<p>";
    echo $_SERVER['SERVER_PORT'];
    echo "<p>";
    $num=1;
    for($j=1;$j<=100;$j++){
        $num+=$j;
    }
    echo $num;
    echo "<p>";
    echo rand(5,646464);
    echo "<p>";
    if(!isset($_COOKIE["vtime"])){
        setcookie("vtime", date("y-m-d H-i-s"));
    }else{
        setcookie("vtime", date("y-m-d H-i-s"),time()+60);
        echo "<p>";
        echo "上一次".$_COOKIE["vtime"];
    }
    echo "<p>";
    echo date("y-m-d H-i-s");
    echo "<p>";
    $str=crypt($i,CRYPT_STD_DES);
    echo $str;
    echo "<p>";
    echo md5("hellp world");
    echo "<p>";
    echo sha1("hellp world");
    echo "<p>";
    if($connID=mysqli_connect("47.111.94.182","root","","stzg")){
        echo "连接成功";
    }else{
        echo "连接失败";
    }
    mysqli_query($connID,"set names utf8");//写库
    mysqli_query($connID,"set character set 'utf8'");//读库 
    $rs=mysqli_query($connID,"select * from wxyy");
    while($row=mysqli_fetch_assoc($rs)){
        echo "<p>";
        echo var_dump($row);
    }
    phpinfo();
?>
</body>
</html>