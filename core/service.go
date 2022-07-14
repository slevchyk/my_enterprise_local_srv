package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedpassword := string(bs)

	return hashedpassword, nil
}

func IsPasswordCorrect(hashedpassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))

	return err == nil
}

func GenerateOtp() string {
	randomeSource := rand.NewSource(time.Now().UnixNano())
	randomeData := rand.New(randomeSource)
	code := strconv.Itoa(randomeData.Int())[:6]

	return code
}

func ComposeMimeMail(to string, from string, subject string, body string) []byte {
	header := make(map[string]string)
	header["From"] = formatEmailAddress(from)
	header["To"] = formatEmailAddress(to)
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}

// Never fails, tries to format the address if possible
func formatEmailAddress(addr string) string {
	e, err := mail.ParseAddress(addr)
	if err != nil {
		return addr
	}
	return e.String()
}

func EncodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func DecodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}

func EncodeHmac(key, message string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	result := fmt.Sprintf("%x", h.Sum(nil))

	return result
}

func GanerateAccessToken(id uint64) (string, error) {

	var accessToken struct {
		Id        uint64    `json:"id"`
		ExpiresAt time.Time `json:"expires_at"`
	}

	accessToken.Id = id
	accessToken.ExpiresAt = time.Now().Add(24 * time.Hour)

	accessToken64, err := EncodeToBase64(accessToken)
	if err != nil {
		return "", fmt.Errorf("помилка кодування токену доступу: %s", err.Error())
	}

	codedAccessToken := EncodeHmac("password", accessToken64)

	authAccessToken := fmt.Sprintf("%s|%s", codedAccessToken, accessToken64)

	return authAccessToken, nil
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Eod(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 23, 59, 59, 59, t.Location())
}
