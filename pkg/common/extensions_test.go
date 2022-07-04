package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScheduledTournament_HasCourse(t *testing.T) {
	tournament := ScheduledTournament{
		Courses: []Course{{
			CourseNumber: "001",
		}},
	}

	assert.True(t, tournament.HasCourse("001"))
}

func TestScheduledTournament_NotHasCourse(t *testing.T) {
	tournament := ScheduledTournament{
		Courses: []Course{{
			CourseNumber: "001",
		}},
	}

	assert.False(t, tournament.HasCourse("000"))
}

func TestScheduledTournament_HostCourse_OneCourse(t *testing.T) {
	tournament := ScheduledTournament{
		Courses: []Course{{
			CourseNumber: "001",
		}},
	}

	expected := Course{
		CourseNumber: "001",
	}

	res, _ := tournament.GetHostCourse()
	// Show that a course in a tournament with one course is the host.
	assert.Equal(t, expected, res)
}

func TestScheduledTournament_HostCourse_TwoCourses(t *testing.T) {
	tournament := ScheduledTournament{
		Courses: []Course{{
			CourseNumber: "001",
		}, {
			CourseNumber: "002",
			IsHost:       true,
		}},
	}

	expected := Course{
		CourseNumber: "002",
		IsHost:       true,
	}

	res, _ := tournament.GetHostCourse()
	// Show that the boolean flag indicates the host.
	assert.Equal(t, expected, res)
}
