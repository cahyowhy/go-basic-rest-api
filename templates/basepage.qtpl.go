// This file is automatically generated by qtc from "basepage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.

//line basepage.qtpl:2
package templates

//line basepage.qtpl:2
import "os"

//line basepage.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line basepage.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line basepage.qtpl:4
type Page interface {
	//line basepage.qtpl:4
	MetaInfo() string
	//line basepage.qtpl:4
	StreamMetaInfo(qw422016 *qt422016.Writer)
	//line basepage.qtpl:4
	WriteMetaInfo(qq422016 qtio422016.Writer)
	//line basepage.qtpl:4
	Body() string
	//line basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
	//line basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
	//line basepage.qtpl:4
	CSSExternal() string
	//line basepage.qtpl:4
	StreamCSSExternal(qw422016 *qt422016.Writer)
	//line basepage.qtpl:4
	WriteCSSExternal(qq422016 qtio422016.Writer)
	//line basepage.qtpl:4
	ScriptExternal() string
	//line basepage.qtpl:4
	StreamScriptExternal(qw422016 *qt422016.Writer)
	//line basepage.qtpl:4
	WriteScriptExternal(qq422016 qtio422016.Writer)
//line basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line basepage.qtpl:13
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	//line basepage.qtpl:13
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="x-ua-compatible" content="ie=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/public/css/app.css">
    <link rel='shortcut icon' type='image/x-icon' href='/public/images/favicon.ico' />
    <link rel="manifest" href="/public/manifest.json">
    <link rel="stylesheet" href="https://cdn.materialdesignicons.com/3.0.39/css/materialdesignicons.min.css">
    <meta name="og:site_name" content="Todo App"/>
    <script type="text/javascript" src="/public/vendor/js/jquery.min.js"></script>
    <script type="text/javascript" src="/public/vendor/js/jquery.migrate.min.js"></script>
    <script type="text/javascript" src="/public/js/app.js" defer></script>
    <script type="text/javascript" src="/public/js/sw-controller.js"></script>
    <script>
        document.addEventListener("DOMContentLoaded", function() {
            new SW();
        });
    </script>
    `)
	//line basepage.qtpl:35
	qw422016.N().S(p.CSSExternal())
	//line basepage.qtpl:35
	qw422016.N().S(`
    `)
	//line basepage.qtpl:36
	qw422016.N().S(p.MetaInfo())
	//line basepage.qtpl:36
	qw422016.N().S(`
</head>
<body>
	<div id="app">
        `)
	//line basepage.qtpl:40
	qw422016.N().S(p.Body())
	//line basepage.qtpl:40
	qw422016.N().S(`
    </div>
    `)
	//line basepage.qtpl:42
	qw422016.N().S(p.ScriptExternal())
	//line basepage.qtpl:42
	qw422016.N().S(`
    `)
	//line basepage.qtpl:43
	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "DEV" {
		//line basepage.qtpl:43
		qw422016.N().S(`
        <script type="text/javascript" src="http://localhost:35729/livereload.js"></script>
    `)
		//line basepage.qtpl:45
	}
	//line basepage.qtpl:45
	qw422016.N().S(`
</body>
</html>
`)
//line basepage.qtpl:48
}

//line basepage.qtpl:48
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	//line basepage.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line basepage.qtpl:48
	StreamPageTemplate(qw422016, p)
	//line basepage.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line basepage.qtpl:48
}

//line basepage.qtpl:48
func PageTemplate(p Page) string {
	//line basepage.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line basepage.qtpl:48
	WritePageTemplate(qb422016, p)
	//line basepage.qtpl:48
	qs422016 := string(qb422016.B)
	//line basepage.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line basepage.qtpl:48
	return qs422016
//line basepage.qtpl:48
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line basepage.qtpl:52
type BasePage struct{}

//line basepage.qtpl:53
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line basepage.qtpl:53
	qw422016.N().S(`This is a base title`)
//line basepage.qtpl:53
}

//line basepage.qtpl:53
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line basepage.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line basepage.qtpl:53
	p.StreamTitle(qw422016)
	//line basepage.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line basepage.qtpl:53
}

//line basepage.qtpl:53
func (p *BasePage) Title() string {
	//line basepage.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
	//line basepage.qtpl:53
	p.WriteTitle(qb422016)
	//line basepage.qtpl:53
	qs422016 := string(qb422016.B)
	//line basepage.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
	//line basepage.qtpl:53
	return qs422016
//line basepage.qtpl:53
}

//line basepage.qtpl:54
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
	//line basepage.qtpl:54
	qw422016.N().S(`This is a base body`)
//line basepage.qtpl:54
}

//line basepage.qtpl:54
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
	//line basepage.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line basepage.qtpl:54
	p.StreamBody(qw422016)
	//line basepage.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line basepage.qtpl:54
}

//line basepage.qtpl:54
func (p *BasePage) Body() string {
	//line basepage.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
	//line basepage.qtpl:54
	p.WriteBody(qb422016)
	//line basepage.qtpl:54
	qs422016 := string(qb422016.B)
	//line basepage.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
	//line basepage.qtpl:54
	return qs422016
//line basepage.qtpl:54
}
