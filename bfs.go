package main

import "fmt"

func is_seller(name string) bool {
	return name[len(name)-1] == 'm'
}

var graph = make(map[string][]string)

func search(name string) bool {
	searched := make(map[string]bool)
	var queue []string
	var person string
	queue = append(queue, graph[name]...)
	for len(queue) != 0 {
		person = queue[0]
		queue = queue[1:]
		if searched[person] == false {
			if is_seller(person) {
				fmt.Println(person)
				return true
			}
			queue = append(queue, graph[person]...)
			searched[person] = true
		}
	}
	return false

}

func main() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	search("you")
}
