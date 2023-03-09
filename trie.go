package gout

import "strings"

// Tree node
type node struct {
	pattern  string  // route to be matched, e.g. /p/:lang
	part     string  // part of the route, e.g. :lang
	children []*node // children node, e.g. [doc, tutorial, intro]
	isWild   bool    // exact match or not, true when containing : and *
}

// The first node that matches successfully, for insertion
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part && child.isWild {
			return child
		}
	}
	return nil
}

// All nodes that match successfully for finding
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
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
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
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
