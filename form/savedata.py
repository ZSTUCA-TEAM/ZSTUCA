import pymysql
from config import *
import datetime
def savedata(name,college,stu_num,qq,phone,detail):
    nowtime = str(datetime.datetime.now())[:10]
    conn = pymysql.connect(host=host, port= 3306, user=user, password=password, db=db)
    sql = "INSERT INTO `zstuca`.`book` (`姓名`, `学院`, `学号`, `qq`, `手机`, `问题简述`, `时间`) VALUES ('{0}','{1}','{2}','{3}','{4}','{5}','{6}')".format(name,college,stu_num,qq,phone,detail,nowtime)
    try:
        with conn.cursor() as cursor:
            cursor.execute(sql)
        conn.commit()
    except Exception as e:
        conn.rollback()
    conn.close()
