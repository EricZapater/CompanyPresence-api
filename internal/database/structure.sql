--v0.0.1
CREATE TABLE users(
	ID uuid not null primary key,
	Name varchar(150) not null,
	Surname varchar(250) not null,
	Email varchar(250) not null,
	Password varchar(250) not null,
	IpAddress varchar (50) not null,
	IsAdmin bool,
	Active bool
);
create unique index uq_users_email on users(email);
create index idx_users_email_password on users(email, password);

CREATE TABLE workschedules(
	ID uuid not null primary key,
	UserID uuid not null references users(id),
	NormalWorkingHours decimal(4,2),
	NormalStartTime time,
	NormalNoonRest decimal(4,2),
	FridayWorkingHours decimal(4,2),
	FridayStartTime time
);
create index idx_workschedules_userid on workschedules(userid);

CREATE TABLE timetrackings (
	ID uuid not null primary key,
	UserID uuid not null references users(id),
	WorkingDate date,
	ClockIn time,
	ClockOut time,
	IpAddress varchar(50) not null
);
create index idx_timetrackings_userid on timetrackings(userid);