$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#stuTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#stuTable').jsGrid('refresh');
    });
    $("#stuTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'stu_id',title:'学号',width:60},
            {name:'name',title:'姓名',sorting:false},
            {name:'gender',title:'性别',sorting:false,itemTemplate:function(value){
                    switch(value){
                        case 1:return "男";
                        case 2:return "女";
                    }
                }},
            {name:'email',title:'邮箱',sorting:false},
            {name:'tel',title:"联系电话",sorting:false},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改学生信息');
                    $('#stuModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除该学生信息吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/info/student/delStu?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#stuTable').jsGrid('loadData');
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
                var stuId=$('input[name="stuIdS"]').val();
                var name=$('input[name="nameS"]').val();
                if(stuId!=""){filter.stuId=stuId;}
                if(name!=""){filter.name=name;}
                return $.ajax({
                    type: "GET",
                    url: "/info/student/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('新建学生信息');
        $('#stuModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $("#stuForm :input").not(":radio").val("");
        $('#stuForm').validate().resetForm();
        $('#stuModal').modal("hide");
    })
    $('#searchBtn').on('click',function(){
        $('#stuTable').jsGrid('loadData');
    })
}

function initForm(){
    $("#stuForm").validate({
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
            $('#stuForm').ajaxSubmit({
                url:'/info/student/saveStu',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#stuModal').modal('hide');
                        $('#stuTable').jsGrid('loadData');
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
    $('#stuId').val(item.stu_id);
    $('#gender').val(item.gender);
    $('#email').val(item.email);
    $('#tel').val(item.tel);
}
