package main

import "fmt"

type Node struct {
	number int
	name   string
	flag   bool
}

func main() {

	p := Node{number: 0, name: "Петя"}
	i := Node{number: 1, name: "Игорь", flag: true}
	v := Node{number: 2, name: "Вася"}
	s := Node{number: 3, name: "Слава"}
	k := Node{number: 4, name: "Кирилл"}
	d := Node{number: 5, name: "Даша"}
	m := Node{number: 6, name: "Маша"}
	visited := make(map[int]bool)
	graph := [7][]Node{
		{i, v},          //0
		{p, v},          //1
		{p, m, i, k, s}, //2
		{v},             //3
		{v, d},          //4
		{k, m},          //5
		{v, d},          //6
	}
	fmt.Print(dfs(0, visited, graph))
}

func dfs(start int, visited map[int]bool, graph [7][]Node) int {
	count := 0
	root := graph[start]
	visited[start] = true
	for node := range root {
		if !visited[root[node].number] {
			if !root[node].flag {
				count += dfs(root[node].number, visited, graph)
			} else {
				count++
				break
			}
		}
	}
	return count
}
