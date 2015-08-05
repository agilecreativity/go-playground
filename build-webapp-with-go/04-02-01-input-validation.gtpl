<html>
<head>
  <title>Test Form</title>
</head>
<body>
  <form action="/login" method="post">
    Username : <input type="text"     name="username"><br/>
    Password : <input type="password" name="password"><br/>
    Age      : <input type="text"     name="age"><br/>
    <!-- sample checkbox -->
    <select name="fruit">
      <option value="apple">apple</option>
      <option value="pear">pear</option>
      <option value="banana">banana</option>
    </select>

    <!-- sample radio button -->
    <input type="radio" name="gender" value="M">Male</input>
    <input type="radio" name="gender" value="F">Female</input>

    <input type="submit" value="login">
  </form>
</body>
</html>

