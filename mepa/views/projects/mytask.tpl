<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    {{template "inc/my-nav.tpl" .}}
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 任务 / 总数：{{.countTask}}<span class="tools pull-right">
			<a href="/my/task?filter=accept" class="btn btn-default {{if eq .condArr.filter "accept"}}active{{end}}" style="padding:6px;">指派给我</a>
			<a href="/my/task?filter=create" class="btn btn-default {{if eq .condArr.filter "create"}}active{{end}}" style="padding:6px;">由我创建</a>
			<a href="/my/task?filter=complete" class="hidden-xs btn btn-default {{if eq .condArr.filter "complete"}}active{{end}}" style="padding:6px;">由我解决</a>	
			<a href="/my/task?filter=close" class="hidden-xs btn btn-default {{if eq .condArr.filter "close"}}active{{end}}" style="padding:6px;">由我关闭</a>	
			<a href="/my/task?filter=cancel" class="hidden-xs btn btn-default {{if eq .condArr.filter "cancel"}}active{{end}}" style="padding:6px;">由我取消</a>			
			<a href="javascript:;" class="fa fa-chevron-down"></a>
            </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>名称</th>
                        <th>状态</th>
                        <th>截止日期</th>
                        <th>指派给</th>
                        <th>完成者</th>
                        <th>预工时</th>
                        <th>需求</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .tasks}}
                    <tr>
                      <td><a href="/task/show/{{$v.Id}}">{{getProjectname $v.Projectid}}-{{$v.Name}}</a></td>
                      <td>{{getTaskStatus $v.Status}}</td>
                      <td>{{getDate $v.Ended}}</td>
                      <td><a href="/user/show/{{$v.Acceptid}}">{{getRealname $v.Acceptid}}</a></td>
                      <td><a href="/user/show/{{$v.Completeid}}">{{getRealname $v.Completeid}}</a></td>
                      <td>{{$v.Tasktime}}</td>
                      <td><a href="/need/show/{{$v.Needsid}}">{{getNeedsname $v.Needsid}}</a></td>
                      <td><a href="#acceptModal" data-toggle="modal" data-id="{{$v.Id}}">指派</a> <a href="javascript:;" data-id="{{$v.Id}}" class="js-task-status" data-status="2">开始</a> <a href="javascript:;" data-id="{{$v.Id}}" class="js-task-status" data-status="3">完成</a> <a href="/task/edit/{{$v.Id}}">编辑</a> </td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <div aria-hidden="true" aria-labelledby="acceptModalLabel" role="dialog" tabindex="-1" id="acceptModal" class="modal fade">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
            <h4 class="modal-title">任务指派给?</h4>
          </div>
          <div class="modal-body">
            <select id="acceptid" class="form-control">
              <option value="">请选择指派给</option>
              
				{{range .teams}}
					
              <option value="{{.Userid}}">{{getRealname .Userid}}</option>
              
					{{end}}
				
            </select>
            <p></p>
            <textarea id="note" placeholder="备注说明" style="height:90px;" class="form-control"></textarea>
          </div>
          <div class="modal-footer">
            <input type="hidden" id="taskid" />
            <button data-dismiss="modal" class="btn btn-default" type="button">取消</button>
            <button class="btn btn-primary js-dialog-taskaccept" type="button">提交</button>
          </div>
        </div>
      </div>
    </div>
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script>
$(function(){
	$('#acceptModal').on('show.bs.modal', function (e) {
		$('#taskid').val($(e.relatedTarget).attr('data-id'))
	})
		
})
</script>
</body>
</html>
