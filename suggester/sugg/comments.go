package sugg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var comments = &commentList{
	comments: map[string]string{},
}

type commentList struct {
	comments map[string]string
}

var plHolder = regexp.MustCompile("%{.+?}")

// GetCommentText returns the final comment text
func (s *commentList) GetCommentText(cmt Comment) (string, error) {
	text, err := s.getText(cmt)
	if err != nil {
		return "", err
	}

	switch c := cmt.(type) {
	case *placeholderComment:
		placeHolder := plHolder.FindAllString(text, -1)
		for _, plH := range placeHolder {
			plHName := plH[2 : len(plH)-1]
			cnt, ok := c.params[plHName]
			if !ok {
				continue
			}

			text = strings.ReplaceAll(text, plH, cnt)
		}
		return text, nil
	case *comment:
		return text, nil
	}

	return "", errors.New("unknown comment type")
}

func (s *commentList) getText(cmt Comment) (string, error) {
	text, ok := s.comments[cmt.ID()]
	if ok {
		return text, nil
	}
	if err := s.load(cmt); err != nil {
		return "", err
	}

	return s.comments[cmt.ID()], nil
}

func (s *commentList) load(cmt Comment) error {
	uri := getURL(cmt)
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("loading comment %s returned with status code %d: %s", cmt.ID(), resp.StatusCode, uri)
	}

	s.comments[cmt.ID()] = string(bytes)
	return nil
}

var baseURL = "https://raw.githubusercontent.com/exercism/go-analyzer/comments/comments/"

func getURL(cmt Comment) string {
	return baseURL + strings.ReplaceAll(cmt.ID(), ".", "/") + ".md"
}
