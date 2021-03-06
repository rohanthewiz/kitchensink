// This file is automatically generated by qtc from "mainpage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line mainpage.qtpl:1
package templates

//line mainpage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Main page template implements BasePage methods

//line mainpage.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line mainpage.qtpl:3
type MainPage struct {
	BasePage // Inherit so we can use BasePage functions for layout
	// CTX *fasthttp.RequestCtx // or any other attribute we want to pass
}

// Override any BasePage function we want to change
//

//line mainpage.qtpl:11
func (p *MainPage) StreamBody(qw422016 *qt422016.Writer) {
	//line mainpage.qtpl:11
	qw422016.N().S(`
  <h1>A Demonstration of a Full Stack App</h1>
  <strong>Techs in use:</strong>
  <ul>
    <li><strong>Go</strong> - The backend is written in Go (golang)</li>
    <li><strong>Iris</strong> - The fastest web framework for Go, period! (https://github.com/kataras/iris)
    <li><strong>FastHttp</strong> - The base router used is the fastest available for Go (https://github.com/valyala/fasthttp)</li>
    <li><strong>Quicktemplate</strong> - Templates are compiled right into the executable (very fast!) (https://github.com/valyala/quicktemplate)</li>
    <li><strong>Reactjs</strong> - The front end is written in cutting edge Reactjs</li>
    <li><strong>Custom select control</strong> - The select control is written from scratch and operates asynchronously with caching</li>
  </ul>
  <hr />
`)
//line mainpage.qtpl:23
}

//line mainpage.qtpl:23
func (p *MainPage) WriteBody(qq422016 qtio422016.Writer) {
	//line mainpage.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line mainpage.qtpl:23
	p.StreamBody(qw422016)
	//line mainpage.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line mainpage.qtpl:23
}

//line mainpage.qtpl:23
func (p *MainPage) Body() string {
	//line mainpage.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line mainpage.qtpl:23
	p.WriteBody(qb422016)
	//line mainpage.qtpl:23
	qs422016 := string(qb422016.B)
	//line mainpage.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line mainpage.qtpl:23
	return qs422016
//line mainpage.qtpl:23
}
