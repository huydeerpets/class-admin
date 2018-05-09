{{template "public/header.tpl" .}}
<link rel="stylesheet" href="/static/css/AdminLTE.min.css">
<script src="/static/plugins/ckeditor/ckeditor.js"></script>
<script src="/static/js/jquery.form.min.js"></script>
<script src="/static/js/messager.js"></script>


<body class="hold-transition skin-blue layout-top-nav">
<div class="wrapper">

    <header class="main-header">
        <nav class="navbar navbar-static-top">
            <div class="container">
                <div class="navbar-header">
                    <div class="logo">
                        <span class="logo-mini"><b>课堂互动</b></span>
                        <span class="logo-lg"><b>课堂互动</b>平台后台管理系统</span>
                    </div>
                </div>
            </div>
        </nav>

    </header>
    <!-- Full Width Column -->
    <div class="content-wrapper">
        <div class="container">
            <!-- Content Header (Page header) -->
            <section class="content-header">
                <h1>
                    解答问题
                </h1>
                <ol class="breadcrumb">
                    <li><a href="#"><i class="fa fa-dashboard"></i> 首页</a></li>
                    <li><a href="#">课程问题</a></li>
                </ol>
            </section>

            <section class="content">
                <div class="box box-primary">
                    <form class="form-horizontal" id="resultForm">
                        <input type="hidden" name="id" id="id" value="{{.ques.id}}">
                        <div class="box-body">
                            <div class="form-group col-md-4">
                                <label class="col-sm-4 control-label">课程代码:</label>
                                <div class="col-sm-8">
                                    {{.ques.les_no}}
                                </div>
                            </div>
                            <div class="form-group col-md-4">
                                <label for="typeId" class="col-sm-4 control-label">课程名称:</label>
                                <div class="col-sm-8">
                                   {{.ques.les_name}}
                                </div>
                            </div>
                            <div class="form-group col-md-4">
                                <label for="date" class="col-sm-4 control-label">上课时间:</label>
                                <div class="col-sm-8">
                                    {{.ques.class_time}}
                                </div>
                            </div>
                            <div class="form-group col-md-4">
                                <label for="source" class="col-sm-4 control-label">提问学生:</label>
                                <div class="col-sm-8">
                                    {{.ques.stu_name}}
                                </div>
                            </div>
                            <div class="form-group col-md-4">
                                <label for="title" class="col-sm-4 control-label">发布时间:</label>
                                <div class="col-sm-8">
                                    {{.ques.created_at}}
                                </div>
                            </div>

                            <div class="col-md-4">
                                <button id="saveBtn" type="button" class="btn btn-primary" style="float:right;">保存</button>
                            </div>
                            <div class="form-group col-md-12">
                                <label for="url" class="col-sm-2 control-label">问题详情:</label>
                                <textarea rows="4" class="col-sm-8" rows="6" disabled>
                                    {{.ques.question}}
                                </textarea>
                            </div>
                            <div class="form-group col-md-12">
                                <textarea id="answer" name="answer" rows="10" cols="80">
                                </textarea>
                            </div>
                        </div>
                    </form>
                    <div class="overlay" style="display: none;">
                        <i class="fa fa-refresh fa-spin"></i>
                    </div>
                </div>
            </section>
        </div>
    </div>
</div>
<script type="text/javascript">
    var editor;
    $(function(){
        editor=CKEDITOR.replace( 'answer' );
        editor.setData({{.ques.answer}});
        init();
    });
    function init(){

        $('#saveBtn').on("click",function(){
            $('#answer').val(editor.getData());
            $(".overlay").show();
            $('#resultForm').ajaxSubmit({
                url:'/question/save',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#id').val(data.msg);
                    }else{
                        $(".overlay").hide();
                        Messager.alert(data.msg);
                    }
                },
                error:function(){
                    Messager.alert("保存失败");
                    $(".overlay").hide();
                }
            })
        })

    }
</script>
{{template "public/footer.tpl" .}}
