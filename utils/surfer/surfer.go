package surfer

import (
	"net/http"
	"os"
	"sync"
)

var (
	surf          Surfer
	phantom       Surfer
	once_surf     sync.Once
	once_phantom  sync.Once
	tempJsDir     = "./tmp"
	phantomjsFile = os.Getenv("GOPATH") + `\src\github.com\henrylee2cn\surfer\phantomjs\phantomjs`
)

func Download(req Request) (resp *http.Response, err error) {
	switch req.GetDownloaderID() {
	case SurfID:
		once_surf.Do(func() { surf = New() })
		resp, err = surf.Download(req)
	case PhomtomJsID:
		once_phantom.Do(func() { phantom = NewPhantom(phantomjsFile, tempJsDir) })
		resp, err = phantom.Download(req)
	}
	return
}

//销毁Phantomjs的js临时文件
func DestroyJsFiles() {
	if pt, ok := phantom.(*Phantom); ok {
		pt.DestroyJsFiles()
	}
}

// Downloader represents an core of HTTP web browser for crawler.
type Surfer interface {
	// GET @param url string, header http.Header, cookies []*http.Cookie
	// HEAD @param url string, header http.Header, cookies []*http.Cookie
	// POST PostForm @param url, referer string, values url.Values, header http.Header, cookies []*http.Cookie
	// POST-M PostMultipart @param url, referer string, values url.Values, header http.Header, cookies []*http.Cookie
	Download(Request) (resp *http.Response, err error)
}
