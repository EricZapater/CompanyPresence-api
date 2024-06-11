package repositories

import (
	"companypresence-api/internal/database"
	"companypresence-api/internal/models"
	"context"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository{
	return &UserRepository{}
}

func (r *UserRepository)CreateUSer(ctx context.Context, user *models.User) error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	password, bcErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcErr != nil {
		return bcErr
	}
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	sql := `INSERT INTO public.users(ID, name, surname, email, password, ipaddress, isadmin, active)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.ExecContext(ctx, sql, ID, user.Name, user.Surname, user.Email, password, user.IpAddress, user.IsAdmin, user.Active)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository)GetUserById(ctx context.Context, id string)(user models.User, err error){	
	db, err := database.NewDatabase()
	if err != nil {
		return user, err
	}
	defer db.Close();

	sql := `SELECT ID, name, surname, email, password, ipaddress, isadmin, active FROM public.users where ID = $1`

	row := db.QueryRowContext(ctx, sql, id)
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.IpAddress,&user.IsAdmin, &user.Active)
	if err !=nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository)GetUserByMail(ctx context.Context, mail string)(user models.User, err error){	
	db, err := database.NewDatabase()
	if err != nil {
		return user, err
	}
	defer db.Close();

	sql := `SELECT ID, name, surname, email, password, ipaddress, isadmin, active FROM public.users where email = $1`

	row := db.QueryRowContext(ctx, sql, mail)
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.IpAddress, &user.IsAdmin, &user.Active)
	if err !=nil {
		return user, err
	}
	return user, nil
}


func (r *UserRepository)UpdateUser(ctx context.Context, user *models.User)error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	sql := `UPDATE public.users 
			SET Name = $1,
				Surname = $2,
				Email = $3,
				Password = $4,
				IpAddress = $5,
				IsAdmin = $6,
				Active = $7
			WHERE ID = $8`
	_, err = db.ExecContext(ctx, sql, user.Name, user.Surname, user.Email, user.Password, user.IpAddress,user.IsAdmin, user.Active, user.ID)
	return err
}
func (r *UserRepository)DeleteUser(ctx context.Context, id string) error{
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close();
	user, err := r.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	user.Active = false
	err = r.UpdateUser(ctx, &user)
	return err
}

func (r *UserRepository)GetUsers(ctx context.Context, active bool)(users []models.User, err error){
	db, err := database.NewDatabase()
	if err != nil {
		return users, err
	}
	defer db.Close();
	var sql string;
	if active {
		sql = `SELECT ID, name, surname, email, password, ipaddress, isadmin, active FROM public.users WHERE active = true`
	}else{
		sql = `SELECT ID, name, surname, email, password, ipaddress, isadmin, active FROM public.users`
	}
	

	var user models.User

	rows, err := db.QueryContext(ctx, sql)	
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.Password, &user.IpAddress, &user.IsAdmin, &user.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}