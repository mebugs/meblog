package constant

const (
	DocHtml = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="renderer" content="webkit">
    <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <title>Stone - API文档 - ReDoc版本</title>
	<meta name="description" content="Stone 物联基石项目 API文档 ReDoc版本"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="icon" href="/docs/file/icon.png"> 
  	<link rel="shortcut icon" href="/docs/file/icon.png"> 
  	<link rel="apple-touch-icon-precomposed" href="/docs/file/icon.png"> 
    <style>
        body {
            margin: 0;
            padding: 0;
        }
    </style>
</head>
<body>
<redoc spec-url='/docs/file/swagger.yaml' expand-responses="200,400,401,403,500" pagination="section"></redoc>
<script src="/docs/file/redoc.standalone.js"> </script>
</body>
</html>`

	SwaggerHtml = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="renderer" content="webkit">
  <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <title>Stone - API文档 - SwaggerUI版本</title>
  <meta name="description" content="Stone 物联基石项目 API文档 SwaggerUI版本"/>
  <link rel="icon" href="/docs/file/icon.png"> 
  <link rel="shortcut icon" href="/docs/file/icon.png"> 
  <link rel="apple-touch-icon-precomposed" href="/docs/file/icon.png"> 
  <link rel="stylesheet" href="/docs/file/swagger.ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="/docs/file/swagger.ui.bundle.js" crossorigin></script>
<script>
  window.onload = () => {
    window.ui = SwaggerUIBundle({
      url: '/docs/file/swagger.json',
      dom_id: '#swagger-ui',
    });
  };
</script>
</body>
</html>`
)
