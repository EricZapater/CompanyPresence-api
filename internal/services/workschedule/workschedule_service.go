package workscheduleservice

import (
	"companypresence-api/internal/models"
	"companypresence-api/internal/repositories"
	"context"
)

type WorkscheduleService struct {
	workScheduleRepository repositories.WorkScheduleRepository
}

func NewWorkScheduleService(workScheduleRepository repositories.WorkScheduleRepository) *WorkscheduleService{
	return &WorkscheduleService{
		workScheduleRepository: workScheduleRepository,
	}
}

func (s *WorkscheduleService)CreateWorkSchedule(ctx context.Context, workschedule models.WorkSchedule)error{
	return s.workScheduleRepository.CreateWorkSchedule(ctx, workschedule)
}

func (s *WorkscheduleService)GetWorkScheduleById(ctx context.Context, id string)(workschedule models.WorkSchedule, err error){
	return s.workScheduleRepository.GetWorkScheduleById(ctx, id)
}

func (s *WorkscheduleService)GetWorkSchedulesByUserId(ctx context.Context, userid string)(workschedule models.WorkSchedule, err error){
	return s.workScheduleRepository.GetWorkSchedulesByUserId(ctx, userid)
}

func (s *WorkscheduleService)GetWorkSchedules(ctx context.Context, id string)(workschedules []models.WorkSchedule, err error){
	return s.workScheduleRepository.GetWorkSchedules(ctx, id)
}

func (s *WorkscheduleService)UpdateWorkSchedule(ctx context.Context, workschedule models.WorkSchedule)error{
	return s.workScheduleRepository.UpdateWorkSchedule(ctx, workschedule)
}

func (s *WorkscheduleService)DeleteWorkSchedule(ctx context.Context, id string)error{
	return s.workScheduleRepository.DeleteWorkSchedule(ctx, id)
}