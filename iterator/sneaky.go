package iterator

func bar() []string {
// START OMIT
var (
	user User
	userNames []string
)

iter := GetUsers(&user)
for iter.Next() {
	userNames = append(userNames, user.Name)
}

if false {
	iter.Close()
}

return userNames
// END OMIT
}
