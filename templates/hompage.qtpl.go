// This file is automatically generated by qtc from "hompage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line hompage.qtpl:1
package templates

//line hompage.qtpl:1
import "go-basic-rest-api/models"

//line hompage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line hompage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line hompage.qtpl:4
type HomePage struct {
	User     models.User
	UserJSON []byte
}

//line hompage.qtpl:10
func (p *HomePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line hompage.qtpl:10
	qw422016.N().S(`
	Beranda `)
	//line hompage.qtpl:11
	qw422016.E().S(p.User.Username)
	//line hompage.qtpl:11
	qw422016.N().S(`
`)
//line hompage.qtpl:12
}

//line hompage.qtpl:12
func (p *HomePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line hompage.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line hompage.qtpl:12
	p.StreamTitle(qw422016)
	//line hompage.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line hompage.qtpl:12
}

//line hompage.qtpl:12
func (p *HomePage) Title() string {
	//line hompage.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line hompage.qtpl:12
	p.WriteTitle(qb422016)
	//line hompage.qtpl:12
	qs422016 := string(qb422016.B)
	//line hompage.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line hompage.qtpl:12
	return qs422016
//line hompage.qtpl:12
}

//line hompage.qtpl:14
func (p *HomePage) StreamBody(qw422016 *qt422016.Writer) {
	//line hompage.qtpl:14
	qw422016.N().S(`
    <div class="home-page">
        <user-home user="`)
	//line hompage.qtpl:16
	qw422016.E().V(string(p.UserJSON))
	//line hompage.qtpl:16
	qw422016.N().S(`"></user-home>  
    </div>
`)
//line hompage.qtpl:18
}

//line hompage.qtpl:18
func (p *HomePage) WriteBody(qq422016 qtio422016.Writer) {
	//line hompage.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line hompage.qtpl:18
	p.StreamBody(qw422016)
	//line hompage.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line hompage.qtpl:18
}

//line hompage.qtpl:18
func (p *HomePage) Body() string {
	//line hompage.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line hompage.qtpl:18
	p.WriteBody(qb422016)
	//line hompage.qtpl:18
	qs422016 := string(qb422016.B)
	//line hompage.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line hompage.qtpl:18
	return qs422016
//line hompage.qtpl:18
}

//line hompage.qtpl:20
func (p *HomePage) StreamCSSExternal(qw422016 *qt422016.Writer) {
	//line hompage.qtpl:20
	qw422016.N().S(`
`)
//line hompage.qtpl:21
}

//line hompage.qtpl:21
func (p *HomePage) WriteCSSExternal(qq422016 qtio422016.Writer) {
	//line hompage.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line hompage.qtpl:21
	p.StreamCSSExternal(qw422016)
	//line hompage.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line hompage.qtpl:21
}

//line hompage.qtpl:21
func (p *HomePage) CSSExternal() string {
	//line hompage.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
	//line hompage.qtpl:21
	p.WriteCSSExternal(qb422016)
	//line hompage.qtpl:21
	qs422016 := string(qb422016.B)
	//line hompage.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
	//line hompage.qtpl:21
	return qs422016
//line hompage.qtpl:21
}

//line hompage.qtpl:23
func (p *HomePage) StreamScriptExternal(qw422016 *qt422016.Writer) {
	//line hompage.qtpl:23
	qw422016.N().S(`
`)
//line hompage.qtpl:24
}

//line hompage.qtpl:24
func (p *HomePage) WriteScriptExternal(qq422016 qtio422016.Writer) {
	//line hompage.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line hompage.qtpl:24
	p.StreamScriptExternal(qw422016)
	//line hompage.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line hompage.qtpl:24
}

//line hompage.qtpl:24
func (p *HomePage) ScriptExternal() string {
	//line hompage.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
	//line hompage.qtpl:24
	p.WriteScriptExternal(qb422016)
	//line hompage.qtpl:24
	qs422016 := string(qb422016.B)
	//line hompage.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
	//line hompage.qtpl:24
	return qs422016
//line hompage.qtpl:24
}
