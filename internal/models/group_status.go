package models

type GroupStatus int

const (
	UndefinedStatus GroupStatus = 0
	CreatedStatus   GroupStatus = 1
	InvitedStatus   GroupStatus = 2
	MappedStatus    GroupStatus = 3
	FinishedStatus  GroupStatus = 4
)

var GroupStatusNameMap = map[int]string{
	0: "undefined_status",
	1: "created_status",
	2: "invited_status",
	3: "mapped_status",
	4: "finished_status",
}

var GroupStatusValueMap = map[string]int{
	"undefined_status": 0,
	"created_status":   1,
	"invited_status":   2,
	"mapped_status":    3,
	"finished_status":  4,
}
