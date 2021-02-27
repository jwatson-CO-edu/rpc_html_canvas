// GOARCH=wasm GOOS=js GO111MODULE=auto go build -o lib.wasm canvas_api.go

package rpc_html_canvas

import (
	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

func serve_display( params ...int64 ){
	
}

func set_up_canvas( h int, w int ) ( web.Canvas, canvas.Context2D ) {
	// 1. Fetch refs
	window := web.GetWindow()
	doc    := window.Document()
	cnvs   := doc.CreateCanvas()
	cntxt  := cnvs.Context2D()
	// 2. Set params
	cnvs.SetHeight(h)
	cnvs.SetWidth(w)
	return cnvs, cntxt
}

func main(){
	
}
