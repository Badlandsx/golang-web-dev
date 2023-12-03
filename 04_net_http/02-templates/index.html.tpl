<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Net http golang with templates</title>
    <style>
        table { text-align: center; margin: 0 auto; border-collapse: collapse; border: 1px solid black; }
        table tr, table td { text-align: left; }
        table th { font-weight: bold; }
    </style>
</head>



<body>

<h1>Information about request:</h1>


<p>Method: {{ .Method }}</p>
<p>URL: {{ .URL.String }}</p>
<p>Host: {{ .Host }}</p>
{{if .Proto}}<p>Proto: {{ .Proto}}</p>{{end}}



{{if .Header}}
<table width="100%" cellpadding="0" cellspacing="0">
  <th>
    <td>Key</td>
    <td>Value</td>
  </th>
  {{range $key, $value := .Header }}
  <tr>
    <td>{{$key}}</td>
    <td>{{$value}}</td>
  </tr>
  {{end}}
</table>
{{end}}

<hr/>
<form action="/" method="POST">
        <input type="text" name="fname" placeholder="first name" autofocus autocomplete="off">
        <input type="submit" name="submit-btn" value="onda button">
</form>

{{if .Form}}
<h1>Form result:</h1>
<main>
    <p><strong>variable names</strong> (identifiers) and <em>values</em>:</p>
    {{range $key, $value := .Form }}
            <p><strong>{{$key}}</strong></p>
            <ul>{{range $value}}<li><em>{{.}}</em></li>{{end}}</ul>
    {{end}}
</main>
{{end}}

</body>
</html>