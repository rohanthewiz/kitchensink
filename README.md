# A Full stack app demonstration - React frontend, Go backend
This is a work in-progress, but demonstrable here: http://gonotes.net:8094

### Technologies Employed

On the backend
- Go (golang)
- Iris - the fastest web framework for Go, period. https://github.com/kataras/iris
- The base router of Iris is Valyala's Fasthttp - https://github.com/valyala/fasthttp
- Valyala's QuickTemplate - https://github.com/valyala/quicktemplate

On the frontend
- React
- Stylus CSS preprocessing
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
- Point your browser to the address and port listed - typically `localhost:8094`