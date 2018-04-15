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
                        <span class="hidden-xs">{{.session.username}}</span>
                    </a>
                    <ul class="dropdown-menu">
                        <li class="user-footer">
                            <div class="pull-right">
                                <a href="#" class="btn btn-default btn-flat">Sign out</a>
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
        <ul class="sidebar-menu" data-widget="tree">
            <li class="treeview" id="user-li">
                <a href="#">
                    <i class="fa fa-user"></i><span>信息管理</span>
                    <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
                </a>
                <ul class="treeview-menu">
                    <li id="owner-li"><a href="/user/user"><i class="fa fa-male"></i> 学生信息</a></li>
                    <li id="card-li"><a href="/user/card"><i class="fa fa-paw"></i> 教师信息 </a></li>
                    <li id="owner-li"><a href="/user/user"><i class="fa fa-male"></i> 课程信息</a></li>
                </ul>
            </li>

            <li id="act-li"><a href="/act/act"><i class="fa fa-bullhorn"></i><span> 课程资料</span></a></li>
        </ul>
    </section>
</aside>
