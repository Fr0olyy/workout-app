package services

import "traning/internal/repository"

type UserService struct {
	repo repository.Profile
}

func NewUserService(repo repository.Profile) *UserService {
	return &UserService{repo: repo}
}

type UserProfileResponse struct {
	ID         uint     `json:"id"`
	Email      string   `json:"email"`
	Name       string   `json:"name"`
	Images     []string `json:"images"`
	Statistics []struct {
		Label string      `json:"label"`
		Value interface{} `json:"value"`
	} `json:"statistics"`
}

func (s *UserService) GetUserProfile(userID uint) (*UserProfileResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	exerciseCount, _ := s.repo.CountCompletedExerciseLogs(userID)
	totalWeight, _ := s.repo.SumCompletedWeights(userID)
	workoutCount, _ := s.repo.CountCompletedWorkouts(userID)

	stats := []struct {
		Label string      `json:"label"`
		Value interface{} `json:"value"`
	}{
		{"Minutes", int64(float64(exerciseCount) * 2.3)},
		{"Workouts", workoutCount},
		{"Kgs", totalWeight},
	}

	return &UserProfileResponse{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Images:     user.Images,
		Statistics: stats,
	}, nil
}
