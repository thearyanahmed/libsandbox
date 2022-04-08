package linked_list

import "fmt"

type Tweet struct {
	Id     int
	UserId int
}

type Twitter struct {
	Tweets            []Tweet
	FollowersRegistry map[int][]int
}

// method name Constructor() in leetcode
func NewTwitter() Twitter {
	var tweets []Tweet
	followers := make(map[int][]int)

	return Twitter{
		Tweets:            tweets,
		FollowersRegistry: followers,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{UserId: userId, Id: tweetId}
	this.Tweets = append(this.Tweets, tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// get followers id
	l := len(this.FollowersRegistry[userId])
	usersToBeChecked := make(map[int]bool, len(this.FollowersRegistry[userId])+1)

	for i := 0; i < l; i++ {
		usersToBeChecked[this.FollowersRegistry[userId][i]] = true
	}

	usersToBeChecked[userId] = true
	// fmt.Println("users_to_be_checked", usersToBeChecked, "user_id", userId)

	var feed []int

	feedCount := 0 

	for i := len(this.Tweets) - 1; i >= 0; i-- {
		if _, ok := usersToBeChecked[this.Tweets[i].UserId]; ok {
			if feedCount > 9 {
				break
			}
			feed = append(feed, this.Tweets[i].Id)

			feedCount++
		}
	}

	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// if a list exists, add the follower to the list
	if following, ok := this.FollowersRegistry[followerId]; ok {
		this.FollowersRegistry[followerId] = append(following, followeeId)
		return
	}

	// else create new list
	this.FollowersRegistry[followerId] = []int{followeeId}
}

func (this *Twitter) GetFollowers(followeeId int) {
	fmt.Println("followers of ", followeeId, this.FollowersRegistry[followeeId])
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if following, ok := this.FollowersRegistry[followerId]; ok {
		l := len(following)

		for i := 0; i < l; i++ {
			if following[i] == followeeId {
				// take it to the very last
				following[i], following[l-1] = following[l-1], following[i]
				break
			}
		}
		if l > 0 {
			// remove the last member
			following = following[:l-1]
		}

		this.FollowersRegistry[followerId] = following
	}
}
