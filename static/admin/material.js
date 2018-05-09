$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#matTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#matTable').jsGrid('refresh');
    });
    $("#matTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'les_no',title:'课程代码',sorting:false},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'name',title:'标题',width:150},
            {name:'class_time',title:'上课时间',sorting:false},
            {name:'extension',title:'文件格式',sorting:false},
            {name:'updated_at',title:'发布时间',itemTemplate:function (value,item) {
                    return value.substr(0,10)
                }},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改资料');
                    $('#matModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/material/del?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#matTable').jsGrid('loadData');
                                } else {
                                    Messager.alert(data.info);
                                }
                            },
                            error:function () {
                                Messager.alert("请求失败");
                            }
                        })
                    })
                });
                var $download=$('<span>').addClass("glyphicon glyphicon-download-alt").tooltip({title:'下载'}).on("click",function () {
                    $(window.location).prop('href', item.url);
                });
                return $('<div>').addClass("operate").append($edit).append($delete).append($download);
            }}
        ],
        controller: {
            loadData: function (filter) {
                var lessonNo=$('input[name="lessonNoS"]').val();
                var lessonName=$('input[name="lessonNameS"]').val();
                if(lessonNo!=""){filter.lessonNo=lessonNo;}
                if(lessonName!=""){filter.lessonName=lessonName;}
                return $.ajax({
                    type: "GET",
                    url: "/material/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('发布通知');
        $('#matModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#matForm').validate().resetForm();
        $('#matModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#matTable').jsGrid('loadData');
    })
}

function initForm(){
    $.ajax({
        url:'/info/lecture/lectureByTea',
        success:function(data){
            if(data!=null){
                $.each(data,function(i,item){
                    $('#lectureId').append(
                        '<option value='+item.id+'>'+item.les_name+'&nbsp'+item.class_time+'</option>'
                    );
                })
            }
        }
    });
    $("#matForm").validate({
        rules: {
            title: "required"
        },
        messages: {
            title: "请输入名称"
        },
        errorPlacement: function(error, element) {
            error.appendTo(element.parent().parent());
        },
        submitHandler:function(){
            $(".overlay").show();
            $('#matForm').ajaxSubmit({
                url:'/material/save',
                type:'post',
                success:function(data){
                    if(data.status){
                        $('#matForm').validate().resetForm();
                        $(".overlay").hide();
                        $('#matModal').modal('hide');
                        $('#matTable').jsGrid('loadData');
                    }else{
                        $(".overlay").hide();
                        Messager.alert(data.info);
                    }
                },
                error:function(){
                    $(".overlay").hide();
                }
            })
        }
    });
}
function fillForm(item){
    $('#id').val(item.id);
    $('#name').val(item.name);
    $('#type').val(item.type);
    $('#brief').val(item.brief);
    $('#lectureId').val(item.lecture_id)
}
