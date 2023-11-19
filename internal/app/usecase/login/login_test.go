package login

import (
	"grapefruit/internal/app/service"
	"grapefruit/internal/app/service/candidate"
	"grapefruit/internal/app/service/employer"
	"testing"
)

func TestUser(t *testing.T) {
	var a int
	switch a {
	case 0:
		Use(&candidate.Candidate{})
	case 1:
		Use(&employer.Employer{})

	}
}

func Use(user service.IUser) {

}
