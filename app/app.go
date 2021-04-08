package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manojown/restApi/app/handler"
	"github.com/manojown/restApi/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func ConfigAndRun(config *config.Config) {
	app := new(App)
	app.initialize(config)
	app.run(config.ServerHost)

}

func (app *App) initialize(config *config.Config) {
	app.Router = mux.NewRouter()
	// app.UseMiddleware()
	app.DB = config.Connect("test")
	app.setRouter()

}

func (app *App) run(host string) {
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	// go func() {
	log.Fatal(http.ListenAndServe(":3005", app.Router))
	// }()
	log.Printf("Server is listning on http://%s\n", host)
	// sig := <-sigs
	//log.Println("Signal: ", sig)

}

func (app *App) setRouter() {
	app.Router.HandleFunc("/person", app.funcHandler(handler.CreatePerson)).Methods("POST")
	app.Router.HandleFunc("/testing", app.funcHandler(handler.GetTesting)).Methods("GET")
	app.Router.HandleFunc("/person/{id}", app.funcHandler(handler.GetPerson)).Methods("Get")
}

type HandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

func (app *App) funcHandler(handler HandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
