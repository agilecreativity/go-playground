<html>
<head>
  <title>Test Form</title>
</head>
<body>
  <form action="/login" method="post">

    Username:<input type="text" name="username">
    Password:<input type="password" name="password">

    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="Login">

  </form>
</body>
</html>
