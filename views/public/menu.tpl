
<script type="text/javascript">
    $( function() {
        if({{.tree}}!=""){
            localStorage.setItem('menu', JSON.stringify({{.tree}}));
        }
        if({{.userinfo.Nickname}}!=""){
            localStorage.setItem('username', {{.userinfo.Nickname}});
        }
        var username=localStorage.getItem("username");
        $('span.hidden-xs').html(username);
        var tree=JSON.parse(localStorage.getItem("menu"));
        $('#menu').empty();
        $.each(tree,function(i,item){
            if(item.children.length==0){
                $('#menu').append($('<li id="'+item.attributes.div_id+'"><a href="'+item.attributes.url+'"><i class="'+item.iconCls+'"></i>'+item.text+'</a></li>'));
                return;
            }
            var $li=$('<li class="treeview" id="'+item.attributes.div_id+'">');
            var $a=$('<a href="#">').append($('<i>').addClass(item.iconCls)).append($('<span>'+item.text+'</span>')).append(
                $('<span class="pull-right-container"><i class="fa fa-angle-left pull-right"></i></span>')
            );
            var $ul=$('<ul class="treeview-menu">');
            $.each(item.children,function(i,val){
                $ul=$ul.append($('<li id="'+val.attributes.div_id+'"><a href="'+val.attributes.url+'"><i class="'+val.iconCls+'"></i>'+val.text+'</a></li>'))
            });
            $('#menu').append($li.append($a).append($ul));
        });
        activeDiv();
    });
    activeDiv=function () {
    }
</script>
<header class="main-header">
    <div class="logo">
        <span class="logo-mini"><b>课堂互动</b></span>
        <span class="logo-lg"><b>课堂互动平台</b>后台管理</span>
    </div>
    <nav class="navbar navbar-static-top">
        <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
            <span class="sr-only">Toggle navigation</span>
        </a>

        <div class="navbar-custom-menu">
            <ul class="nav navbar-nav">
                <li class="dropdown user user-menu">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                        <span class="hidden-xs"></span>
                    </a>
                    <ul class="dropdown-menu">
                        <li class="user-footer">
                            <div class="pull-right">
                                <a href="/public/logout" class="btn btn-default btn-flat">Sign out</a>
                            </div>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
    </nav>
</header>
<aside class="main-sidebar">
    <section class="sidebar">
        <ul class="sidebar-menu" data-widget="tree" id="menu">
        </ul>
    </section>
</aside>
