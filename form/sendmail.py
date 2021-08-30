# coding:utf-8
import smtplib
from email.mime.base import MIMEBase
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.header import Header
from email import encoders
from config import *


class Mail:
    def __init__(self, sender, receivers):
        self.mail_host = mail_host       # 设置服务器:这个是qq邮箱服务器
        self.mail_pass = mail_pass           # 授权码  需到服务提供商处获取
        self.sender = sender      # 你的邮箱地址
        self.receivers = receivers  # 收件人的邮箱地址，可设置为你的QQ邮箱或者其他邮箱，可多个

    def send(self, subject, content, sendfile, filepath, sender_name, receivers_name):
        msg = MIMEMultipart()                              # multipart类型（附件）
        msg['From'] = Header(sender_name, 'utf-8')
        msg['To'] = Header(receivers_name[0], 'utf-8')     # 收件人名字，但只能传字符串，因此取列表第一个
        msg['Subject'] = Header(subject, 'utf-8')

        # msg.attach(MIMEText(content, 'html', 'utf-8'))    # html类型
        msg.attach(MIMEText(content, 'plain', 'utf-8'))

        # 添加附件
        if sendfile:
            with open(filepath, 'rb') as f:
                # 设置附件的MIME和文件名，这里是png类型:
                mime = MIMEBase('application', 'octet-stream', filename=filepath.split('/')[-1])
                # 加上必要的头信息:
                mime.add_header('Content-Disposition', 'attachment', filename=filepath.split('/')[-1])
                mime.add_header('Content-ID', '<0>')
                mime.add_header('X-Attachment-Id', '0')
                mime.set_payload(f.read())      # 把附件的内容读进来
                encoders.encode_base64(mime)    # 用Base64编码
                msg.attach(mime)                # 添加到MIMEMultipart
        try:
            smtpObj = smtplib.SMTP_SSL(self.mail_host, 465)
            smtpObj.login(self.sender, self.mail_pass)
            smtpObj.sendmail(self.sender, self.receivers, msg.as_string())
            smtpObj.quit()
            print('邮件发送成功')
        except smtplib.SMTPException as e:
            print('邮件发送失败')


if __name__ == '__main__':
    sender = 'root@jojo-m.cn'           # 发件人邮箱
    receivers = ['root@jojo-m.cn']      # 收件人列表
    subject = 'test subject'            # 发送的主题
    content = 'This is test content'    # 发送的内容
    sendfile = 0                        # 是否发送文件
    filepath = ''          # 文件路径
    sender_name = "ruokeqx"             # 发件人姓名
    receivers_name = [':)']             # 收件人姓名

    mail = Mail(sender, receivers)
    mail.send(subject, content, sendfile, filepath, sender_name, receivers_name)
