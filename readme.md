# zip
> Simple zip utility with a pluggable filesystem backend.

`go get github.com/jackmordaunt/zip`

The filesystem defaults to `afero.OsFs`, which operates on the host OS native 
filesystem.

Example use:
```go
func main() {
        _ = zip.Create("path/to/archive.zip", "path/to/source")
        _ = zip.Extract("path/to/achive.zip", "path/to/dest")
}
```

To set the filesystem backend: 
```go
func main() {
        filesystem := afero.NewMemMapFs()
        zip.SetFs(filesystem)
        
        _ = zip.Create("path/to/archive.zip", "path/to/source")
        _ = zip.Extract("path/to/achive.zip", "path/to/dest")
}
```