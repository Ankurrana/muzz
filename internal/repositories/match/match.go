package match

type MatcherRepository interface {
	Swipe(fromUserId, toUserId int) (bool, int, error)
	// gets "count" number of matches for this user
	Matches(userId int) ([]int, error)

	// gets "count" number of next potencial matches for this user
	Discover(userId int) ([]int, error)

	AddUser(userId int) error
}
