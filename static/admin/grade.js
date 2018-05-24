$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#classTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#classTable').jsGrid('refresh');
    });
    $("#classTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'stu_no',title:'学生学号',width:60},
            {name:'stu_name',title:'学生姓名',sorting:false},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'tea_name',title:'教师姓名',sorting:false},
            {name:'term',title:'学期',sorting:false},
            {name:'score',title:'成绩'},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改排课信息');
                    $('#classModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/info/class/del?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#classTable').jsGrid('loadData');
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
                var stuNo=$('input[name="stuNoS"]').val();
                if(stuNo!=""){filter.stuNo=stuNo;}
                return $.ajax({
                    type: "GET",
                    url: "/info/class/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('新建选课信息');
        $('#classModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#classForm').validate().resetForm();
        $('#classModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#classTable').jsGrid('loadData');
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
    $("#classForm").validate({
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
            $('#classForm').ajaxSubmit({
                url:'/grade/import',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#classModal').modal('hide');
                        $('#classTable').jsGrid('loadData');
                        $('#classForm').validate().resetForm();
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
    $('#lectureId').val(item.lecture_id);
    $('#stuNo').val(item.stu_no);
    $('#grade').val(item.score);
}
