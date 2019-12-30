package handler

import (
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lukasjarosch/genki/logger"

	"github.com/lukasjarosch/secret-santa/internal/models"
)

const GroupNavigationName = "groups"
const MyGroupNavigationName = "findgroup"
const NewGroupNavigationName = "newgroup"

type GroupHandler struct {
}

func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

func (h *GroupHandler) ListAll(w http.ResponseWriter, r *http.Request) {

	tplData := struct {
		Groups    []models.Group
		Members   []models.GroupMember
		NavActive string
	}{
		Groups:    groups,
		Members:   members,
		NavActive: GroupNavigationName,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/listGroups.html.tmpl")
	err := tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

func (h *GroupHandler) ShowGroup(w http.ResponseWriter, r *http.Request) {
	var groupID string
	group := groups[0]

	switch strings.ToLower(r.Method) {
	case "post":
		if err := r.ParseForm(); err != nil {
			logger.Warnf("failed to parse form: %s", err)
			w.WriteHeader(400)
			break
		}
		groupID = r.PostForm.Get("groupID")
		logger.Infof("post data contained groupID: %s", groupID)
		break
	case "get":
		vars := mux.Vars(r)
		if id, ok := vars["groupID"]; ok {
			groupID = id
			logger.Infof("url contains groupID: %s", groupID)
			break
		}
	default:
		w.WriteHeader(400)
		return
	}

	uuidGroupID, err := uuid.Parse(groupID)
	if err != nil {
		logger.Warnf("invalid group id, not a valid uuid: %s", err)
		// TODO: proper error handling...
		w.WriteHeader(400)
		return
	}
	logger.Infof("groupID is a valid uuid")

	group.GroupID = uuidGroupID

	selfLink := "http://" + path.Join(r.Host, "group", group.GroupID.String())

	tplData := struct {
		SelfLink   string
		Group     models.Group
		Members   []models.GroupMember
		NavActive string
	}{
		SelfLink: selfLink,
		Group:     group,
		NavActive: MyGroupNavigationName,
		Members:   members,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/showGroup.html.tmpl")
	err = tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

func (h *GroupHandler) FindGroup(w http.ResponseWriter, r *http.Request) {
	tplData := struct {
		NavActive string
	}{
		NavActive: MyGroupNavigationName,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/findGroup.html.tmpl")
	err := tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	tplData := struct {
		NavActive string
	}{
		NavActive: NewGroupNavigationName,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/newGroup.html.tmpl")
	err := tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

var groups = []models.Group{
	{
		ID:        12,
		GroupID:   uuid.New(),
		Name:      "Testgruppe",
		Status:    models.AllWishesCollectedStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        13,
		GroupID:   uuid.New(),
		Name:      "Spezi Gruppe",
		Status:    models.WishListEmailSentStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var members = []models.GroupMember{
	{
		ID:        1,
		MemberID:  uuid.New(),
		GroupID:   groups[0].GroupID,
		Name:      "Hans Peter",
		Email:     "hans@peter.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		MemberID:  uuid.New(),
		GroupID:   groups[0].GroupID,
		Name:      "Max Muster",
		Email:     "hans@peter.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
