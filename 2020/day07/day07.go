package day07

import (
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	key      string
	parents  []*node
	children []childInfo
}

type childInfo struct {
	node   *node
	number int
}

func Resolve(input []byte) ([]interface{}, error) {
	n, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	return []interface {
	}{
		countParents(n),
		countChildren(n),
	}, nil
}

func parseInput(input []byte) (*node, error) {
	var (
		m          = map[string]*node{}
		rules      = strings.Split(string(input), "\n")
		childRegex = regexp.MustCompile(`(\d+)\s([a-z]+\s[a-z]+)`)
	)

	for _, rule := range rules {
		parentKey := strings.Split(rule, " bags contain")[0]
		parentNode := findOrCreateNode(m, parentKey)

		for _, match := range childRegex.FindAllStringSubmatch(rule, -1) {
			number, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}

			childNode := findOrCreateNode(m, match[2])
			parentNode.children = append(parentNode.children, childInfo{childNode, number})
			childNode.parents = append(childNode.parents, parentNode)
		}
	}

	return m["shiny gold"], nil
}

func findOrCreateNode(m map[string]*node, key string) *node {
	if _, ok := m[key]; !ok {
		m[key] = &node{key: key}
	}
	return m[key]
}

func countParents(n *node) int {
	type m map[string]bool
	var collect func(*node, m) m

	collect = func(n *node, parents m) m {
		for _, parent := range n.parents {
			parents[parent.key] = true
			collect(parent, parents)
		}
		return parents
	}

	return len(collect(n, m{}))
}

func countChildren(n *node) int {
	total := 0
	for _, childInfo := range n.children {
		total += childInfo.number + childInfo.number*countChildren(childInfo.node)
	}
	return total
}
