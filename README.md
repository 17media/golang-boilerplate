## 17 Media backend boilerplate

A boilerplate application for building backend service using go.


## Folder structure
- `controllers/` - defines your app routes and their logic
- `utils/` - code and functionality to be shared by different parts of the project
- `middlewares/` - Express middlewares which process the incoming requests before handling them down to the routes
- `models/` - represents data, implements business logic and handles storage
- `config/` - this folder contains configuration files, such as application settings, constants, ..., etc.
- `public/` - contains all static files like images, styles and javascript assets
- `glide.yaml` - dependencies pacakge management
- `glide.lock` - paired with `glide.yaml`
- `.editorconfig` - settings of editor such as indent size/style, charset, newline character, ..., etc. Please install EditorConfig to automatically apply settings
- `main.go` - app start from here

## Dependencies
1. run `go get github.com/tools/godep`
2. append `export GO15VENDOREXPERIMENT=1` in your `.zshrc`, `.bashrc`, `.bash_profile` or `.profile`
3. run `godep restore`
4. run `go install`
5. `go build`
6. execute the binary

**make sure you have GOPATH set up properly**

## Note
- To add new dependency, run `glide get github.com/Masterminds/cookoo`
- please put `_test.go` files adjacent to the file, e.g. `controllers/index.go` will have `controllers/index_test.go` for testing

## Todo
- Environment specific configs
