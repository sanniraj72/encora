package main

import (
	"encoding/json"
	"log"
)

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func parse(v string) (*node, error) {

	root := &node{}
	// Parse here
	stack := []string{}
	// r is something to create a string, by which will create node.
	// We append alphabet to make an element
	r := ""
	for _, c := range v {
		if string(c) == "]" {
			// If "]" will encountered, then ready to create node from stack
			if r != "" {
				stack = append(stack, r)
			}
			r = ""
			root = parseToJson(stack, root)
			stack = []string{}
		} else if string(c) == "[" {
			// If "[" will encountered then simply add to stack
			if r != "" {
				stack = append(stack, r)
			}
			r = ""
			stack = append(stack, string(c))
		} else if string(c) == "," {
			// If "," will encountered then simply add to stack
			if r != "" {
				stack = append(stack, r)
			}
			r = ""
			stack = append(stack, string(c))
		} else {
			// create element
			r = r + string(c)
		}
	}
	return root, nil
}

// parseToJson - create json object from stack
func parseToJson(stack []string, root *node) *node {
	// Iterate through stack
	for i := len(stack) - 1; i > 0; i-- {
		// If [ will encountered in stack then create new node
		// and append to children till all [ is finished
		if stack[i] == "[" {
			// decrease counter by 1 to skip [
			i -= 1
			n := &node{}
			n.Name = stack[i]
			n.Children = append(n.Children, root)
			root = n
			continue
		}
		// If "," occurs, the create a new node and make it as children
		// of first node in children array
		if stack[i] == "," {
			// decrease counter by 1 to skip comma
			i -= 1
			n := node{
				Name:     stack[i],
				Children: []*node{},
			}
			root.Children = append(root.Children, &n)
			continue
		}
		// If there is no "[" and "," then there must be an element.
		// Create node and append to root children
		root.Children = append(root.Children, &node{
			Name:     stack[i],
			Children: []*node{},
		})
	}
	return root
}

func main() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
