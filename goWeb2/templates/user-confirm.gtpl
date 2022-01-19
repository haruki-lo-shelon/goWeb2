<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  <h2>登録内容の確認</h2>
  <form action="user-registered" method="post">
    <table>
      <tr>
        <td>アカウント名：</td>
        <td>{{.account}}</td>
      </tr>
      <tr>
        <td>お名前：</td>
        <td>{{.name}}</td>
      </tr>
      <tr>
        <td>パスワード：</td>
        <td>{{.passwd}}</td>
      </tr>
    </table>
    <input type="hidden" name="account" value="{{.hid_account}}">
    <input type="hidden" name="name" value="{{.hid_name}}">
    <input type="hidden" name="passwd" value="{{.hid_passwd}}">
    <button type="button" class="btn-push" onclick="history.back()">入力画面に戻る</button>
    <input type="submit" class="btn-push btn-push-blue" value="この内容で登録する">
  </form>
</body>
</html>
