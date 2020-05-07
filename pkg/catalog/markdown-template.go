package catalog

const NewProjectTemplate = `
| Package |
----|{{range $key, $val := ModuleAsPackages .}}
[{{$key}}]({{$key}}/README.md)|{{end}}

<img src="{{CreateIntegrationDiagram . "" false}}">

<img src="{{CreateIntegrationDiagram . "" true}}">

`

const NewPackageTemplate = `
[Back](../README.md)
{{$packageName := ModulePackageName .}}

# {{$packageName}}

## Service Index
| Service Name | Method |
----|----{{$Apps := .Apps}}{{range $appName := AlphabeticalApps .Apps}}{{$app := index $Apps $appName}}{{$Endpoints := $app.Endpoints}}{{range $endpointName := AlphabeticalEndpoints $Endpoints}}{{$endpoint := index $Endpoints $endpointName}}
{{$appName}} | [{{$endpoint.Name}}](#{{$appName}}-{{SanitiseOutputName $endpoint.Name}}){{end}}{{end}}



![]({{CreateIntegrationDiagram . $packageName false}})

{{range $appName := AlphabeticalApps .Apps}}{{$app := index $Apps $appName}}
{{if eq (hasPattern $app.Attrs "ignore") false}}
# {{$appName}}
{{AppComment $app}}
{{range $e := $app.Endpoints}}
{{if eq (hasPattern $e.Attrs "ignore") false}}
## {{$appName}} {{SanitiseOutputName $e.Name}}
{{Attribute "description" $e.GetAttrs}}

![]({{CreateSequenceDiagram $appName $e}})

### Request types
{{range $param := $e.Param}}
![]({{CreateParamDataModel $app $param}})
{{end}}

{{with $e.RestParams}}

{{range $param := $e.RestParams.QueryParam}}
![]({{CreateQueryParamDataModel $appName $param}})
{{end}}

{{range $param := $e.RestParams.UrlParam}}
![]({{CreatePathParamDataModel $appName $param}})
{{end}}

{{end}}

### Response types
{{range $s := $e.Stmt}}{{$diagram := CreateReturnDataModel $appName $s $e}}{{if ne $diagram ""}}
![]({{$diagram}})
{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}


{{range $appName := AlphabeticalApps .Apps}}{{$app := index $Apps $appName}}
{{if hasPattern $app.Attrs "db"}}

## Database
![]({{GenerateDataModel $app}})
{{end}}{{end}}


### Types

<table>
<tr>
<th>App Name</th>
<th>Diagram</th>
<th>Comment</th>
<th>Full Diagram</th>
{{range $appName := AlphabeticalApps .Apps}}{{$app := index $Apps $appName}}{{$types := $app.Types}}
{{if ne (hasPattern $app.Attrs "db") true}}
</tr>

{{range $typeName := AlphabeticalTypes $types}}{{$type := index $types $typeName}}
<tr>
<td>

{{$appName}}.<br>{{$typeName}}
</td>
<td>

<img src="{{CreateTypeDiagram  $app $typeName $type false}}">
</td>
<td> 

{{if ne (TypeComment $type) ""}}<details closed><summary>Comment</summary><br>{{TypeComment $type}}</details>{{end}} 
</td>
<td>

<a href="{{CreateTypeDiagram  $app $typeName $type true}}">Link</a>
</td>
</tr>{{end}}{{end}}{{end}}
</table>

`
