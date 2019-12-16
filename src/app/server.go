package main

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shadow1163/logger"
	"github.com/shadow1163/new-server/src/account"
	"github.com/shadow1163/new-server/src/chatroom"
	"github.com/shadow1163/new-server/src/fileserver"
	"github.com/shadow1163/new-server/src/game"
)

var (
	jsPath    = "public/js/"
	cssPath   = "public/css/"
	log       = logger.NewLogger()
	httpPort  = ":80"
	httpsPort = ":443"
	cert      = "certificate/cert.pem"
	key       = "certificate/key.pem"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	host, _, err := net.SplitHostPort(req.Host)
	if err != nil {
		host = req.Host
	}
	target := "https://" + host + httpsPort + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	jsfs := http.FileServer(http.Dir(jsPath))
	cssfs := http.FileServer(http.Dir(cssPath))
	fileserverfs := http.FileServer(http.Dir(fileserver.UploadPath))

	r.HandleFunc("/", fileserver.FileserverIndex)
	r.HandleFunc("/upload", fileserver.UploadFileHandler)
	r.HandleFunc("/key", game.MiniGame)

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsfs))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssfs))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fileserverfs))

	//session
	r.HandleFunc("/signin", account.Signin)
	r.HandleFunc("/uploadpage", account.UploadPage)
	r.HandleFunc("/refresh", account.Refresh)
	r.HandleFunc("/logout", account.LogoutHandler).Methods("POST")

	//ChatRoom
	r.HandleFunc("/chatroom", chatroom.ChatRoom)
	r.HandleFunc("/chatroom/counter", chatroom.GetChatRoomCounter).Methods("Get")
	r.HandleFunc("/ws", chatroom.HandleWSConnections)

	server := &http.Server{
		Addr:    httpsPort,
		Handler: r,
	}

	go http.ListenAndServe(httpPort, http.HandlerFunc(redirect))
	log.Info("Listening...")
	server.ListenAndServeTLS(cert, key)
}
