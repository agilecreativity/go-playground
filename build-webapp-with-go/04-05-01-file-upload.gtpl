<html>
  <head>
    <title>File Upload Example</title>
  </head>
  <body>
    <form enctype="multipart/form-data"
          action="http://localhost:9090/upload"
          method="post">

      <input type="file" name="uploadfile"/>
      <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="Upload"/>

    </form>
  </body>
</html>

