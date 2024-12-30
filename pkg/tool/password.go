package tool

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	EncryptHeader string = "pbkdf2:sha256" // 用户密码加密
)

func NewPassword(password string) string {
	return EncryptPassword(password, RandomString(8), int(RandomInt(50000)+100000))
}

// ComparePassword
// if password in and db are same, return nil
func ComparePassword(in, db string) error {
	strs := strings.Split(db, "$")
	if len(strs) != 3 {
		return fmt.Errorf("password in db invalid: %s", db)
	}

	encs := strings.Split(strs[0], ":")
	if len(encs) != 3 {
		return fmt.Errorf("password in db invalid: %s", db)
	}

	encIteration, err := strconv.Atoi(encs[2])
	if err != nil {
		return fmt.Errorf("password in db invalid: %s, convert iter err: %s", db, err)
	}

	if EncryptPassword(in, strs[1], encIteration) != db {
		return fmt.Errorf("password input and db not same")
	}

	return nil
}

func EncryptPassword(password, salt string, iter int) string {
	hash := pbkdf2.Key([]byte(password), []byte(salt), iter, 32, sha256.New)
	encrypted := hex.EncodeToString(hash)
	return fmt.Sprintf("%s:%d$%s$%s", EncryptHeader, iter, salt, encrypted)
}

func CheckPassword(password string) error {
	if len(password) < 8 || len(password) > 32 {
		return errors.New("密码长度不符合")
	}

	var (
		err          error
		match        bool
		patternList  = []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[!@#%]+`} //, `[~!@#$%^&*?_-]+`}
		matchAccount = 0
		tips         = []string{"缺少数字", "缺少小写字母", "缺少大写字母", "缺少'!@#%'"}
		locktips     = make([]string, 0)
	)

	for idx, pattern := range patternList {
		match, err = regexp.MatchString(pattern, password)
		if err != nil {
			return errors.New("密码强度不够")
		}

		if match {
			matchAccount++
		} else {
			locktips = append(locktips, tips[idx])
		}
	}

	if matchAccount < 3 {
		return fmt.Errorf("密码强度不够, 可能 %s", strings.Join(locktips, ", "))
	}

	return nil
}
