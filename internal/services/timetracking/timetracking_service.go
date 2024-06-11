package timetrackingservice

import (
	"companypresence-api/internal/models"
	"companypresence-api/internal/repositories"
	userservice "companypresence-api/internal/services/users"
	workscheduleservice "companypresence-api/internal/services/workschedule"
	"context"
	"fmt"
	"math/rand"
	"time"
)

type TimeTrackingService struct {
	timeTrackingRepository repositories.TimeTrackingRepository	
	userService userservice.UserService
	workscheduleservice workscheduleservice.WorkscheduleService
	
}

func NewTimeTrackingService(timeTrackingRepository repositories.TimeTrackingRepository, userService userservice.UserService, workscheduleservice workscheduleservice.WorkscheduleService) *TimeTrackingService{
	return &TimeTrackingService{timeTrackingRepository: timeTrackingRepository, 
								userService: userService,
								workscheduleservice: workscheduleservice}
}

func (s *TimeTrackingService)Scheduler(){
	ctx := context.Background()
	location, _ := time.LoadLocation("Europe/Madrid")
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	fmt.Println("Scheduler started")

	for {
		fmt.Println(time.Now())
		select {
		case t := <-ticker.C:
			currentTime := t.In(location)
			fmt.Println(time.Now())
			isFriday := currentTime.Weekday() == time.Friday
			

			users, err := s.userService.GetActiveUsers(ctx)
			if err != nil {
				fmt.Errorf("error %s", err.Error())
				continue
			}

			for _, user := range users {
				s.SetTimetracking(ctx, isFriday, user.ID)
			}
		case <-ctx.Done():
			fmt.Println("Scheduler stopping")
			return 
		}
	}
}

func (s *TimeTrackingService) SetTimetracking(ctx context.Context, isFriday bool, userid string) error {
	location, _ := time.LoadLocation("Europe/Madrid")
	workSchedule, err := s.workscheduleservice.GetWorkSchedulesByUserId(ctx, userid)
	if err != nil {
		return err
	}
	user, err := s.userService.GetUserById(ctx, userid)
	if err != nil {
		return err
	}
	if !isFriday{
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
		afternoonInitime := getRandomTime(endTime.Add(time.Duration(workSchedule.NormalNoonRest) * time.Hour).Add(-10 * time.Minute),endTime.Add(time.Duration(workSchedule.NormalNoonRest) * time.Hour).Add(10 * time.Minute))
		elapsedTime := endTime.Sub(iniTime).Minutes()
		restOfDay := 510 - elapsedTime
		afternoonEndTime := getRandomTime(afternoonInitime.Add(time.Duration(restOfDay)*time.Minute).Add(-5*time.Minute), afternoonInitime.Add(time.Duration(restOfDay)*time.Minute).Add(5*time.Minute))
		secondTimeTracking := models.TimeTracking{
			ID:"",
			UserID: userid,
			WorkingDate: time.Now().In(location),
			ClockIn: afternoonInitime,
			ClockOut: afternoonEndTime,
			IpAddress: user.IpAddress,
		}
		err = s.timeTrackingRepository.CreateTimeTracking(ctx, secondTimeTracking)
		if err != nil {
			return err
		}
	}else{
		iniTime := getRandomTime(workSchedule.NormalStartTime.Add(-5 * time.Minute), workSchedule.NormalStartTime.Add(10 * time.Minute))
		endTime := getRandomTime(iniTime.Add(6 * time.Hour).Add(45 * time.Minute), iniTime.Add(7 * time.Hour).Add(15 *time.Minute))
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
	}
	
	return nil
}

func getRandomTime(minTime time.Time, maxTime time.Time)(rTime time.Time){
	//rand.Seed(time.Now().UnixNano())
	difference := maxTime.Sub(minTime).Seconds()
	randomSecs := rand.Float64() * difference
	return  minTime.Add((time.Duration(randomSecs)* time.Second))
}