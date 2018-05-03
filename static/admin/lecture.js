$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#lecTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#lecTable').jsGrid('refresh');
    });
    $("#lecTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'lesson_no',title:'课程代码',width:60},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'teacher_no',title:'教师编号',sorting:false},
            {name:'tea_name',title:'教师姓名',sorting:false},
            {name:'term',title:'学期',sorting:false},
            {name:'place',title:"上课地点",sorting:false},
            {name:'people',title:"人数",sorting:false},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改排课信息');
                    $('#lecModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/info/lecture/delLec?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#lecTable').jsGrid('loadData');
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
                var teacherNo=$('input[name="teacherNoS"]').val();
                var lessonNo=$('input[name="lessonNoS"]').val();
                var term=$('input[name="termS"]').val();
                if(teacherNo!=""){filter.teacherNo=teacherNo;}
                if(lessonNo!=""){filter.lessonNo=lessonNo;}
                if(term!=""){filter.term=term;}
                return $.ajax({
                    type: "GET",
                    url: "/info/lecture/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('新建排课信息');
        $('#lecModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#lecForm').validate().resetForm();
        $('#lecModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#lecTable').jsGrid('loadData');
    })
}

function initForm(){
    $("#lecForm").validate({
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
            $('#lecForm').ajaxSubmit({
                url:'/info/lecture/saveLec',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#lecModal').modal('hide');
                        $('#lecTable').jsGrid('loadData');
                        $('#lecForm').validate().resetForm();
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
    $('#lessonNo').val(item.lesson_no);
    $('#teacherNo').val(item.teacher_no);
    $('#place').val(item.place);
    $('#term').val(item.term);
    $('#classTime').val(item.class_time);
    $('#people').val(item.people);
}
