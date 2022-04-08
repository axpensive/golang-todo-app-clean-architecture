package repository

import (
	"log"
	"time"

	"github.com/axpensive/golang-todo-app-clean-architecture/app/entity"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func CreateUser(name string, email string, password string) (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd, createUUID(), name, email, Encrypt(password), time.Now())
	return err
}

func Getuser(id int) (user entity.User, err error) {
	user = entity.User{}
	cmd := `select id, uuid, name ,email,password,created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func UpdateUser(name string, email string, id int) (err error) {
	cmd := `update users set name = ?, email =? where id = ?`
	_, err = Db.Exec(cmd, name, email, id)
	return err
}

func DeleteUser(id int) (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, id)
	return err
}

func GetUserByEmail(email string) (user entity.User, err error) {
	user = entity.User{}
	cmd := `select id , uuid, name, email, password, created_at from users where email = ?`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func CreateSession(name string, email string, id int) (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, createUUID(), email, id, time.Now())

	if err != nil {
		return session, err
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, id, email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id,created_at from sessions where uuid =?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)

	if err != nil {
		valid = false
		return
	}

	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id,uuid, name, email,created_at from users where id =?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&sess.Email,
		&user.CreatedAt,
	)
	return user, err
}
