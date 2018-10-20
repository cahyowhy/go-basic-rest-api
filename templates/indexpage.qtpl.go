// This file is automatically generated by qtc from "indexpage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line indexpage.qtpl:1
package templates

//line indexpage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line indexpage.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line indexpage.qtpl:2
type IndexPage struct {
}

//line indexpage.qtpl:7
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line indexpage.qtpl:7
	qw422016.N().S(`
	Todos Websites
`)
//line indexpage.qtpl:9
}

//line indexpage.qtpl:9
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line indexpage.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line indexpage.qtpl:9
	p.StreamTitle(qw422016)
	//line indexpage.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line indexpage.qtpl:9
}

//line indexpage.qtpl:9
func (p *IndexPage) Title() string {
	//line indexpage.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line indexpage.qtpl:9
	p.WriteTitle(qb422016)
	//line indexpage.qtpl:9
	qs422016 := string(qb422016.B)
	//line indexpage.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line indexpage.qtpl:9
	return qs422016
//line indexpage.qtpl:9
}

//line indexpage.qtpl:11
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
	//line indexpage.qtpl:11
	qw422016.N().S(`
    <div class="login-page">
        <form-login></form-login>  
    </div>
`)
//line indexpage.qtpl:15
}

//line indexpage.qtpl:15
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
	//line indexpage.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line indexpage.qtpl:15
	p.StreamBody(qw422016)
	//line indexpage.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line indexpage.qtpl:15
}

//line indexpage.qtpl:15
func (p *IndexPage) Body() string {
	//line indexpage.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line indexpage.qtpl:15
	p.WriteBody(qb422016)
	//line indexpage.qtpl:15
	qs422016 := string(qb422016.B)
	//line indexpage.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line indexpage.qtpl:15
	return qs422016
//line indexpage.qtpl:15
}

//line indexpage.qtpl:17
func (p *IndexPage) StreamCSSExternal(qw422016 *qt422016.Writer) {
	//line indexpage.qtpl:17
	qw422016.N().S(`
`)
//line indexpage.qtpl:18
}

//line indexpage.qtpl:18
func (p *IndexPage) WriteCSSExternal(qq422016 qtio422016.Writer) {
	//line indexpage.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line indexpage.qtpl:18
	p.StreamCSSExternal(qw422016)
	//line indexpage.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line indexpage.qtpl:18
}

//line indexpage.qtpl:18
func (p *IndexPage) CSSExternal() string {
	//line indexpage.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line indexpage.qtpl:18
	p.WriteCSSExternal(qb422016)
	//line indexpage.qtpl:18
	qs422016 := string(qb422016.B)
	//line indexpage.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line indexpage.qtpl:18
	return qs422016
//line indexpage.qtpl:18
}

//line indexpage.qtpl:20
func (p *IndexPage) StreamScriptExternal(qw422016 *qt422016.Writer) {
	//line indexpage.qtpl:20
	qw422016.N().S(`
`)
//line indexpage.qtpl:21
}

//line indexpage.qtpl:21
func (p *IndexPage) WriteScriptExternal(qq422016 qtio422016.Writer) {
	//line indexpage.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line indexpage.qtpl:21
	p.StreamScriptExternal(qw422016)
	//line indexpage.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line indexpage.qtpl:21
}

//line indexpage.qtpl:21
func (p *IndexPage) ScriptExternal() string {
	//line indexpage.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
	//line indexpage.qtpl:21
	p.WriteScriptExternal(qb422016)
	//line indexpage.qtpl:21
	qs422016 := string(qb422016.B)
	//line indexpage.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
	//line indexpage.qtpl:21
	return qs422016
//line indexpage.qtpl:21
}
