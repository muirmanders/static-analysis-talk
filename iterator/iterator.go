package iterator

import "fmt"

// ITER_DEF_START OMIT
type Iter interface {
	Next() bool
	Close() error
}

// ITER_DEF_END OMIT

type iterImpl struct{}

func (iterImpl) Next() bool {
	return false
}

func (iterImpl) Close() error {
	return nil
}

type User struct {
	Name string
}

func GetUsers(*User) Iter {
	return iterImpl{}
}

func foo() {
// ITER_USE_START OMIT
var user User

iter := GetUsers(&user)
for iter.Next() {
	fmt.Println("user", user.Name)
}

if err := iter.Close(); err != nil {
	fmt.Println("error fetching users:", err)
}
// ITER_USE_END OMIT
}

func bar() []string {
// ITER_MISUSE_START OMIT
var (
	user User
	userNames []string
)

iter := GetUsers(&user)
for iter.Next() {
	userNames = append(userNames, user.Name)
}

// Oops! forgot iter.Close()

return userNames
// ITER_MISUSE_END OMIT
}
