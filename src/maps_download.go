package opennox

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/maps"
)

var mapsend struct {
	srv *maps.Server
}

func init() {
	if os.Getenv("NOX_MAPS_HTTP") != "false" {
		registerOnDataPathSet(func() {
			mapsend.srv = maps.NewServer(datapath.Maps())
			noxServer.HTTP().Handle("/api/v0/maps/", http.HandlerFunc(serveMapHTTP))
		})
	}
}

// serveMapHTTP serves map downloads over HTTP. When the running coop map was
// restored from a world save, joining clients must receive that saved state
// instead of the stock map files.
func serveMapHTTP(w http.ResponseWriter, r *http.Request) {
	if f := coopSavedMapDownload(r); f != "" {
		http.ServeFile(w, r, f)
		return
	}
	mapsend.srv.ServeHTTP(w, r)
}

func coopSavedMapDownload(r *http.Request) string {
	saved := noxServer.mapSend.savedFile
	if saved == "" || r.Method != http.MethodGet {
		return ""
	}
	// /api/v0/maps/<name>/download
	rest := strings.TrimPrefix(r.URL.Path, "/api/v0/maps/")
	name, _, ok := strings.Cut(rest, "/download")
	if !ok || name == "" || strings.ContainsAny(name, "/\\") {
		return ""
	}
	cur := strings.TrimSuffix(strings.ToLower(filepath.Base(noxServer.getServerMap())), maps.Ext)
	if strings.ToLower(name) != cur {
		return ""
	}
	if _, err := os.Stat(saved); err != nil {
		return ""
	}
	return saved
}

func clientGetServerMap() string {
	if gameIsNotMultiplayer {
		return nox_xxx_mapFilenameGetSolo_4DB260()
	}
	return noxServer.nox_server_currentMapGetFilename_409B30()
}
