# goreplay

`goreplay` is a simple utility for exploring the gorethink driver and displaying retrieved items in the browser.

### Installation

```shell
go get -u github.com/mfridman/goreplay
```

### Basic usage

Navigate to app and open up `main.go`. Locate start of play comment and start playing with rethinkDB via the go driver [GoRethink](https://github.com/GoRethink/gorethink)

The default example displays an array of tables from database.

```go
/*
	start of play
*/

cur, err := r.TableList().Run(session)
if err != nil {
	http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	return
}

all := make([]string, 0)

if err := cur.All(&all); err != nil {
	http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	return
}

/*
	end of play
*/
```

Save file `main.go`, build/run app and access via browser at http://localhost:3001/playground

It's possible the database doesn't contain any tables, but if all went well, you should see an empty array `[]`

### Config options

Expecting a file named `config` in yaml format in the root directory

```shell
~/go/src/github.com/mfridman/goreplay
.
├── config
├── main.go
```

```yaml
# mandatory defaults

# rethinkdb connection
re_database: test
re_ip: localhost
re_port: 28015

# web app connection
http_address: localhost
http_port: 3001
```

### Optional

Use [gin](https://github.com/codegangsta/gin) for live reloading without building the binary each time.

```shell
go get -u github.com/codegangsta/gin
```

- Verify `gin` is installed by running: `gin -h`
- Then navigate to `goreplay` directory and run it via: `gin run main.go`
- The app will be accessible at http://localhost:3000/playground
    - **Not a type**. port 3000 is the default proxy port `gin` uses, `goreplay` is still listening/servering on port 3001