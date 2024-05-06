package services

import (
	"companypresence-api/internal/models"
	"companypresence-api/internal/repositories"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type TimeTrackingService struct {
	timeTrackingRepository repositories.TimeTrackingRepository
	userRepository repositories.UserRepository
	workScheduleRepository repositories.WorkScheduleRepository
}

func NewTimeTrackingService(timeTrackingRepository repositories.TimeTrackingRepository, userRepository repositories.UserRepository, workScheduleRepository repositories.WorkScheduleRepository) *TimeTrackingService{
	return &TimeTrackingService{timeTrackingRepository: timeTrackingRepository, 
								userRepository: userRepository,
								workScheduleRepository: workScheduleRepository}
}

func (s *TimeTrackingService)Scheduler(){
	ctx := context.Background()
	location, _ := time.LoadLocation("Europe/Madrid")
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	go func() {
		for t:=range ticker.C {
			currentTime := t.In(location)
			isFriday := currentTime.Weekday() == time.Friday
			if currentTime.Hour() == 23 && currentTime.Minute() == 30 {				
				users, err := s.userRepository.GetActiveUsers(ctx)
				if err != nil {
					fmt.Errorf("error %s", err)
					continue
				}
				for _, user := range users {
					s.SetTimetracking(ctx, isFriday, user.ID)
				}
			}
		}
	}()	
}

func (s *TimeTrackingService) SetTimetracking(ctx context.Context, isFriday bool, userid string) error {
	location, _ := time.LoadLocation("Europe/Madrid")
	workSchedule, err := s.workScheduleRepository.GetWorkSchedulesByUserId(ctx, userid)
	if err != nil {
		return err
	}
	user, err := s.userRepository.GetUserById(ctx, userid)
	if err != nil {
		return err
	}
	iniTime := getRandomTime(workSchedule.NormalStartTime.Add(-5 * time.Minute), workSchedule.NormalStartTime.Add(10 * time.Minute))
	endTime := getRandomTime(iniTime.Add(5 * time.Hour).Add(45 * time.Minute), iniTime.Add(6 * time.Hour).Add(15 *time.Minute))
	timeTracking := models.TimeTracking{
		ID: "",
		UserID: userid,
		WorkingDate: time.Now().In(location),
		ClockIn: iniTime,
		ClockOut: endTime,
		IpAddress: user.IpAddress,
	}
	err = s.timeTrackingRepository.CreateTimeTracking(ctx, timeTracking)
	if err != nil {
		return err
	}
	return nil
}

func getRandomTime(minTime time.Time, maxTime time.Time)(rTime time.Time){
	//rand.Seed(time.Now().UnixNano())
	difference := maxTime.Sub(minTime).Seconds()
	randomSecs := rand.Float64() * difference
	return  minTime.Add((time.Duration(randomSecs)* time.Second))
}