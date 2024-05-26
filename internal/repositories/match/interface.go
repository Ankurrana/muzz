package match

type MatcherRepository interface {
	Swipe(fromUserId, toUserId int) (bool, int, error)

	Matches(userId int) ([]int, error)

	Discover(userId int) ([]int, error)

	AddUser(userId int) error
}
