
{{template "public/header.tpl" .}}
<script src="/static/js/jquery.form.min.js"></script>
<script src="/static/js/messager.js"></script>
<script type="text/javascript">
    function fromsubmit() {
        // if($('select[name="userType"] option:selected').val()=="0"){
        //     Messager.alert("请选择用户类型")
        //     return
        // }
        if($('input[name="username"]').val()==""){
            Messager.alert("请输入用户名")
            return
        }
        if($('input[name="password"]').val()==""){
            Messager.alert("请输入密码")
            return
        }
        $("#loginForm").ajaxSubmit({
            url: '/public/login?isajax=1',
            type: 'post',
            success: function (r) {
                if (r.status) {
                    location.href = "/public/index"
                } else {
                    Messager.alert(r.info);
                }
            }
        });
    }
    //这个就是键盘触发的函数
    var SubmitOrHidden = function (evt) {
        evt = window.event || evt;
        if (evt.keyCode == 13) {//如果取到的键值是回车
            fromsubmit();
        }
    };
    window.document.onkeydown = SubmitOrHidden;//当有键按下时执行函数
</script>
<body class="hold-transition login-page">
    <div class="login-box">
        <div class="login-logo">
            <a href="#"><b>课堂互动平台</b>后台管理</a>
        </div>
        <div class="login-box-body">
            <p class="login-box-msg">Sign in to start your session</p>
            <form id="loginForm" >
                {{/*<div class="form-group has-feedback">*/}}
                    {{/*<select class="form-control" name="userType">*/}}
                        {{/*<option value="0">——用户类型——</option>*/}}
                        {{/*<option value="1">管理员</option>*/}}
                        {{/*<option value="2">教师</option>*/}}
                    {{/*</select>*/}}
                {{/*</div>*/}}
                <div class="form-group has-feedback">
                    <input type="text" class="form-control" name="username" placeholder="账号">
                    <span class="glyphicon glyphicon-user form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback">
                    <input type="password" class="form-control" name="password" placeholder="密码">
                    <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                </div>
                <div class="row">
                    <div class="col-xs-8">
                    </div>
                    <div class="col-xs-4">
                        <button type="button" class="btn btn-primary btn-block btn-flat" onclick="fromsubmit()">Sign In</button>
                    </div>
                </div>
            </form>
        </div>
    </div>

{{template "public/footer.tpl" .}}
