<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>菜单管理</title>
  {{template "header_css.tpl"}}
</head>
<body>
  <div id="tree">
  
  <ul class="sidebar-menu" data-widget="tree">
        <!--
        <li class="header">MAIN NAVIGATION</li>
        -->
        <!-- 菜单管理 -->
        {{range $key,$val:=.menu}}
        <li class="treeview active menu-open">
         
          {{if $val.SubMenu}}
               <a href="#">
                <i class="fa fa-pie-chart"></i>
                <span>{{$val.Menu.Name}}</span>
                <span style="position:absolute;right:100px;">
                  <button class="btn btn-success btn-xs add" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-plus"></i>
                  </button>
                  <button class="btn btn-success btn-xs delete" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-minus"></i>
                  </button>
                   <button class="btn btn-success btn-xs edit" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-edit"></i>
                  </button>
                </span>
                
                
                <span class="pull-right-container">
                  <i class="fa fa-angle-left pull-right"></i>
                </span>
              </a>
          {{else}}
            <a href="#" active="{{$val.Menu.Url}}">
            <i class="{{$val.Menu.Icon}}"></i> 
            <span>{{$val.Menu.Name}}</span>
               <span style="position:absolute;right:100px;">
                  <button class="btn btn-success btn-xs add" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-plus"></i>
                  </button>
                  <button class="btn btn-success btn-xs delete" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-minus"></i>
                  </button>
                   <button class="btn btn-success btn-xs edit" parentId="{{$val.Menu.Id}}">
                  <i class="glyphicon glyphicon-edit"></i>
                  </button>
                </span>
            </a>
          {{end}}
          <ul class="treeview-menu">
            {{range $key2,$val2:=$val.SubMenu}}
                 <li><a href="#" active="{{$val2.Url}}"><i class="{{$val2.Icon}}"></i>{{$val2.Name}}</a></li>
            {{end}}
          </ul>
        </li>
        {{end}}
        
      </ul>
  </div>
  
  
  <div class="form">
      <input type="text" class="" >
      <input type="text" class="" >
      <input type="text" class="" >
  </div>
  
  
  
  <!-- 引入外部的js -->
{{template "header_js.tpl"}}


<script>
$(function(){
  var $dig=$(".form")
  $(".treeview").on("click",".add",function(event){
    event.stopPropagation();
     BootstrapDialog.show({
            title: '加载远程页面',
            message: function (dialog) {
                var $message = $dig;
                return $message;
            },
            size: BootstrapDialog.SIZE_WIDE,
            cssClass: "fade",
            closeable: true,
        });
    console.log("添加");
  })
  
   $(".treeview").on("click",".edit",function(event){
    console.log("修改");
    BootstrapDialog.confirm('Hi Apple, are you sure?', function(result){
            if(result) {
                alert('Yup.');
            }else {
                alert('Nope.');
            }
        });
    event.stopPropagation();
  })
  $(".treeview").on("click",".delete",function(event){
    console.log("删除");
    event.stopPropagation();
  })

});


</script>
</body>

