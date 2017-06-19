package zctr

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var GLogger *log.Logger
var GSessions = sessions.NewCookieStore(securecookie.GenerateRandomKey(16))
var GCustomVars map[string]interface{}

type HTTPSystemConf struct {
	Port    string
	Access  string
	Error   string
	TLSCert string
	TLSKey  string
}
type HTTPConf struct {
	System HTTPSystemConf
	Custom map[string]interface{}
}

type ZController struct { //{{{
}

func (c ZController) GetLogger() *log.Logger {
	return GLogger
}
func (c ZController) GetSessions() *sessions.CookieStore {
	return GSessions
}
func (c ZController) GetVar(key string) interface{} {
	return GCustomVars[key]
}
func (c ZController) JavaScript(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	http.ServeFile(w, r, "js/"+vars["name"])
}
func (c ZController) Vars(r *http.Request) map[string]string { //{{{
	return mux.Vars(r)
} //}}}

//}}}

type ZRouter struct { //{{{
	alog     io.Writer
	elog     io.Writer
	router   *mux.Router
	port     string
	sys_conf HTTPSystemConf
}                                       //}}}
func NewZRouter(conf string) *ZRouter { //{{{
	conf_content, err := ioutil.ReadFile(conf)
	if err != nil {
		panic(err)
		return nil
	}
	var http_config HTTPConf
	json.Unmarshal([]byte(string(conf_content)), &http_config)

	r := new(ZRouter)
	r.sys_conf = http_config.System
	GCustomVars = http_config.Custom
	r.defaultSetting()
	r.router = mux.NewRouter()
	return r
}                                    //}}}                                                                                      //}}}
func (r *ZRouter) defaultSetting() { //{{{
	if r.sys_conf.Port == "" {
		r.port = "80"
	} else {
		r.port = r.sys_conf.Port
	}
	if r.sys_conf.Access == "" {
		r.alog = os.Stdout
	} else {
		r.alog, _ = os.OpenFile(r.sys_conf.Access, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	if r.sys_conf.Error == "" {
		r.elog = os.Stdout
	} else {
		r.elog, _ = os.OpenFile(r.sys_conf.Error, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	GLogger = log.New(r.elog, "Logger:", log.Ldate|log.Ltime|log.Lshortfile)

}                                                                                      //}}}
func (r ZRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) { //{{{
	r.router.HandleFunc(path, f)
}                        //}}}
func (r ZRouter) Run() { //{{{
	if r.alog == nil {
		r.alog = os.Stdout
	}
	loggedRouter := handlers.CombinedLoggingHandler(r.alog, r.router)
	log.Println("init port:" + r.port)
	var err error
	if r.sys_conf.TLSCert != "" && r.sys_conf.TLSKey != "" {
		err = http.ListenAndServeTLS(":"+r.port, r.sys_conf.TLSCert, r.sys_conf.TLSKey, context.ClearHandler(loggedRouter))
	} else {
		err = http.ListenAndServe(":"+r.port, context.ClearHandler(loggedRouter))
	}

	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
} //}}}
