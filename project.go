package main

import(
"fmt"
"math"
)
//////////////////////////////////////////////
//////////////////////////////////////////////

type edge struct {

source int
dest int
cost float64
} 		// Edge between two nodes

type graph struct {
n int //number of nodes
matrix [][]float64 //Adjacency matrix(undirected)
} 	//Graph represented by Adjacency matrix
//////////////////////////////////////////////
//////////////////////////////////////////////

func new_graph(size int) graph{
G:= graph{n:size}
matrix:=make([][]float64,size)
for i:=0;i<size;i++{
	matrix[i]=make([]float64,size)
	for j:=0;j<size;j++{
		
		matrix[i][j]= math.Inf(1)
}
}
G.matrix=matrix
return G
} 	//Function to initialize a new graph
///////////////////////////////////////////////
//////////////////////////////////////////////

func (G *graph) AddEdge(e edge){

G.matrix[e.source][e.dest]=e.cost
G.matrix[e.dest][e.source]=e.cost
} 	//Method of graph:Adding new edge between nodes
//////////////////////////////////////////////
//////////////////////////////////////////////


func min_distance(G graph,dist[]float64,visited[]bool)int{
	min:= math.Inf(1)
	var min_index int
	for i:=0;i<G.n;i++{
		if visited[i]==false && dist[i]<=min{
			min=dist[i]
			min_index=i
		}
	}
	return min_index
}	 //Helper function to get minimum element of the array dist
//////////////////////////////////////////////
//////////////////////////////////////////////

func Dijkstra(G graph,source int, destination int) float64{

	distance:=make([]float64,G.n)
	visited:=make([]bool,G.n)
	for i:=0 ;i<G.n;i++{
		distance[i]= math.Inf(1)
		visited[i]=false
	}
	distance[source]=0


	for i:=0;i<G.n;i++{
		min:=min_distance(G,distance,visited)
		if visited[min]==false{
			visited[min]=true
			for j:=0;j<G.n;j++{
				if visited[j]==false{
					short:= distance[min] + G.matrix[min][j]
					if short<distance[j]{
						distance[j]=short
					}
				}
			}
		}
	}
	return distance[destination]
}		//Dijkstra shortest path algorithm
//////////////////////////////////////////////
//////////////////////////////////////////////

func main(){
var E1 edge
var E2 edge
var E3 edge
var E4 edge
var E5 edge
var E6 edge
var E7 edge
var E8 edge
var E9 edge
var E10 edge
var E11 edge
var E12 edge
var E13 edge
var E14 edge
var E15 edge

E1.source=0
E1.dest=1
E1.cost=50

E2.source=1
E2.dest=2
E2.cost=10

E3.source=0
E3.dest=4
E3.cost=80

E4.source=0
E4.dest=3
E4.cost=70

E5.source=0
E5.dest=10
E5.cost=60

E6.source=2
E6.dest=5
E6.cost=120

E7.source=4
E7.dest=5
E7.cost=20

E8.source=4
E8.dest=6
E8.cost=80

E9.source=3
E9.dest=6
E9.cost=100

E10.source=6
E10.dest=8
E10.cost=40

E11.source=6
E11.dest=7
E11.cost=30

E12.source=5
E12.dest=7
E12.cost=50

E13.source=7
E13.dest=9
E13.cost=40

E14.source=8
E14.dest=9
E14.cost=60

E15.source=10
E15.dest=8
E15.cost=150

G:=new_graph(11)
G.AddEdge(E1)
G.AddEdge(E2)
G.AddEdge(E3)
G.AddEdge(E4)
G.AddEdge(E5)
G.AddEdge(E6)
G.AddEdge(E7)
G.AddEdge(E8)
G.AddEdge(E9)
G.AddEdge(E10)
G.AddEdge(E11)
G.AddEdge(E12)
G.AddEdge(E13)
G.AddEdge(E14)
G.AddEdge(E15)
//////////////////////////////////////////////
//////////////////////////////////////////////
var name string
var source,destination int

	fmt.Print("Hello to MyGuide app. Please enter your name: ")
	fmt.Scanln(&name)
	fmt.Println(" ")
	fmt.Println("Welcome",name,)
	fmt.Println("Please read the following list:")
	fmt.Println("CITY NAME:		INDEX:")
	fmt.Println("New Cairo		0")
	fmt.Println("Madinaty		1")
	fmt.Println("Sherouk City		2")
	fmt.Println("Nasr City		3")
	fmt.Println("Heliopolis		4")
	fmt.Println("Sheraton		5")
	fmt.Println("Downtown		6")
	fmt.Println("Zamalek			7")
	fmt.Println("Dokki			8")
	fmt.Println("Sheikh Zayed		9")
	fmt.Println("Maadi			10")
	fmt.Println("From this list, please enter the index of your location.")
	fmt.Scanln(&source)
	fmt.Println("From the same list, please enter the index of your desired destination.")
	fmt.Scanln(&destination)
//////////////////////////////////////////////
//////////////////////////////////////////////
cost:=Dijkstra(G,source,destination) //finding minimum cost of moving from source to destination
fmt.Println("This trip will cost you",cost,"EGP")
} // My_Guide Program
