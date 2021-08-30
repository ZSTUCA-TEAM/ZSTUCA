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
    return render_template('endterm.html')

if __name__ == '__main__':
    app.run()
