$(function() {
    initTable();
    initBtnEvent();
});


function initTable(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#quesTable').jsGrid('option','pageSize',$('#pageSize').val());
        $('#quesTable').jsGrid('refresh');
    });
    $("#quesTable").jsGrid({
        pageSize:$('#pageSize').val(),
        fields:[
            {name:'les_no',title:'课程代码',sorting:false},
            {name:'les_name',title:'课程名称',sorting:false},
            {name:'title',title:'标题',width:150},
            {name:'stu_name',title:'学生姓名',sorting:false},
            {name:'updated_at',title:'发布时间',itemTemplate:function (value,item) {
                    return value.substr(0,10)
                }},
            {name:'operate',title:'操作',width:80,sorting:false,itemTemplate:function(value,item){
                var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'解答'}).on('click',function(){
                    window.open("/question/ansIndex?id="+item.id)
                });

                return $('<div>').addClass("operate").append($edit);
            }}
        ],
        controller: {
            loadData: function (filter) {
                var lessonNo=$('input[name="lessonNoS"]').val();
                var lessonName=$('input[name="lessonNameS"]').val();
                var stuName=$('input[name="stuNameS"]').val();
                if(lessonNo!=""){filter.lessonNo=lessonNo;}
                if(lessonName!=""){filter.lessonName=lessonName;}
                if(stuName!=""){filter.stuName=stuName;}
                return $.ajax({
                    type: "GET",
                    url: "/question/index?isajax=1",
                    data: filter
                });
            }
        }
    });
}

function initBtnEvent(){
    $('.closeBtn').on('click',function(){
        $('#quesForm').validate().resetForm();
        $('#quesModal').modal("hide");
    });
    $('#searchBtn').on('click',function(){
        $('#quesTable').jsGrid('loadData');
    })
}

