# coding=utf-8
from flask import Flask, render_template, session, redirect, url_for, flash
from flask_bootstrap import Bootstrap
from flask_moment import Moment
import time, os
import sendmail
from config import *
from bookform import *
import savedata

app = Flask(__name__)
app.config['SECRET_KEY'] = 'csrfdog'
bootstrap = Bootstrap(app)
moment = Moment(app)


@app.errorhandler(404)
def page_not_found(e):
    return render_template('404.html'), 404


@app.errorhandler(500)
def internal_server_error(e):
    return render_template('500.html'), 500


@app.route('/', methods=['GET', 'POST'])
def book():
    booker = BookForm()
    # handle
    if booker.validate_on_submit():
        # getdata
        name = booker.name.data
        college = booker.college.data
        stu_num = booker.stu_num.data
        qq = booker.qq.data
        phone = booker.phone.data
        detail = booker.detail.data
        # sql
        savedata.savedata(name,college,stu_num,qq,phone,detail)
        # sendmail
        content = 'name:%s\ncollege:%s\nstu_num:%s\nQQ:%s\nphone:%s\ndetail:%s\n' %(name,college,stu_num,qq,phone,detail)  # mail content
        mail = sendmail.Mail(sender, receivers)
        mail.send(subject, content, sendfile, filepath, sender_name, receivers_name)
        # flash                                             
        flash('上传成功!')
        return redirect(url_for('book'))
    return render_template('book.html', form=booker)

if __name__ == '__main__':
    app.run()
