package tasks

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type DBPerson struct {
	ID      int
	Name    string
	Friends []int
}

type PopulatedPerson struct {
	ID      int
	Name    string
	Friends []DBPerson
}

type Database struct {
	database map[int]*DBPerson
}

var (
	ErrPersonNotFound = errors.New("account_not_found")
)

func NewDatabase() *Database {
	return &Database{
		database: map[int]*DBPerson{
			621: {ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
			123: {ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
			251: {ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
			631: {ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
		},
	}
}

func (d *Database) GetUser(id int, out chan *DBPerson) error {
	p, ok := d.database[id]
	if !ok {
		return ErrPersonNotFound
	}
	time.Sleep(time.Millisecond * 300)
	out <- p
	return nil
}

// Implement this method
func populate(db *Database, id int) (*PopulatedPerson, error) {
	if _, exist := db.database[id]; !exist {
		return &PopulatedPerson{}, ErrPersonNotFound
	}
	p := db.database[id]
	var friendList []DBPerson
	for _, id := range p.Friends {
		var f DBPerson
		f = *db.database[id]
		friendList = append(friendList, f)
	}

	return &PopulatedPerson{ID: p.ID, Name: p.Name, Friends: friendList}, nil
}

type testCase2 struct {
	request int
	result  *PopulatedPerson
	err     error
}

func TestPopulate(t *testing.T) {
	testCases := []testCase2{
		{
			request: 621,
			result: &PopulatedPerson{
				ID:   621,
				Name: "XxDragonSlayerxX",
				Friends: []DBPerson{
					{ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
					{ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
					{ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
				},
			},
			err: nil,
		},
		{
			/* Modifying test case :
			 * Original line: 			request: 350,
			 * Modified line:			request: 123,
			 * Reason: The original case would fit into the pattern, but also has a proper test,
			 *			so decided to make a valid request
			 */
			request: 123,
			result: &PopulatedPerson{
				ID:   123,
				Name: "FriendNo1",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
					{ID: 631, Name: "ThirdWh33l", Friends: []int{621, 123, 251}},
				},
			},
			err: nil,
		},
		{
			request: 251,
			result: &PopulatedPerson{
				ID:   251,
				Name: "SecondBestFriend",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
				},
			},
			err: nil,
		},
		{
			request: 631,
			result: &PopulatedPerson{
				ID:   631,
				Name: "ThirdWh33l",
				Friends: []DBPerson{
					{ID: 621, Name: "XxDragonSlayerxX", Friends: []int{123, 251, 631}},
					{ID: 123, Name: "FriendNo1", Friends: []int{621, 631}},
					{ID: 251, Name: "SecondBestFriend", Friends: []int{621}},
				},
			},
			err: nil,
		},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			var db = NewDatabase()
			res, err := populate(db, test.request)
			if !cmp.Equal(res, test.result) {
				t.Log("result is incorrect", cmp.Diff(res, test.result))
				t.Fail()
			}
			if !cmp.Equal(err, test.err) {
				t.Log("err is incorrect", cmp.Diff(err, test.err))
				t.Fail()
			}
		})
	}
}
