package match

import (
	"github.com/ankur-toko/muzz/internal/repositories/sql"
	"gorm.io/gorm"
)

type MatchPostgresRepo struct {
	Conn *gorm.DB
}

// Swipe u
func (mDB *MatchPostgresRepo) Swipe(fromUserId, toUserId int) (bool, int, error) {
	swipe := sql.Swipe{
		FromUser: fromUserId,
		ToUser:   toUserId,
		Swipe:    true,
	}

	res := mDB.Conn.Create(&swipe)
	if res.Error != nil {
		return false, -1, res.Error
	}

	var swipes []sql.Swipe

	mDB.Conn.
		Where("from_user = ?", toUserId).
		Where("to_user = ?", fromUserId).
		Where("swipe = ?", true).
		Find(&swipes)

		//
	mDB.Conn.Exec("UPDATE user_summaries SET swipe_count = swipe_count + 1 WHERE user_id = ?", fromUserId)

	// It user is already swiped by the other user, we create a new match
	if len(swipes) > 0 {
		// Create a match as well
		match := sql.Match{UserA: toUserId, UserB: fromUserId}
		res := mDB.Conn.Create(&match)
		if res.Error != nil {
			return false, -1, res.Error
		}
		mDB.Conn.Exec("UPDATE user_summaries SET match_count = match_count + 1 WHERE user_id = ?", fromUserId)
		mDB.Conn.Exec("UPDATE user_summaries SET match_count = match_count + 1 WHERE user_id = ?", toUserId)

		return true, int(match.ID), nil
	}

	return false, -1, nil

}

// gets matches for this user
func (mDB *MatchPostgresRepo) Matches(userId int) ([]int, error) {
	matches := []sql.Match{}
	res := mDB.Conn.
		Or("user_a = ?", userId).
		Or("user_b = ?", userId).
		Find(&matches)

	if res.Error != nil {
		return nil, res.Error
	}

	ids := []int{}

	for i := 0; i < len(matches); i++ {
		if matches[i].UserB == userId {
			ids = append(ids, matches[i].UserA)
		} else {
			ids = append(ids, matches[i].UserB)
		}
	}
	return ids, nil
}

// gets all potencial matches for this user
func (mDB *MatchPostgresRepo) Discover(userId int) ([]int, error) {
	// select user_id from user_swipes where user_id not in (select to_user where from_user = userId)
	userIds := []int{}
	res := mDB.Conn.Raw("select user_id from user_summaries where user_id not in (select to_user from swipes where from_user = ?) and user_id <> ?", userId, userId).Scan(&userIds)
	if res.Error != nil {
		return userIds, res.Error
	}

	return userIds, nil

}

// adds user to the match summaries table
func (mDB *MatchPostgresRepo) AddUser(userId int) error {
	s := sql.UserSummary{UserId: userId, SwipeCount: 0, MatchCount: 0}
	res := mDB.Conn.Create((&s))
	if res.Error != nil {
		return res.Error
	}
	return nil
}
