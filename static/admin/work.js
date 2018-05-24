$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#workTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#workTable').jsGrid('refresh');
    });
    $("#workTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'title',title:'标题',width:60},
            {name:'content',title:'内容',sorting:false},
            {name:'les_no',title:'课程代码',sorting:false},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'class_time',title:'上课时间',sorting:false},
            {name:'updated_at',title:'发布时间',itemTemplate:function (value,item) {
                    return value.substr(0,10)
                }},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改作业');
                    $('#workModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/work/del?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#workTable').jsGrid('loadData');
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
                return $('<div>').addClass("operate").append($edit).append($delete);
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
                    url: "/work/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('布置作业');
        $('#workModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#workForm').validate().resetForm();
        $('#workModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#workTable').jsGrid('loadData');
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
    $("#workForm").validate({
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
            $('#workForm').ajaxSubmit({
                url:'/work/save',
                type:'post',
                success:function(data){
                    if(data.status){
                        $('#workForm').validate().resetForm();
                        $(".overlay").hide();
                        $('#workModal').modal('hide');
                        $('#workTable').jsGrid('loadData');
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
    $('#title').val(item.title);
    $('#content').val(item.content);
}
