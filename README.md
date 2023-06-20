# practice-go

## commands

### test

```
go test -v practice-go/{package_name} -run {function_name}
```

### benchmark

-benchtimeで回数を指定できる

```
go test -bench {regexp} -benchmem {directory_path} -benchtime 000x
```

### module名変更

```
go mod edit -module practice-go
```

### go.modのバージョンアップ

```
go mod tidy -go="version"
```