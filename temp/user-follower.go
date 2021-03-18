package temp

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type BadgeType string

const (
	Official   BadgeType = "Official"
	Unofficial BadgeType = "Unofficial"
)

var (
	ErrUserNum     = errors.New("ユーザーの数 n は 1<=n<=10^4 の整数を入力してください。")
	ErrFollowerNum = errors.New("フォロワーの人数 x は 1<=x<=10^9 の整数を入力してください。")
	ErrUserID      = errors.New("ユーザーIDは半角英字アルファベット1文字以上20文字以下に設定してください。")
	ErrBadgeType   = errors.New("バッジは'Official'か'Unofficial'のどちらかに設定してください。")
	ErrCondition   = errors.New("条件ユーザーの数(n)とフォロワーの人数(x)は'n x'形式で入力してください。")
	ErrUser        = errors.New("ユーザーID(id)、バッジの付与(s)とフォロワーの人数(f)は'id s f'形式で入力してください。")
	ErrUniqueUser  = errors.New("ユーザーIDはすでに存在しています。")
	ErrNoUserMatch = errors.New("条件を満たすユーザーが無いです。")
)

func (b BadgeType) String() string {
	return string(b)
}

type Condition struct {
	UserNum     int
	FollowerNum int
}

type User struct {
	UserID      string
	Badge       BadgeType
	FollowerNum int
	LineNo      int
}

func (c *Condition) setUserNum(n string) error {
	userNum, err := strconv.Atoi(n)
	if err != nil || userNum < 1 || userNum > int(math.Pow10(4)) {
		return ErrUserNum
	}
	c.UserNum = userNum
	return nil
}

func (c *Condition) setFollowerNum(n string) error {
	followerNum, err := strconv.Atoi(n)
	if err != nil || followerNum < 1 || followerNum > int(math.Pow10(9)) {
		return ErrFollowerNum
	}
	c.FollowerNum = followerNum
	return nil
}

func (u *User) setUserID(userID string) error {
	r := regexp.MustCompile(`[A-Za-z]{1,20}`)
	if r.MatchString(userID) {
		u.UserID = userID
		return nil
	} else {
		err := fmt.Errorf("'%s'の%w", userID, ErrUserID)
		return err
	}
}
func (u *User) setBadge(b string) error {
	badge := BadgeType(b)
	if badge != Official && badge != Unofficial {
		return ErrBadgeType
	}
	u.Badge = badge
	return nil
}

func (u *User) setFollowerNum(n string) error {
	followerNum, err := strconv.Atoi(n)
	if err != nil || followerNum < 1 || followerNum > int(math.Pow10(9)) {
		return ErrFollowerNum
	}
	u.FollowerNum = followerNum
	return nil
}

func getStdin() []string {
	stdin := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if stdin.Text() == "" {
			break
		}
		lines = append(lines, stdin.Text())
	}
	return lines
}

func UserFollower() {
	lines := getStdin()
	con := Condition{}
	userMap := make(map[string]User)
	userLineMap := make(map[int]User)
	for i, v := range lines {
		line := strings.Split(v, " ")
		if i == 0 {
			if len(line) != 2 {
				fmt.Printf("第%d行error:%s", i, ErrCondition.Error())
				return
			} else {
				if err := con.setUserNum(line[0]); err != nil {
					fmt.Printf("第%d行error:%s", i, err)
					return
				}
				if err := con.setFollowerNum(line[1]); err != nil {
					fmt.Printf("第%d行error:%s", i, err)
					return
				}
			}

		} else {
			user := User{}
			if len(line) != 3 {
				fmt.Printf("第%d行error:%s:", i, ErrUser)
				return
			} else {
				if err := user.setUserID(line[0]); err != nil {
					fmt.Printf("第%d行error:%s:", i, err)
					return
				}
				if err := user.setBadge(line[1]); err != nil {
					fmt.Printf("第%d行error:%s:", i, err)
					return
				}
				if err := user.setFollowerNum(line[2]); err != nil {
					fmt.Printf("第%d行error:%s:", i, err)
					return
				}
				user.LineNo = i + 1
			}
			if _, ok := userMap[user.UserID]; ok {
				fmt.Printf("第%d行error:'%s'の%s", i, user.UserID, ErrUniqueUser)
				return
			} else {
				userMap[user.UserID] = user
			}
		}
	}

	for _, user := range userMap {
		userLineMap[user.LineNo] = user
	}
	var tmpSlice = make([]int, 0, len(userLineMap))
	for k := range userLineMap {
		tmpSlice = append(tmpSlice, k)
	}
	sort.Ints(tmpSlice)
	i := 0
	for _, key := range tmpSlice {
		user := userLineMap[key]
		if user.Badge == Official && user.FollowerNum >= con.FollowerNum {
			fmt.Println(user.UserID)
		}
		i += 1
	}
	if i == 0 {
		fmt.Println(ErrNoUserMatch)
	}
}
