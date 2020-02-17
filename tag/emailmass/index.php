<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>邮件群发</title>
	<!--
    Template 2105 Input
	http://www.tooplate.com/view/2105-input
	-->
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,400">
    <link rel="stylesheet" href="css/bootstrap.min.css">
    <link rel="stylesheet" href="css/materialize.min.css">
    <link rel="stylesheet" href="css/tooplate.css">
</head>

<body id="survey" class="font-weight-light">
<?php
include $_SERVER["DOCUMENT_ROOT"].'/public/php/mail.php';
			
// send_email(["1600337300@qq.com",'1308294811@qq.com','1367138194@qq.com','571080761@qq.com'],"维修队预约通知",'<html>
// 	<body>
// 		<h1>维修预约通知</h1>
// 		<table  width="100%" border="0">
// 			<tr><td width="15%">预约人： </td><td width="15%">'. $xm . '</td><td>预约时间：'. $time . ' </td></tr>
// 			<tr><td>基本信息： </td>
// 				<td colspan="3">'. $xm ."|".$dh."|".$qq."|" . $address . ' </td>
// 			</tr>
// 			<tr><td>问题描述： </td>
// 				<td colspan="3">'. $text . '</td>
// 			</tr>
// 		</table>
// 		<a href="http://47.95.142.33/tag/reservationphp/note.php" target="_blank" rel="noopener">查看详情信息</a>
// 		<div><strong>[维修部信息处]</strong></div>
// 	</body>
// </html>');


send_email("1600337300@qq.com","维修队预约通知",'<html>
    <body>
        <h1>维修预约通知</h1>
        <table  width="100%" border="0">
            <tr><td width="15%">预约人： </td><td width="15%"></td><td>预约时间： </td></tr>
            <tr><td>基本信息： </td>
                <td colspan="3"> </td>
            </tr>
            <tr><td>问题描述： </td>
                <td colspan="3"></td>
            </tr>
        </table>
        <a href="http://47.95.142.33/tag/reservationphp/note.php" target="_blank" rel="noopener">查看详情信息</a>
        <div><strong>[维修部信息处]</strong></div>
    </body>
</html>');
?>
    <div class="container tm-container-max-800">
        <div class="row">
            <div class="col-12">
                <header class="mt-5 mb-5 text-center">
                    <h1 class="font-weight-light tm-form-title">会员邮件群发系统</h1>
                    <p class="tm-form-description">浙江理工大学研发部</p>
                </header>
                <form action="" method="post">
                        <div class="row">
                            <div class="col-xl-12">
                                <textarea class="p-3" name="message" id="message" cols="30" rows="3" placeholder="Additional Message...(optional)"></textarea>
                            </div>
                        </div>
                    <div class="text-center mt-5">
                        <button type="submit" class="waves-effect btn-large">Submit</button>
                    </div>
                </form>
            </div>
        </div>
        <footer class="row tm-mt-big mb-3">
            <div class="col-xl-12">
                <p class="text-center grey-text text-lighten-2 tm-footer-text-small">
                    Copyright &copy; 2018 ZSTUCA
                    
                    - <a rel="nofollow" href="http://www.tooplate.com/view/2105-input">Input</a> by 
                    <a rel="nofollow" href="http://tooplate.com" class="tm-footer-link">Tooplate</a>
                </p>
            </div>
        </footer>
    </div>

    <script src="js/jquery-3.2.1.slim.min.js"></script>
    <script src="js/materialize.min.js"></script>
    <script>
        $(document).ready(function () {
            $('select').formSelect();
        });
    </script>
</body>

</html>