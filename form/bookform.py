# coding=utf-8
from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField, SelectField
from wtforms.validators import DataRequired


class BookForm(FlaskForm):
    name = StringField('姓名', validators=[DataRequired()])
    college = StringField('学院/部门', validators=[DataRequired()])
    stu_num = StringField('学号/工号', validators=[DataRequired()])
    qq = StringField('qq/微信', validators=[DataRequired()])
    phone = StringField('手机', validators=[DataRequired()])
    detail = StringField('简述问题', validators=[DataRequired()])
    submit = SubmitField("预约")

# defination
name = None
college = None
stu_num = None
qq = None
phone = None
detail = None
