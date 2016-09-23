# A Full stack app demonstration - React frontend, Go backend
(This is a somewhat functional work in-progress)

### Technologies Employed

On the backend
- Go (golang)
- Valyala Fasthttp - https://github.com/valyala/fasthttp
- Valyala QuickTemplate - https://github.com/valyala/quicktemplate

On the frontend
- React
- Node for asset compilation

### Setup and Compilation of React components
You will need to have `nodejs` and `npm` already installed. Follow instructions at https://nodejs.org/
- Change to the *project root*: `cd kitchensink`
- Install project dependencies:  `npm install`
- Install browserify if not installed `npm install -g browserify`
- You can watch for changed files or manually build.
    - To automatically build when files in `src` change:
        - Open another terminal and change directory to the project root
        - Run `npm run watch`
    - To manually build: `npm run build`

### Setup and Compilation - backend
- If changes are made to templates, you need to run `qtc` in the templates folder.
See https://github.com/valyala/quicktemplate#quick-start
- If changes are made to any `*.go` files, run `go build` at the project root

### Starting the server
- `./kitchensink`
- Point your browser to the address and port listed - typically `localhost:8080`