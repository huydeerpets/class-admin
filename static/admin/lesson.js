$(function() {
    initForm();
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#lesTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#lesTable').jsGrid('refresh');
    });
    $("#lesTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'number',title:'课程代码',width:60},
            {name:'name',title:'名称',sorting:false},
            {name:'credit',title:"学分",sorting:false},
            {name:'type',title:'类型',sorting:false},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改课程信息');
                    $('#lesModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除该课程信息吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/info/lesson/delLes?id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#lesTable').jsGrid('loadData');
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
                var number=$('input[name="numberS"]').val();
                var name=$('input[name="nameS"]').val();
                if(number!=""){filter.number=number;}
                if(name!=""){filter.name=name;}
                return $.ajax({
                    type: "GET",
                    url: "/info/lesson/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('新建学生信息');
        $('#lesModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#lesForm').validate().resetForm();
        $('#lesModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#lesTable').jsGrid('loadData');
    })
}

function initForm(){
    $("#lesForm").validate({
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
            $('#lesForm').ajaxSubmit({
                url:'/info/lesson/saveLes',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#lesModal').modal('hide');
                        $('#lesTable').jsGrid('loadData');
                        $('#lesForm').validate().resetForm();
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
    $('#number').val(item.number);
    $('#credit').val(item.credit);
    $('#type').val(item.type);
}
