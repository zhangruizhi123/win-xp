<!DOCTYPE html>

<html>
<head>
   <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>某某管理系统</title>
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <link rel="stylesheet" href="/admin-lte/bower_components/bootstrap/dist/css/bootstrap.min.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/font-awesome/css/font-awesome.min.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/Ionicons/css/ionicons.min.css">
  <link rel="stylesheet" href="/admin-lte/dist/css/AdminLTE.min.css">
  <link rel="stylesheet" href="/admin-lte/dist/css/skins/_all-skins.min.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/morris.js/morris.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/jvectormap/jquery-jvectormap.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/bootstrap-datepicker/dist/css/bootstrap-datepicker.min.css">
  <link rel="stylesheet" href="/admin-lte/bower_components/bootstrap-daterangepicker/daterangepicker.css">
    <link rel="stylesheet" href="/admin-lte/icheck-1.x/skins/all.css">
  </head>

<body class="hold-transition login-page">


<div class="login-box">
      <div class="login-logo">
        <b>某某管理系统</b>
      </div><!-- /.login-logo -->
      <div class="login-box-body">
        <p class="login-box-msg">&nbsp;{{.msg}}</p>
        <form action="/admin/login.do" method="post">
          <div class="form-group has-feedback">
            <input type="text" name="userName" class="form-control" placeholder="用户名">
            <span class="glyphicon glyphicon-user form-control-feedback"></span>
          </div>
          <div class="form-group has-feedback">
            <input type="password" name="password" class="form-control" placeholder="密码">
            <span class="glyphicon glyphicon-lock form-control-feedback"></span>
          </div>
          <div class="row">
            <div class="col-xs-8">
                 <input tabindex="13" type="checkbox" id="flat-checkbox-1" class="save">
                 <label for="flat-checkbox-1">记住</label>
            </div><!-- /.col -->
            <div class="col-xs-4">
              <button type="submit" class="btn btn-primary btn-block btn-flat">登录</button>
            </div><!-- /.col -->
          </div>
        </form>
        <!--
        <div class="social-auth-links text-center">
          <p>- OR -</p>
          <a href="#" class="btn btn-block btn-social btn-facebook btn-flat"><i class="fa fa-facebook"></i> Sign in using Facebook</a>
          <a href="#" class="btn btn-block btn-social btn-google btn-flat"><i class="fa fa-google-plus"></i> Sign in using Google+</a>
        </div>
        -->
        <!-- /.social-auth-links -->

        <a href="#">忘记密码?</a><br>
        <a href="register.html" class="text-center">注册</a>

      </div><!-- /.login-box-body -->
    </div><!-- /.login-box -->

  
</body>

<!-- jQuery 3 -->
<script src="/admin-lte/bower_components/jquery/dist/jquery.min.js"></script>
<!-- jQuery UI 1.11.4 -->
<script src="/admin-lte/bower_components/jquery-ui/jquery-ui.min.js"></script>
<!-- Resolve conflict in jQuery UI tooltip with Bootstrap tooltip -->
<script>
  $.widget.bridge('uibutton', $.ui.button);
</script>
<!-- Bootstrap 3.3.7 -->
<script src="/admin-lte/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- Morris.js charts -->
<script src="/admin-lte/bower_components/raphael/raphael.min.js"></script>
<script src="/admin-lte/bower_components/morris.js/morris.min.js"></script>
<!-- Sparkline -->
<script src="/admin-lte/bower_components/jquery-sparkline/dist/jquery.sparkline.min.js"></script>

<!-- jQuery Knob Chart -->
<script src="/admin-lte/bower_components/jquery-knob/dist/jquery.knob.min.js"></script>
<!-- daterangepicker -->
<script src="/admin-lte/bower_components/moment/min/moment.min.js"></script>
<script src="/admin-lte/bower_components/bootstrap-daterangepicker/daterangepicker.js"></script>
<!-- datepicker -->
<script src="/admin-lte/bower_components/bootstrap-datepicker/dist/js/bootstrap-datepicker.min.js"></script>

<!-- Slimscroll -->
<script src="/admin-lte/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/admin-lte/bower_components/fastclick/lib/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/admin-lte/dist/js/adminlte.min.js"></script>
<script src="/admin-lte/dist/js/pages/dashboard.js"></script>
<script src="/admin-lte/dist/js/demo.js"></script>
<script src="/admin-lte/bower_components/icheck-1.x/icheck.min.js"></script>
<script>
 $(function () {
   //$('.save').iCheck();
  });
</script>

</html>
