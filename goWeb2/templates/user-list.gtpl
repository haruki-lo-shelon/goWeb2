<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  <h2>ユーザ一覧</h2>
  <table class="list">
    <tr>
      <th>ID</th>
      <th>アカウント名</th>
      <th>ユーザ名</th>
      <th>登録日時</th>
    </tr>
    {{ range . }}
    <tr>
      <td>{{.Id}}</td>
      <td>{{.Account}}</td>
      <td>{{.Name}}</td>
      <td>{{.Created}}</td>
    </tr>
    {{ end }}
  </table>
  </form>
</body>
</html>
