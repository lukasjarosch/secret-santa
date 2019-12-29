package main

import (
	"log"
	http2 "net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lukasjarosch/genki"
	"github.com/lukasjarosch/genki/cli"
	"github.com/lukasjarosch/genki/config"
	"github.com/lukasjarosch/genki/logger"
	"github.com/lukasjarosch/genki/server/http"
	"github.com/lukasjarosch/genki/types/nullable"

	"github.com/gorilla/mux"

	"github.com/lukasjarosch/secret-santa/internal/handler"
	"github.com/lukasjarosch/secret-santa/internal/models"
)

const Service = "stringer"

func main() {
	if err := logger.NewLogger(config.GetString(logger.LogLevelConfigKey)); err != nil {
		log.Fatal(err.Error())
	}

	group := models.Group{
		ID:          1,
		GroupID:     uuid.New(),
		Name:        "asdf",
		Description: nullable.String{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	member := models.GroupMember{
		ID:        1,
		MemberID:  uuid.New(),
		GroupID:   group.GroupID,
		Name:      "derp",
		Email:     "asdf@a.b",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	valid, errs := models.ValidateStruct(group)
	if !valid {
		for _, err := range errs {
			logger.Warnf("%s", err)
		}
	}

	valid, errs = models.ValidateStruct(member)
	if !valid {
		for _, err := range errs {
			logger.Warnf("%s", err)
		}
	}

	router := mux.NewRouter()

	groupHandler := handler.NewGroupHandler()

	router.HandleFunc("/group/", groupHandler.GetGroup)
	router.HandleFunc("/group", groupHandler.GetGroup)
	router.HandleFunc("/group/{groupID}", groupHandler.ShowGroup)
	router.HandleFunc("/admin/groups/", groupHandler.ListAll)
	router.HandleFunc("/admin/groups", groupHandler.ListAll)
	router.PathPrefix("/").Handler(http2.FileServer(http2.Dir("./static/")))

	httpServer := http.NewServer()
	httpServer.Handle("/", router)

	app := genki.NewApplication()
	app.AddServer(httpServer)

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

// init takes care of setting up the CLI flags, parsing and ultimately binding
// them to the configuration. After they are bound, they are globally accessible via the config package.
func init() {
	flags := cli.NewFlagSet(Service)
	flags.Add(logger.Flags, http.Flags)
	flags.Parse()
	config.BindFlagSet(flags.Set())
}
