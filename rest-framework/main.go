package main

import "net/http"

// TODO: Define a Router struct that stores handlers keyed by method+path.
type Router struct{
	routes 	map[string]map[string]http.HandlerFunc
}
// TODO: Implement NewRouter() *Router.
func NewRouter() *Router{
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

// TODO: Implement Handle(method string, path string, handler http.HandlerFunc).
func (r *Router) Handle(method string, path string , handler http.HandlerFunc) {
	if r.routes[path] == nil{
		r.routes[path]= make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}
// TODO: Implement ServeHTTP(w http.ResponseWriter, r *http.Request):
//   - Look up handler by method + path
//   - If path exists but method does not match, return 405
//   - If path not found, return 404
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request){
	methods,ok := r.routes[req.URL.Path]
	if !ok{
		http.Error(w, "Not Found",http.StatusNotFound)
		return
	}
	handler,ok := methods[req.Method]
	if !ok{
		http.Error(w, "method not allowed",http.StatusMethodNotAllowed)
		return
	}
	handler(w,req)
}

func main() {}
