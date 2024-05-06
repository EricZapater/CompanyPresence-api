package repositories

import (
	"companypresence-api/internal/database"
	"companypresence-api/internal/models"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TimeTrackingRepository struct {
	db *sql.DB
}

func NewTimeTrackingRepository(db *sql.DB) *TimeTrackingRepository {
	return &TimeTrackingRepository{db: db}
}

func (r *TimeTrackingRepository) CreateTimeTracking(ctx context.Context, timetracking models.TimeTracking)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	userRepository := NewUserRepository(r.db)
	user,err := userRepository.GetUserById(ctx, timetracking.UserID)
	if err != nil{
		return err
	}
	sql := `INSERT INTO public.timetracking(ID, UserID, WorkingDate, ClockIn, ClockOut, IpAddress)`
	_, err = db.ExecContext(ctx, sql, ID, timetracking.UserID, timetracking.WorkingDate.Format("2006-01-02"), timetracking.ClockIn.Format("15:04:05")  , timetracking.ClockOut.Format("15:04:05")  , user.IpAddress)
	if err != nil{
		return err
	}
	return err
}

func (r *TimeTrackingRepository)GetTimeTrackings(ctx context.Context)(timetrackings []models.TimeTracking, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return timetrackings,err
	}
	defer db.Close();
	var timetracking models.TimeTracking
	sql := `SELECT ID, UserID, WorkingDate, ClockIn, ClockOut, IpAddress FROM public.timetracking WHERE ClockOut IS NOT NULL`
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return timetrackings, err
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&timetracking.ID, &timetracking.UserID, &timetracking.WorkingDate, &timetracking.ClockIn, &timetracking.ClockOut, &timetracking.IpAddress)
		if err != nil {
			return nil, err
		}
		timetrackings = append(timetrackings, timetracking)
	}
	return timetrackings, nil
}
func (r *TimeTrackingRepository)GetTimeTrackingsByUserId(ctx context.Context, userid string)(timetrackings []models.TimeTracking, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return timetrackings,err
	}
	defer db.Close();
	var timetracking models.TimeTracking
	sql := `SELECT ID, UserID, WorkingDate, ClockIn, ClockOut, IpAddress FROM public.timetracking WHERE ClockOut IS NOT NULL AND userid = $1`
	rows, err := db.QueryContext(ctx, sql, userid)
	if err != nil {
		return timetrackings, err
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&timetracking.ID, &timetracking.UserID, &timetracking.WorkingDate, &timetracking.ClockIn, &timetracking.ClockOut, &timetracking.IpAddress)
		if err != nil {
			return nil, err
		}
		timetrackings = append(timetrackings, timetracking)
	}
	return timetrackings, nil
}
func (r *TimeTrackingRepository)GetOpenTimeTrackingByUserId(ctx context.Context, userid string)(timetracking models.TimeTracking, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return timetracking,err
	}
	defer db.Close();
	sql := `SELECT ID, UserID, WorkingDate, ClockIn, ClockOut, IpAddress FROM public.timetracking WHERE ClockOut IS NULL AND userId = $1`
	row := db.QueryRowContext(ctx, sql, userid)
	err = row.Scan(&timetracking.ID, &timetracking.UserID, &timetracking.WorkingDate, &timetracking.ClockIn, &timetracking.ClockOut, &timetracking.IpAddress)
	if err != nil {
		return timetracking, err
	}
	return timetracking, nil
}
func (r *TimeTrackingRepository)UpdateOpenTimeTtrackingOfUserId(ctx context.Context, userid string, closeTime time.Time)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	timeTracking, err := r.GetOpenTimeTrackingByUserId(ctx, userid)
	if err != nil {
		return err
	}
	sql := `UPDATE public.timetrackings SET ClockOut = $1 WHERE id = $2`
	_, err = db.ExecContext(ctx, sql, closeTime, timeTracking.ID)
	if err != nil {
		return err
	}
	return nil
}