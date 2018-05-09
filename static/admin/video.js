$(function(){
    $('#video-li').addClass("active");
    initInput();
    initShow();
    initBtnEvent();
});

//初始化视频列表
function initShow(){
    //切换每页大小
    $('#pageSize').on('change',function(){
        $('#videoShow').jsGrid('option','pageSize',$('#pageSize').val());
    });
    //分页展示照片
    $("#videoShow").jsGrid({
        pageSize:$('#pageSize').val(),
        controller: {
            loadData: function(filter) {
                var start=$('input[name="startDate"]').val();
                var end=$('input[name="endDate"]').val();
                var name=$('input[name="nameS"]').val();
                if(start!=""){filter.startDate=start;}
                if(end!=""){filter.endDate=end;}
                if(name!=""){filter.name=name;}
                return $.ajax({
                    type: "GET",
                    url: '/video/index?isajax=1',
                    data: filter
                });
            }
        },
        rowRenderer: function(item) {
            var $videobox = $("<div>").addClass("video-view box box-primary");
            var $video = $("<div>").addClass("box-body");
            var $img=$("<div>").addClass("video-img").on("click",function(){
                var videoObj = videojs("video");
                videoObj.src(item.url);
                $('#videoModal').modal();
            }).append($('<img>').attr("src",item.poster).attr("alt","无图片"));
            var $info = $("<div>").addClass("video-info")
                .append($("<p>").html("<strong>名称：</strong>"+item.name))
                .append($("<p>").html("<strong>时间：</strong>"+item.updated_at.substr(0,10)));

            var $edit=$('<span>').addClass("glyphicon glyphicon-edit").tooltip({title:'编辑'}).on('click',function(){
                fillForm(item);
                $('#modalLabel').text('编辑视频信息');
                $('#editModal').modal();
            });
            var $delete=$('<span>').addClass("glyphicon glyphicon-trash").tooltip({title:'删除'}).on('click',function(){
                deleteVideo(item.id)
            });
            var $button=$('<div>').addClass("operate").append($edit).append($delete);

            return $videobox.append($video.append($img).append($button).append($info));
        }

    });

}
//绑定按钮事件
function initBtnEvent(){
    //搜索按钮
    $('#videoSearch').on('click',function(){
        $('#videoShow').jsGrid('loadData');
    });
    $('#addBtn').on('click',function(){
        $('#modalLabel').text('添加视频');
        $('#editModal').modal();
    });
    $('.closeBtn').on('click',function(){
        $('#videoForm').validate().resetForm();
        $('#editModal').modal("hide");
    });
}
//初始化搜索输入框（日期选择、单选按钮）
function initInput(){
    $('#createTime').daterangepicker({
            startDate: moment().subtract(recentDays,'days'),
            endDate: moment(),
            maxDate : moment(),
            locale: datePickerLocale},
        function(start,end){
            $('input[name="startDate"]').val(start.format('YYYY-MM-DD'));
            $('input[name="endDate"]').val(end.format('YYYY-MM-DD'))
        });
    $("#videoForm").validate({
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
            $('#videoForm').ajaxSubmit({
                url:'/video/saveVideo',
                type:'post',
                success:function(data){
                    if(data.status){
                        $(".overlay").hide();
                        $('#editModal').modal('hide');
                        $('#videoShow').jsGrid('loadData');
                        $('#videoForm').validate().resetForm();
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

function fillForm(item) {
    $('#id').val(item.id);
    $('#name').val(item.name);
    $('#type').val(item.type);
    $('#brief').val(item.brief);
}

function deleteVideo(id) {
    Messager.confirm({message:"确定删除吗"}).on(function(e) {
        if (!e) return;
        $.ajax({
            url: '/video/delVideo?id=' + id,
            success: function (data) {
                if (data.status) {
                    $('#videoShow').jsGrid('loadData');
                } else {
                    Messager.alert(data.info);
                }
            },
            error:function () {
                Messager.alert("请求失败");
            }
        })
    })
}