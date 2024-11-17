# thumbnails

CLI for downloading youtube thumbnails.

```bash
# this look wrong, probably a local setup issue; not worth working out what...
go install ./cmd/main.go
mv $GOPATH/bin/main $GOPATH/bin/thumbnails
```

```bash
thumbnails --url="https://www.youtube.com/watch?v=2AwmwORrepc" --filename=tinzo.jpeg
```
