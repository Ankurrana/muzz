package match

type MatchInMemoryDB struct {
	UserMatchMap map[int]map[int]bool
	MatchesMap   map[int][2]int
	UsersMap     map[int]bool
}

var instance *MatchInMemoryDB

func Instance() *MatchInMemoryDB {
	if instance == nil {
		instance = &MatchInMemoryDB{
			make(map[int]map[int]bool),
			make(map[int][2]int),
			map[int]bool{},
		}
	}
	return instance
}

// Swipe(fromUserId, toUserId int, ts time.Time) (bool, error)
// 	// gets "count" number of matches for this user
// 	Matches(userId int, count int, offset int) ([]int, error)

// 	// gets "count" number of next potencial matches for this user
// 	Discover(userId int, count int, offset int) ([]int, error)

// 	DiscoverByAgeAndGender(userId int, minAge, maxAge int, gender int, count int, offset int) ([]int, error)

// 	DiscoverByAgeGenderAndLoc(userId int, minAge, maxAge int, gender int, maxDis int, count int, offset int) ([]int, error)

func (db *MatchInMemoryDB) Swipe(fromUserId, toUserId int) (bool, int, error) {
	if db.UserMatchMap[fromUserId] == nil {
		db.UserMatchMap[fromUserId] = map[int]bool{}
	}
	db.UserMatchMap[fromUserId][toUserId] = true
	if db.UserMatchMap[toUserId][fromUserId] {
		// create a new match
		l := len(db.MatchesMap)
		db.MatchesMap[l+1] = [2]int{toUserId, fromUserId}
		return true, l + 1, nil
	}
	return false, -1, nil
}

func (db *MatchInMemoryDB) Matches(userId int) ([]int, error) {
	res := []int{}
	for u := range db.UserMatchMap[userId] {
		res = append(res, u)
	}
	return res, nil
}

func (db *MatchInMemoryDB) Discover(userId int) ([]int, error) {
	res := []int{}
	for u := range db.UsersMap {
		if !db.UserMatchMap[userId][u] {
			res = append(res, u)
		}
	}
	return res, nil
}

func (db *MatchInMemoryDB) AddUser(userId int) error {
	db.UsersMap[userId] = true
	return nil
}
