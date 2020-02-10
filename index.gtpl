<!DOCTYPE html>
<html>
<head>
<title>Upload File</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://localhost:5000/upload/" method="post">
<input type="file" name="uploadfile" id="">
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="upload">
</form>
</body>
</html>