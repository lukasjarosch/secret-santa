package handler

import (
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lukasjarosch/genki/logger"
	"github.com/lukasjarosch/genki/types/nullable"

	"github.com/lukasjarosch/secret-santa/internal/models"
)

const GroupNavigationName = "groups"
const MyGroupNavigationName = "mygroup"

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
	group := groups[0]

	tplData := struct {
		Group     models.Group
		NavActive string
	}{
		Group:     group,
		NavActive: MyGroupNavigationName,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/showGroup.html.tmpl")
	err := tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	tplData := struct {
		NavActive string
	}{
		NavActive: MyGroupNavigationName,
	}

	tmpl, _ := template.ParseFiles("templates/layout.html.tmpl", "templates/groups/getGroup.html.tmpl")
	err := tmpl.Execute(w, tplData)
	if err != nil {
		logger.Warnf("failed to execute template: %s", err)
	}
}

var groups = []models.Group{
	{
		ID:          12,
		GroupID:     uuid.New(),
		Name:        "Testgruppe",
		Status:      models.AllWishesCollectedStatus,
		Description: nullable.NewString("meine tolle gruppe"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          13,
		GroupID:     uuid.New(),
		Name:        "Spezi Gruppe",
		Status:      models.WishListEmailSentStatus,
		Description: nullable.NewString("Lorem ipsum dolor sit amet...whatever"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
}
