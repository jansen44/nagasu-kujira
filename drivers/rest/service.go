package rest

import (
	"github.com/jansen44/nagasu-kujira/drivers"
	"github.com/jansen44/nagasu-kujira/drivers/rest/controllers"
	"github.com/jansen44/nagasu-kujira/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIClient struct {
	port       string
	controller controllers.IRestAPIController
}

func NewRestAPIClient(config *util.Config) drivers.IDrivers {
	return &APIClient{
		port:       "8080",
		controller: controllers.NewRestAPIController(config.ProjectRepository),
	}
}

func (client *APIClient) Start() {
	logrus.Info("## Starting REST API Server...")
	client.registerEndPoints()

	logrus.Info("## Service REST API on http://localhost:" + client.port)
	logrus.Fatal(http.ListenAndServe(":"+client.port, nil))
}

func (client *APIClient) registerEndPoints() {
	logrus.Info("## Registering endpoints on route '/projects'")
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			client.controller.GetProjects(w, r)
		case http.MethodPost:
			client.controller.PostProject(w, r)
		case http.MethodPut:
			client.controller.PutProject(w, r)
		case http.MethodDelete:
			client.controller.DeleteProject(w, r)
		}
	})
}
