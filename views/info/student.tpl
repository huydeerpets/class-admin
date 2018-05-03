{{template "public/header.tpl" .}}
        <style>
            input.error, select.error {
                border: 1px solid #ff9999;
                background: #ffeeee;
            }
            label.error {
                font-size: .8em;
                color: #ff6666;
                display:inline;
            }
            .note {
                font-size: .8em;
                color: #a5a5a5;
                display:inline;
                float:left;
            }
            .operate span{
                margin-left:10px;
            }
            .operate span:hover{
                color: #ff7167;
            }
            .newline {
                word-break:break-all;
            }
        </style>
<link rel="stylesheet" href="/static/css/jsgrid.min.css">
<link rel="stylesheet" href="/static/css/jsgrid-theme.min.css">
<script src="/static/js/jsgrid.min.js"></script>
<script src="/static/js/jquery.validate.min.js"></script>
<script src="/static/js/jquery.form.min.js"></script>
<script src="/static/js/messager.js"></script>
<script src="/static/admin/common.js"></script>
<script src="/static/admin/student.js"></script>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
{{template "public/menu.tpl" .}}
    <script type="text/javascript">
        activeDiv=function () {
            $('#info-li').addClass("active");
            $('#stu-li').addClass("active");
        };
    </script>
    <div class="content-wrapper">
        <section class="content-header">
            <h1>学生信息</h1>
            <ol class="breadcrumb">
                <li><a href="/public/index"><i class="fa fa-dashboard"></i> 首页</a></li>
                <li>基本信息</li>
                <li class="active">学生信息</li>
            </ol>
        </section>
        <section class="content">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-body">
                            <div class="form-inline" style="height:28px">
                                <div style="float:left">
                                    <button id="addBtn" class="btn btn-primary btn"><span class="glyphicon glyphicon-plus"></span>&nbsp新建学生信息</button>
                                </div>
                            </div>
                            <div class="form-inline" style="margin:12px 0">
                                <select class="form-control input-sm" id="pageSize">
                                    <option selected="selected">10</option>
                                    <option>25</option>
                                    <option>50</option>
                                    <option>100</option>
                                </select>
                                <div style="float:right;display:inline">
                                    <input type="text" class="form-control input-sm" name="stuIdS" placeholder="学号">
                                    <input type="text" class="form-control input-sm" name="nameS" placeholder="姓名">
                                    <button id="searchBtn" class="btn btn-primary btn-sm">搜索</button>
                                </div>
                            </div>
                            <div id="stuTable" style="margin:10px 0"></div>
                        </div>
                        <div class="overlay" style="display: none;">
                            <i class="fa fa-refresh fa-spin"></i>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
    <div id="stuModal" class="modal fade" data-backdrop="static">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="box">
                    <div class="modal-header">
                        <button type="button" class="close closeBtn" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
                        <h4 class="modal-title" id="modalLabel"></h4>
                    </div>
                    <form class="form-horizontal" id="stuForm">
                    <div class="modal-body" style="text-align:center;">
                            <input type="hidden" name="id" id="id">
                            <div class="box-body">
                                <div class="form-group">
                                    <label for="name" class="col-sm-3 control-label">姓名</label>
                                    <div class="col-sm-6">
                                        <input type="text" id="name" name="name" class="form-control">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="stuId" class="col-sm-3 control-label">学号</label>
                                    <div class="col-sm-6">
                                        <input type="text" id="stuId" name="stuId" class="form-control">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="gender" class="col-sm-3 control-label">性别</label>
                                    <div class="col-sm-6">
                                        <select id="gender" name="gender" class="form-control">
                                            <option value="0">——请选择——</option>
                                            <option value="1">男</option>
                                            <option value="2">女</option>
                                        </select>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="email" class="col-sm-3 control-label">邮箱</label>
                                    <div class="col-sm-8">
                                        <input id="email" name="email" class="form-control">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="tel" class="col-sm-3 control-label">联系方式</label>
                                    <div class="col-sm-8">
                                        <input id="tel" name="tel" class="form-control">
                                    </div>
                                </div>
                            </div>
                    </div>
                    <div class="modal-footer">
                        <button type="submit" class="btn btn-primary" id="submitBtn">确定</button>
                        <button type="button" class="btn btn-default closeBtn">取消</button>
                    </div>
                    </form>
                    <div class="overlay" style="display: none;">
                        <i class="fa fa-refresh fa-spin"></i>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{template "public/footer.tpl" .}}
