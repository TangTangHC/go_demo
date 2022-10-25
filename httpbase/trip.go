package httpbase

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配的路由
	part     string  // 路由的一部分
	children []*node // 子节点
	isWild   bool    // 是否精准匹配 含有:或*时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%tm, children=%s\n}", n.pattern, n.part, n.isWild, n.children)
}

func (n *node) matchChild(part string) *node {
	for _, v := range n.children {
		if v.isWild || v.part == part {
			return v
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, v := range n.children {
		if v.isWild || v.part == part {
			nodes = append(nodes, v)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	matchChild := n.matchChild(part)
	if matchChild == nil {
		matchChild = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, matchChild)
	}
	matchChild.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil

}
