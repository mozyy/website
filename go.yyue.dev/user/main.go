package main

import "go.yyue.dev/user/model"

func main() {

	model.RegisterService()
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handleRPC)
	// log.Println("Listen on :8082")
	// http.ListenAndServe(":8082", mux)
	// cmd.Init()
}

// var cors = map[string]bool{"*": true}

// func handleRPC(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		w.Write([]byte("ok,this is the server ..."))
// 		return
// 	}

// 	// 跨域处理
// 	if origin := r.Header.Get("Origin"); cors[origin] {
// 		w.Header().Set("Access-Control-Allow-Origin", origin)
// 	} else if len(origin) > 0 && cors["*"] {
// 		w.Header().Set("Access-Control-Allow-Origin", origin)
// 	}

// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
// 	w.Header().Set("Access-Control-Allow-Credentials", "true")
// 	if r.Method == "OPTIONS" {
// 		return
// 	}

// 	if r.Method != "POST" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

//http://localhost:53548/user.UserService/Register
// func pathToReceiver(p string) (string, string) {
// 	p = path.Clean(p)
// 	p = strings.TrimPrefix(p, "/")
// 	parts := strings.Split(p, "/")

// 	// If we've got two or less parts
// 	// Use first part as service
// 	// Use all parts as method
// 	if len(parts) <= 2 {
// 		service := ns + strings.Join(parts[:len(parts)-1], ".")
// 		method := strings.Title(strings.Join(parts, "."))
// 		return service, method
// 	}

// 	// Treat /v[0-9]+ as versioning where we have 3 parts
// 	// /v1/foo/bar => service: v1.foo method: Foo.bar
// 	if len(parts) == 3 && versionRe.Match([]byte(parts[0])) {
// 		service := ns + strings.Join(parts[:len(parts)-1], ".")
// 		method := strings.Title(strings.Join(parts[len(parts)-2:], "."))
// 		return service, method
// 	}

// 	// Service is everything minus last two parts
// 	// Method is the last two parts
// 	service := ns + strings.Join(parts[:len(parts)-2], ".")
// 	method := strings.Title(strings.Join(parts[len(parts)-2:], "."))
// 	return service, method
// }
