// Code generated by templ@(devel) DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "strings"

import "strconv"

func border() templ.CSSClass {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`border:1px solid #eeeeee;`)
	templCSSBuilder.WriteString(`border-radius:4px;`)
	templCSSBuilder.WriteString(`margin:10px;`)
	templCSSBuilder.WriteString(`padding-top:30px;`)
	templCSSBuilder.WriteString(`padding-bottom:30px;`)
	templCSSID := templ.CSSID(`border`, templCSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}

func test(other string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(other))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func counts(global, session int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form id=\"countsForm\" action=\"/\" method=\"POST\" hx-post=\"/\" hx-select=\"#countsForm\" hx-swap=\"outerHTML\"><div class=\"columns\">")
		if err != nil {
			return err
		}
		var var_3 = []any{"column", "has-text-centered", "is-primary", border}
		err = templ.RenderCSSItems(ctx, templBuffer, var_3...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_3).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1 class=\"title is-size-1 has-text-centered\">")
		if err != nil {
			return err
		}
		var var_4 string = strconv.Itoa(global)
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"subtitle has-text-centered\">")
		if err != nil {
			return err
		}
		var_5 := `Global`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div><button class=\"button is-primary\" type=\"submit\" name=\"global\" value=\"global\">")
		if err != nil {
			return err
		}
		var_6 := `+1`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></div>")
		if err != nil {
			return err
		}
		var var_7 = []any{"column", "has-text-centered", border}
		err = templ.RenderCSSItems(ctx, templBuffer, var_7...)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(templ.CSSClasses(var_7).String()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1 class=\"title is-size-1 has-text-centered\">")
		if err != nil {
			return err
		}
		var var_8 string = strconv.Itoa(session)
		_, err = templBuffer.WriteString(templ.EscapeString(var_8))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"subtitle has-text-centered\">")
		if err != nil {
			return err
		}
		var_9 := `Session`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><div><button class=\"button is-secondary\" type=\"submit\" name=\"session\" value=\"session\">")
		if err != nil {
			return err
		}
		var_10 := `+1`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div></div></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Page(global, session int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if err != nil {
			return err
		}
		var_12 := `Counts`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"stylesheet\" href=\"/assets/css/bulma.min.css\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/assets/favicon/apple-touch-icon.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/assets/favicon/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/assets/favicon/favicon-16x16.png\"><link rel=\"manifest\" href=\"/assets/favicon/site.webmanifest\"><script src=\"/assets/js/htmx.min.js\">")
		if err != nil {
			return err
		}
		var_13 := ``
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body class=\"bg-gray-100\"><header class=\"hero is-primary\"><div class=\"hero-body\"><div class=\"container\"><h1 class=\"title\">")
		if err != nil {
			return err
		}
		var_14 := `Counts`
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1></div></div></header><section class=\"section\"><div class=\"container\"><div class=\"columns is-centered\"><div class=\"column is-half\">")
		if err != nil {
			return err
		}
		err = counts(global, session).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div></div></section></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
