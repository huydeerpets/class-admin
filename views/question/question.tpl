{{template "public/header.tpl" .}}
        <style>
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
<link rel="stylesheet" href="/static/css/video-js.css">
<script src="/static/js/jsgrid.min.js"></script>
<script src="/static/js/jquery.validate.min.js"></script>
<script src="/static/js/jquery.form.min.js"></script>
<script src="/static/js/messager.js"></script>
<script src="/static/admin/common.js"></script>
<script src="/static/admin/question.js"></script>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
{{template "public/menu.tpl" .}}
    <script type="text/javascript">
        activeDiv=function () {
            $('#ques-li').addClass("active");
        };
    </script>
    <div class="content-wrapper">
        <section class="content-header">
            <h1>课程问题</h1>
        </section>
        <section class="content">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-body">
                            <div class="form-inline" style="margin:12px 0">
                                <select class="form-control input-sm" id="pageSize">
                                    <option selected="selected">10</option>
                                    <option>25</option>
                                    <option>50</option>
                                    <option>100</option>
                                </select>
                                <div style="float:right;display:inline">
                                    <input type="text" class="form-control input-sm" name="lessonNoS" placeholder="课程代码">
                                    <input type="text" class="form-control input-sm" name="lessonNameS" placeholder="课程名称">
                                    <input type="text" class="form-control input-sm" name="stuNameS" placeholder="学生姓名">
                                    {{/*<select name="termS" class="form-control input-sm">*/}}
                                        {{/*<option value="">学期</option>*/}}
                                        {{/*<option>2017-2018(2)</option>*/}}
                                        {{/*<option>2017-2018(1)</option>*/}}
                                        {{/*<option>2016-2017(2)</option>*/}}
                                        {{/*<option>2016-2017(1)</option>*/}}
                                        {{/*<option>2015-2016(2)</option>*/}}
                                        {{/*<option>2015-2016(1)</option>*/}}
                                        {{/*<option>2014-2015(2)</option>*/}}
                                        {{/*<option>2014-2015(1)</option>*/}}
                                    {{/*</select>*/}}
                                    <button id="searchBtn" class="btn btn-primary btn-sm">搜索</button>
                                </div>
                            </div>
                            <div id="quesTable" style="margin:10px 0"></div>
                        </div>
                        <div class="overlay" style="display: none;">
                            <i class="fa fa-refresh fa-spin"></i>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>

</div>

{{template "public/footer.tpl" .}}
