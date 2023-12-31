<h1>Swag-Demo</h1>
<h2>Demonstration repo to show how <a href="https://github.com/swaggo/swag">swaggo</a> library works</h2>

Here I provided a simple http server with several handlers. It works with the global array of <i>items</i> (imitation of database interaction) to mimic basic CRUD operations.

<h3>Installation</h3>
To start with you need to get a <a href="https://github.com/swaggo/swag">swaggo</a> package

```
go install github.com/swaggo/swag/cmd/swag@latest
```

Unfortunately I could not access it directly from console as a regular util so I had to provide the absolute path to the app (I hope you won`t face this problem and your $PATH will work as intended so further I will to refer to this util as <i>swag</i> as it should be):

```
~/go/bin/swag --version
```
</br>
You also need <a href="https://github.com/swaggo/http-swagger">http-swagger</a> to register it in the handler

```
go get -u github.com/swaggo/http-swagger
```

<b>NOTE:</b> If you use web framework try to find the corresponding package from the <a href="https://github.com/swaggo/swag?tab=readme-ov-file#supported-web-frameworks">list</a>.

<h3>Usage</h3>
<p>Swaggo is based on the declarative comments. You start with most basic API infromation e.g 
<ul>
    <li>API version</li> 
    <li>Project name</li> 
    <li>Contact information</li> 
    <li>Licence</li>
    <li>Host</li>
    <li>Root path</li>
</ul>

<pre>
// @title Project Name
// @version 0.0.1
// @host localhost:3000
// @BasePath /
// @contact.name company_name
// @contact.url http://www.link-to-support.io/support
// @contact.email some.company.name@company.domain.com
func main(){
    ...
}
</pre>

Then you need to provide docs for methods in a similar way:
<pre>
// @title list godoc
// @Summary      list items
// @Description  list all items from global var
// @Tags         items
// @Produce      json
// @Success      200  {object}  []models.Item
// @Router       /list [get]“
func list(w http.ResponseWriter, r *http.Request){
    ...
}
</pre>

Please refer to <a href="https://github.com/swaggo/swag?tab=readme-ov-file#declarative-comments-format">swaggo documentation</a> for more available attributes.</br>
<b>NOTE:</b> swaggo generaters models automatically, but you can also hide some fields by using fields attributes.</br>

After you wrote your comments you can format them by calling:

```
swag fmt
```

Then you need to run 

```
swag init
```

To generate docs files.</br>

Finally you can run your server to check swagger works and shows data correctly. Run

```
go run .
```

and follow http://localhost:3000/swagger/.