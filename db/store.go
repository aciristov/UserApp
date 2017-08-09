package db

type DBUser struct{
        ID int
        Username string
        Password string
        Email string
}

type UserRepository map[int]*DBUser

var UserRepo UserRepository = make(UserRepository)

type IDKey int

var IDSeq IDKey = 1

func (seq IDKey) GetNext() IDKey{
        seq = seq + 1
        return seq
}

func (repo UserRepository) GetUserById(ID int) *DBUser{
        return repo[ID]
}

func (repo UserRepository) CreateUser(username string, password string, email string) *DBUser{
        user := &DBUser{
                ID:        (int(IDSeq.GetNext())),
                Username:  username,
                Password:  password,
                Email:     email,
        }
        repo[user.ID] = user
        return user
}

func init(){
        UserRepo[1] = &DBUser{
                ID: 100,
                Username: "testuserrrrr",
                Password: "pass",
                Email: "email@email.com",
        }
        UserRepo[2] = &DBUser{
                ID: 123,
                Username: "secondtestuser",
                Password: "passwoord",
                Email: "email@example.com",
        }
        UserRepo[3] = &DBUser{
                ID: 123,
                Username: "thtestuser",
                Password: "passwoord",
                Email: "email@last.com",
        }

}
