<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>2019干事申请表</title>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300,400">
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <link rel="stylesheet" href="/assets/css/materialize.min.css">
    <link rel="stylesheet" href="/assets/css/tooplate.css">
    <link rel="stylesheet" href="./assets/css/apply.css">
</head>
<%
session("apply")="doing"
%>
<body id="application">
    <div class="container">
        <div class="row">
            <div class="col-xl-12 col-lg-12 col-md-12 col-sm-12  mx-auto">
                <header class="mt-5 mb-5 text-center">
                    <h3>2019会员申请表</h3>
                    <p class="tm-form-description">浙理计协——期待你的加入！</p>
                </header>
                <form id="form" name="form" action="check.asp" method="post" enctype="application/x-www-form-urlencoded" class="tm-form-white tm-font-big" accept-charset="UTF-8"  _charset="UTF-8">
                    <div class="tm-bg-white tm-form-pad-big">
                        <div class="form-group mb-5">
                            <label for="name" class="black-text mb-4 big">姓名</label>
                            <input id="name" name="name" type="text" value="" class="validate" required>
                        </div>
                        <div class="form-group mb-5">
                            <label for="id" class="black-text mb-4 big">学号</label>
                            <input id="id" name="id" type="text" value="" class="validate" maxlength="13" required>
                        </div>
                        <div class="form-group mb-5">
                            <label for="tel" class="black-text mb-4 big">电话</label>
                            <input id="tel" name="tel" type="tel" value="" class="validate" maxlength="11" required>
                        </div>
                        <div class="form-group mb-5">
                            <label for="qq" class="black-text mb-4 big">QQ</label>
                            <input id="qq" name="qq" type="text" value="" class="validate" maxlength="10"  required>
                        </div>
                        <div class="form-group mb-5" required>
                            <label for="college" class="black-text mb-4 big">学院</label>
                            <input id="college" name="college" value="" type="text" class="validate" required>
                        </div>
                        <div class="form-group mb-5">
                            <label for="major" class="black-text mb-4 big">专业</label>
                            <input id="major" name="major" value="" type="text" class="validate" required>
                        </div>
                    </div>
                    <div class="text-center mt-5">
                        <button id="submit" type="submit" class="waves-effect btn-large btn-large-white">Submit</button>
                    </div>
                </form>
            </div>
        </div>
        <a id="ahover" href="https://www.zstuca.club/">
            <div class="web_circle">
                <img src="./assets/img/o_my_blog_circle.png" class="circle">
                <span>浙 理<br>计 协</span>
            </div>
        </a>
        <footer class="row tm-mt-big mb-3">
            <div class="col-xl-12 text-center">
                <p class="d-inline-block tm-bg-black white-text py-2 tm-px-5">
                    Copyright &copy; 2019 ZSTUCA
                    
                    - <a rel="nofollow" href="http://www.tooplate.com/view/2105-input">Input</a> by 
                    <a rel="nofollow" href="http://tooplate.com" class="tm-footer-link">Tooplate</a>
                </p>
            </div>
        </footer>
    </div>
    <script src="/assets/js/jquery-1.11.1.min.js"></script>
    <script src="/public/js/materialize.min.js"></script>
    <script src="/public/js/jquery.serializeObject2.0.3.min.js"></script>
    <script src="/public/js/jquery.lazyload.js"></script>
    <script src="/public/js/effect.js"></script>
</body>
</html>