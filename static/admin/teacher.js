$(function() {
    initForm();
    initTable();
    initBtnEvent();
});

function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#teaTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#teaTable').jsGrid('refresh');
    });
    $("#teaTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'number',title:'编号',width:60},
            {name:'name',title:'姓名',sorting:false},
            {name:'gender',title:'性别',sorting:false,itemTemplate:function(value){
                    switch(value){
                        case 1:return "男";
                        case 2:return "女";
                    }
                }},
            {name:'email',title:'邮箱',sorting:false},
            {name:'tel',title:"联系电话",sorting:false},
            {name:'office',title:"办公室",sorting:false},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                    fillForm(item);
                    $('#modalLabel').text('修改教师信息');
                    $('#teaModal').modal();
                });
                var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                    Messager.confirm({message:"确定删除该教师信息吗"}).on(function(e) {
                        if (!e) return;
                        $.ajax({
                            url: '/info/teacher/delTea?Id=' + item.id,
                            success: function (data) {
                                if (data.status) {
                                    $('#teaTable').jsGrid('loadData');
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
                    url: "/info/teacher/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('新建教师信息');
        $('#teaModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#teaForm').validate().resetForm();
        $('#teaModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#teaTable').jsGrid('loadData');
    })
}

function initForm(){
    $("#teaForm").validate({
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
            $('#teaForm').ajaxSubmit({
                url:'/info/teacher/saveTea',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#teaModal').modal('hide');
                        $('#teaTable').jsGrid('loadData');
                    }else{
                        $(".overlay").hide();
                        Messager.alert(data.info);
                    }
                },
                error:function(){
                    $(".overlay").hide();
                    Messager.alert("请求失败");
                }
            })
        }
    });
}
function fillForm(item){
    $('#id').val(item.id);
    $('#name').val(item.name);
    $('#number').val(item.number);
    $('#gender').val(item.gender);
    $('#office').val(item.office);
    $('#email').val(item.email);
    $('#tel').val(item.tel);
}
