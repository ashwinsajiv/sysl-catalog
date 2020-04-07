# syslcatalog

A markdown/html + Diagram generator for sysl specifications

## Installation

```bash
go get -u -v github.com/anz-bank/sysl-catalog
```

## How to use
1. Set up environment
`export SYSL_PLANTUML=http://www.plantuml.com/plantuml`

2. Run 

```bash
sysl-catalog -o <output directory> <input.sysl>
```
- You can optionally specify the `--type=html` if you want to generate html instead of markdown, which is useful for use with github pages

3. That's it (basically!)

    This will generate markdown with integration diagrams + sequence diagrams + data model diagrams as seen in [demo/markdown/README.md](demo/markdown/README.md) or see html generation at [demo/html/index.html](demo/html/index.html).


## Server Mode
sysl-catalog comes with a `serve` mode which will serve on port `:69` by default

```bash 
sysl-catalog -o <output directory> --server <input directory>
```
This will start a server and filewatchers to watch the input file and its directories recursively, the website will refresh on edit
![example gif](resources/example.gif)

## Requirements
In [demo/markdown/README.md](demo/markdown/README.md) we have an example with a couple of interesting parts:

1. There needs to be a sysl `project` the same name as the filename:

so as this example is called "simple2.sysl" there needs to be a `project` "simple2":
- applications in our integration diagram:
```
simple2[appfmt="%(appname)", ~ignore]:
    _:
        MobileApp
        Server
        MegaDatabase
```
NOTE: must have `appfmt="%(appname)"` as an attribute to render integration diagrams correctly.

2. `@package` attribute must be specified:
- This will create a markdown page for `ApplicationPackage` as seen in [demo/markdown/ApplicationPackage/README.md](demo/markdown/ApplicationPackage/README.md).
 Currently the package name is not inferred from the application name (`MobileApp`), so this needs to be added (`ApplicationPackage`).
```
MobileApp:
    @package = "ApplicationPackage"
    Login(input <: Server.Request):
        Server <- Authenticate
        return ok <: MegaDatabase.Empty
```

3. Application names need to be prefixed to parameter types if the type is defined in another application, since defined parameters are under scope of the application it is defined in:
```diff
MobileApp:
    @package = "ApplicationPackage"
+    Login(input <: Server.Request):
-    Login(input <: Request):
        Server <- Authenticate
        return ok <: MegaDatabase.Empty
```

4. Add `~ignore` to applications/projects that are to be ignored in the markdown creation
```diff
ThisAppShouldntShow[~ignore]:
    NotMySystem:
        ...
# Or ignore only specific endpoints
ThisAppShouldShow[~ignore]:
    NotMySystem[~ignore]:
        ...
```

## Screenshots
![resources/project_view.png](resources/project_view.png)
*project_view*

![resources/package_view.png](resources/package_view.png)
*package_view*
