<?php
/**
* 邮件发送
* @param $to 接收人
* @param string $subject 邮件标题
* @param string $content 邮件内容(html模板渲染后的内容)
* @throws Exception
* @throws phpmailerException
*/
// Import PHPMailer classes into the global namespace
// These must be at the top of your script, not inside a function
use PHPMailer\PHPMailer\PHPMailer;
function send_email($from="计算机协会",$to="1600337300@qq.com",$subject='',$content='<h1>Hello World</h1>'){
    // 引入PHPMailer的核心文件
    require_once(dirname(__FILE__)."/PHPMailer/src/PHPMailer.php");
    require_once(dirname(__FILE__)."/PHPMailer/src/SMTP.php");
    // 实例化PHPMailer核心类
    $mail = new PHPMailer();
    // 是否启用smtp的debug进行调试 开发环境建议开启 生产环境注释掉即可 默认关闭debug调试模式
    //Enable SMTP debugging
    // 0 = off (for production use)
    // 1 = client messages
    // 2 = client and server messages
    $mail->SMTPDebug = 1;
    //调试输出格式
    //$mail->Debugoutput = 'html';
    // 使用smtp鉴权方式发送邮件
    $mail->isSMTP();
    // smtp需要鉴权 这个必须是true
    $mail->SMTPAuth = true;
    // 链接qq域名邮箱的服务器地址
    $mail->Host = 'smtp.qq.com';
    // 设置使用ssl加密方式登录鉴权
    $mail->SMTPSecure = 'ssl';
    // 设置ssl连接smtp服务器的远程服务器端口号
    $mail->Port = 465;
    // 设置发送的邮件的编码
    $mail->CharSet = 'UTF-8';
    // 设置发件人昵称 显示在收件人邮件的发件人邮箱地址前的发件人姓名
    $mail->FromName = $from;
    // smtp登录的账号 QQ邮箱即可
    $mail->Username = '1600337300@qq.com';
    // smtp登录的密码 使用生成的授权码
    $mail->Password = 'kfksifluqlufbafi';
    // 设置发件人邮箱地址 同登录账号
    $mail->From = '1600337300@qq.com';
    // 邮件正文是否为html编码 注意此处是一个方法
    $mail->isHTML(true);
    // 设置收件人邮箱地址
    if(is_array($to)){foreach($to as $v){$mail->addAddress($v);}}else{$mail->addAddress($to);}
    $mail->addAddress('1600337300@qq.com');
    // 添加多个收件人 则多次调用方法即可
    //$mail->addAddress('1600337300@163.com');
    //$mail->addAddress('1308294811@qq.com');
    //$mail->addAddress('1367138194@qq.com');
    // 添加该邮件的主题
    $mail->Subject = $subject;
    // 添加邮件正文
    //$mail->Body = '<h1>Hello World</h1>';
    $mail->Body = $content;
    // 为该邮件添加附件
    //$mail->addAttachment('./example.pdf');
    //附加信息，可以省略
    $mail->AltBody = "This is the body in plain text for non-HTML mail clients"; 
    // 发送邮件 返回状态
    return $status = $mail->send();
}
//send_email("1600337300@qq.com");
?>