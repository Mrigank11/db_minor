# db_minor

## Build Instructions
- Install `golang`
- Clone this repo using 
```
git clone https://github.com/Mrigank11/db_minor
```
- To start server and download dependencies, run:
```
go run .
```
- To build a binary, run: 
```bash
# Make sure that you have `go-bindata` installed:
go get -u github.com/go-bindata/go-bindata/...
# Then, make a .go file from static files
go-bindata -o views/viewdata.go -pkg views templates templates/layout
# Create the binary
go build .
```
