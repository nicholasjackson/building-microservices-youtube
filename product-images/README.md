# Product Images

## Uploading 

Note: standard `-d` strips new lines

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @$PWD/go.mo
```

## Downloading with compression

```
curl -v localhost:9090/1/go.mod --compressed -o file.tmp
```