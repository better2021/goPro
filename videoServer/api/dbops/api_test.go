package dbops

import "testing"

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T) {
	err:= AddUserCredential("admin","123")
	if err!=nil{
		t.Errorf("error of adduser:%v",err)
	}
}

func testGetUser(t *testing.T)  {
	 pwd,err := GetUserCredential("admin")
	if pwd!="123"||err!=nil {
		t.Errorf("error of getUser:%v",err)
	}
}

func testDeleteUser(t *testing.T) {
	err:=DeleteUser("admin","123")
	if err!=nil{
		t.Errorf("error of deleteUser:%v",err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd,err:= GetUserCredential("admin")
	if err!=nil{
		t.Errorf("error of RegetUser:%v",err)
	}

	if pwd!=""{
		t.Errorf("delete user test failed")
	}
}