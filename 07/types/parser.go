package types

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

func buildTree(tokens []Token) (*Entry, error) {
	root := &Entry{
		Name: "/",
		Size: 0,
	}
	cwd := root

	for _, token := range tokens {
		switch token.Type {
		case TokenCommand:
			if token.Args[0] == CommandChange {
				if token.Args[1] == "/" {
					cwd = root
					continue
				}

				if token.Args[1] == ".." {
					cwd = cwd.Parent
					continue
				}

				x := cwd.Child(token.Args[1])
				if x == nil {
					x = &Entry{
						Name:   token.Args[1],
						Size:   0,
						Parent: cwd,
					}
					cwd.Children = append(cwd.Children, x)
				}
				cwd = x
			}

		case TokenOutput:
			if p1 := token.Args[0]; p1 != "dir" {
				sz, err := strconv.Atoi(token.Args[0])
				if err != nil {
					return nil, err
				}

				x := &Entry{
					Name:   token.Args[1],
					Parent: cwd,
				}
				x.UpdateSize(sz)
				cwd.Children = append(cwd.Children, x)
			}
		}
	}

	return root, nil
}

func parseData(r io.Reader) ([]Token, error) {
	var out []Token

	s := bufio.NewScanner(r)

	for s.Scan() {
		line := s.Bytes()

		var t Token

		if line[0] == '$' {
			c := bytes.TrimSpace(line[1:])

			t = Token{
				Type: TokenCommand,
				Args: strings.Split(string(c), " "),
			}
		} else {
			t = Token{
				Type: TokenOutput,
				Args: strings.Split(string(line), " "),
			}
		}

		out = append(out, t)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func Parse(r io.Reader) (*Entry, error) {
	tokens, err := parseData(r)
	if err != nil {
		return nil, err
	}

	return buildTree(tokens)
}
