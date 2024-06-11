package repositories

import (
	"companypresence-api/internal/database"
	"companypresence-api/internal/models"
	"context"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type WorkScheduleRepository struct {
	
}

func NewWorkScheduleRepository() *WorkScheduleRepository {
	return &WorkScheduleRepository{}
}

func (r *WorkScheduleRepository)CreateWorkSchedule(ctx context.Context, workschedule models.WorkSchedule)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	sql := `INSERT INTO public.workschedules(ID, UserID, NormalWorkingHours, NormalStartTime, NormalNoonRest, FridayWorkingHours, FridayStartTime)
			VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.ExecContext(ctx, sql, ID, workschedule.UserID, workschedule.NormalWorkingHours, workschedule.NormalStartTime, workschedule.NormalNoonRest, workschedule.FridayWorkingHours, workschedule.FridayStartTime)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkScheduleRepository)GetWorkScheduleById(ctx context.Context, id string)(workschedule models.WorkSchedule, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return workschedule, err
	}
	defer db.Close();

	sql := `SELECT ID, UserID, NormalWorkingHours, NormalStartTime, NormalNoonRest, FridayWorkingHours, FridayStartTime 
	FROM public.workschedules WHERE id = $1`

	row := db.QueryRowContext(ctx, sql, id)
	err = row.Scan(&workschedule.ID, &workschedule.UserID, &workschedule.NormalWorkingHours, &workschedule.NormalStartTime, &workschedule.NormalNoonRest, &workschedule.FridayWorkingHours, &workschedule.FridayStartTime)
	if err !=nil {
		return workschedule, err
	}	
	return workschedule, nil
}

func (r *WorkScheduleRepository)GetWorkSchedulesByUserId(ctx context.Context, userid string)(workschedule models.WorkSchedule, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return workschedule, err
	}
	defer db.Close();

	sql := `SELECT ID, UserID, NormalWorkingHours, NormalStartTime, NormalNoonRest, FridayWorkingHours, FridayStartTime 
	FROM public.workschedules WHERE userid = $1`

	row := db.QueryRowContext(ctx, sql, userid)
	err = row.Scan(&workschedule.ID, &workschedule.UserID, &workschedule.NormalWorkingHours, &workschedule.NormalStartTime, &workschedule.NormalNoonRest, &workschedule.FridayWorkingHours, &workschedule.FridayStartTime)
	if err !=nil {
		return workschedule, err
	}	
	return workschedule, nil
}

func (r *WorkScheduleRepository)GetWorkSchedules(ctx context.Context, id string)(workschedules []models.WorkSchedule, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return workschedules, err
	}
	defer db.Close();

	sql := `SELECT ID, UserID, NormalWorkingHours, NormalStartTime, NormalNoonRest, FridayWorkingHours, FridayStartTime 
			FROM public.workschedules`
	

	var workschedule models.WorkSchedule

	rows, err := db.QueryContext(ctx, sql)	
	if err != nil{
		return workschedules, err 
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&workschedule.ID, &workschedule.UserID, &workschedule.NormalWorkingHours, &workschedule.NormalStartTime, &workschedule.NormalNoonRest, &workschedule.FridayWorkingHours, &workschedule.FridayStartTime)
		if err != nil{
			return nil, err
		}
		workschedules = append(workschedules, workschedule)
	}
	return workschedules, err
}

func (r *WorkScheduleRepository)UpdateWorkSchedule(ctx context.Context, workschedule models.WorkSchedule)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	sql := `UPDATE public.workschedules
			SET UserID = $1,
				NormalWorkingHours = $2,
				NormalStartTime = $3,
				NormalNoonRest = $4,
				FridayWorkingHours = $5,
				FridayStartTime = $6
			WHERE id = $7`
	_, err = db.ExecContext(ctx, sql, workschedule.UserID, workschedule.NormalWorkingHours, workschedule.NormalStartTime, workschedule.NormalNoonRest, workschedule.FridayWorkingHours, workschedule.FridayStartTime, workschedule.ID)
	return err
}

func (r *WorkScheduleRepository)DeleteWorkSchedule(ctx context.Context, id string)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	sql := `DELETE FROM public.workschedules WHERE id = $1`
	_, err = db.ExecContext(ctx, sql, id)
	return err
}