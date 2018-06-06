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
<script src="/static/admin/notice.js"></script>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
{{template "public/menu.tpl" .}}
    <script type="text/javascript">
        activeDiv=function () {
            $('#notice-li').addClass("active");
        };
    </script>
    <div class="content-wrapper">
        <section class="content-header">
            <h1>通知</h1>
        </section>
        <section class="content">
            <div class="row">
                <div class="col-xs-12">
                    <div class="box">
                        <div class="box-body">
                            <div class="form-inline" style="height:28px">
                                <div style="float:left">
                                    <button id="addBtn" class="btn btn-primary btn"><span class="glyphicon glyphicon-plus"></span>&nbsp新建通知</button>
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
                                    <input type="text" class="form-control input-sm" name="lessonNoS" placeholder="课程代码">
                                    <input type="text" class="form-control input-sm" name="lessonNameS" placeholder="课程名称">
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
                            <div id="noticeTable" style="margin:10px 0"></div>
                        </div>
                        <div class="overlay" style="display: none;">
                            <i class="fa fa-refresh fa-spin"></i>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
    <div id="noticeModal" class="modal fade" data-backdrop="static">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="box">
                    <div class="modal-header">
                        <button type="button" class="close closeBtn" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">Close</span></button>
                        <h4 class="modal-title" id="modalLabel"></h4>
                    </div>
                    <form class="form-horizontal" id="noticeForm">
                    <div class="modal-body" style="text-align:center;">
                            <input type="text" style="display:none;" name="id" id="id">
                            <div class="box-body">
                                <div class="form-group">
                                    <label for="lectureId" class="col-sm-3 control-label">课程</label>
                                    <div class="col-sm-6">
                                        <select id="lectureId" name="lectureId" class="form-control"></select>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="title" class="col-sm-3 control-label">标题</label>
                                    <div class="col-sm-6">
                                        <input type="text" id="title" name="title" class="form-control">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="content" class="col-sm-3 control-label">详细内容</label>
                                    <div class="col-sm-6">
                                        <textarea id="content" name="content" class="form-control" rows="4"></textarea>
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
