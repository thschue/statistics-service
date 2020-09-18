package operations

import (
	"time"
)

// GetStatisticsParams godoc
type GetStatisticsParams struct {
	From time.Time `form:"from" json:"from" time_format:"unix"`
	To   time.Time `form:"to" json:"to" time_format:"unix"`
}

// Statistics godoc
type Statistics struct {
	From     time.Time           `json:"from" bson:"from"`
	To       time.Time           `json:"to" bson:"to"`
	Projects map[string]*Project `json:"projects" bson:"projects"`
}

// Project godoc
type Project struct {
	Name     string              `json:"name" bson:"name"`
	Services map[string]*Service `json:"services" bson:"services"`
}

// Service godoc
type Service struct {
	Name              string         `json:"name" bson:"name"`
	ExecutedSequences int            `json:"executedSequences" bson:"executedSequences"`
	Events            map[string]int `json:"events" bson:"events"`
}

func (s *Statistics) ensureProjectAndServiceExist(projectName string, serviceName string) {
	if s.Projects == nil {
		s.Projects = map[string]*Project{}
	}
	if s.Projects[projectName] == nil {
		s.Projects[projectName] = &Project{
			Name:     projectName,
			Services: map[string]*Service{},
		}
	}
	if s.Projects[projectName].Services[serviceName] == nil {
		s.Projects[projectName].Services[serviceName] = &Service{
			Name:              serviceName,
			ExecutedSequences: 0,
			Events:            map[string]int{},
		}
	}
}

// IncreaseEventTypeCount godoc
func (s *Statistics) IncreaseEventTypeCount(projectName, serviceName, eventType string, increment int) {
	s.ensureProjectAndServiceExist(projectName, serviceName)
	service := s.Projects[projectName].Services[serviceName]
	service.Events[eventType] = service.Events[eventType] + increment
}

// IncreaseExecutedSequencesCount godoc
func (s *Statistics) IncreaseExecutedSequencesCount(projectName, serviceName string, increment int) {
	s.ensureProjectAndServiceExist(projectName, serviceName)
	service := s.Projects[projectName].Services[serviceName]
	service.ExecutedSequences = service.ExecutedSequences + increment
}
