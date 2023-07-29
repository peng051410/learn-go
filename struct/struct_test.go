package struct_test

import (
	"testing"
	"unsafe"
	"time"
)

func TestStructDeclare(t *testing.T) {
	type User struct {
		Name    string
		Age     int
		Address string
	}

	user := User{}
	t.Log("age is: ", user.Age)
}

func TestStructDeclareEmpty(t *testing.T) {
	type User struct {}	
	var user User
	t.Log("size of struct is: ", unsafe.Sizeof(user))
}

func TestStructWithOtherStruct(t *testing.T) {
	type Address struct {
		Name string
		Code string
	}

	type User struct {
		Name    string
		Age     int
		Address Address
	}

	user := User{}
	t.Log("code is: ", user.Address.Name)
}

func TestStructWithImportOtherStruct(t *testing.T) {
	type Address struct {
		Name string
		Code int 
	}

	type User struct {
		Name    string
		Age     int
		Address 
	}	

	user := User{}
	t.Log("code is: ", user.Address.Code)
	t.Log("code is: ", user.Code)
}

func TestStructWithPointer(t *testing.T) {
	type User struct {
		up *User
		us []User
		um map[string]User
	}

	user := User{}
	t.Log("user is: ", user)
}

func TestStructInit(t *testing.T) {

	type User struct {
		Name    string
		Age     int
	}

	user := User{
		Name: "John",
		Age: 1,
	}
	t.Log(user)

	var badUser User
	badUser.Name = "Tom"
	badUser.Age = 0
	t.Log(badUser)
}

func TestStructZeroInit(t *testing.T) {
	type User struct {
		Name string
		Age int
	}

	// Complie success, but no meaning
	var user User
	t.Log(user)
}

func TestStructDeclareBadCase(t *testing.T) {
	type User struct {
		Name string
		// add a new name, construct init complie error
		// Nick string 
		Age int
	}

	user := User{"TOM", 11}
	t.Log(user)
}

