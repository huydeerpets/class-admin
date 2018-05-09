package rbac

import (
	"github.com/astaxie/beego"
	m "class-admin/models"
	"github.com/astaxie/beego/orm"
)

type CommonController struct {
	beego.Controller
	Templatetype string //ui template type
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *CommonController) GetTemplatetype() string {
	return ""
}

func (this *CommonController) GetTree(user m.User) []Tree {
	var nodes []orm.Params
	if user.Username=="admin"{
		nodes, _ = m.GetNodeTree(0, 1)
	}else {
		nodes, _ = m.NodeTreeByUser(user.Id)
	}

	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		tree[k].IconCls = v["Logo"].(string)
		tree[k].Attributes.DivId = v["DivId"].(string)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2)
		if len(children)==0{
			tree[k].Attributes.Url="/"+v["Name"].(string)
		}
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
			tree[k].Children[k1].IconCls = v1["Logo"].(string)
			tree[k].Children[k1].Attributes.DivId = v1["DivId"].(string)
		}
	}
	return tree
}

func init() {

	//验证权限
	AccessRegister()
}
