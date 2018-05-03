$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#noticeTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#noticeTable').jsGrid('refresh');
    });
    $("#noticeTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'title',title:'标题',width:60},
            {name:'content',title:'内容',sorting:false},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改通知');
                    $('#noticeModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/notice/del?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#noticeTable').jsGrid('loadData');
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
                if(lessonNo!=""){filter.lessonNo=lessonNo;}
                return $.ajax({
                    type: "GET",
                    url: "/notice/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('发布通知');
        $('#noticeModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#noticeForm').validate().resetForm();
        $('#noticeModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#noticeTable').jsGrid('loadData');
    })
}

function initForm(){
    $.ajax({
        url:'/info/lecture/lectureByTea',
        success:function(data){
            if(data!=null){
                $.each(data,function(i,item){
                    $('#lectureId').append(
                        '<option value='+item.ID+'>'+item.les_name+'&nbsp'+item.class_time+'</option>'
                    );
                })
            }
        }
    });
    $("#noticeForm").validate({
        rules: {
            name: "required"
        },
        messages: {
            name: "请输入名称"
        },
        errorPlacement: function(error, element) {
            error.appendTo(element.parent().parent());
        },
        submitHandler:function(){
            $(".overlay").show();
            $('#noticeForm').ajaxSubmit({
                url:'/notice/save',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#noticeModal').modal('hide');
                        $('#noticeTable').jsGrid('loadData');
                        $('#noticeForm').validate().resetForm();
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
