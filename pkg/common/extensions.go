package common

func (s *ScheduledTournament) GetHostCourse() (Course, bool) {
	if len(s.Courses) == 1 {
		return s.Courses[0], true
	}

	for _, course := range s.Courses {
		if course.IsHost {
			return course, true
		}
	}

	return Course{}, false
}

func (s *ScheduledTournament) HasCourse(courseNumber string) bool {
	for _, course := range s.Courses {
		if course.CourseNumber == courseNumber {
			return true
		}
	}

	return false
}

func (s *PgaTourScheduleV2) GetTourSchedule(tourCode string) (Schedule, bool) {
	for _, schedule := range s.TourSchedules {
		if schedule.TourCode == tourCode {
			return schedule, true
		}
	}

	return Schedule{}, false
}

func (s *PgaTourScheduleV2) GetScheduledTournament(tourCode string, permNum string) (ScheduledTournament, bool) {
	if schedule, found := s.GetTourSchedule(tourCode); found {
		return schedule.GetScheduledTournament(permNum)
	}

	return ScheduledTournament{}, false
}

func (s *Schedule) GetScheduledTournament(permNum string) (ScheduledTournament, bool) {
	for _, tournament := range s.Tournaments {
		if tournament.TournamentId == permNum {
			return tournament, true
		}
	}

	return ScheduledTournament{}, false
}
