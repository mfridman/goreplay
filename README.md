# goreplay

goreplay is a simple web app for exploring the gorethink driver and displaying retrieved items.

Installation:

```shell
go get -u github.com/mfridman/goreplay
```

Navigate to app directory and build/run the app. If using defaults, access the app at http://localhost:3001/playground

### Optional:

Use [gin](https://github.com/codegangsta/gin) for live reloading without building the binary each time.

```shell
go get -u github.com/codegangsta/gin
```

- Verify `gin` is installed by running: `gin -h`
- Then navigate to `goreplay` directory and run it via: `gin run main.go`
- The app will be accessible at http://localhost:3000/playground
    - **Not a type**. port 3000 is the default proxy port `gin` uses, `goreplay` is still listening/servering on port 3001