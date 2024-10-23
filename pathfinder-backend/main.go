package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

type Point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type Request struct {
    Start Point `json:"start"`
    End   Point `json:"end"`
}

var grid = [20][20]int{}

func dfs(x, y int, endX, endY int, visited *[20][20]bool, path *[]Point) bool {
    if x < 0 || y < 0 || x >= 20 || y >= 20 || visited[x][y] || grid[x][y] == 1 {
        return false
    }

    visited[x][y] = true
    *path = append(*path, Point{X: x, Y: y})

    if x == endX && y == endY {
        return true
    }
    
    
    directions := []Point{
        {0, 1},    
        {1, 0},    
        {0, -1},   
        {-1, 0},   
        {-1, -1},  
        {1, 1},    
        {-1, 1},   
        {1, -1},   
    }
    
    
    for _, dir := range directions {
        if dfs(x+dir.X, y+dir.Y, endX, endY, visited, path) {
            return true
        }
    }

    // Backtrack if no path is found
    *path = (*path)[:len(*path)-1]
    return false
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func findPath(w http.ResponseWriter, r *http.Request) {
    if r.Method == "OPTIONS" {
        enableCors(&w)
        w.WriteHeader(http.StatusOK)
        return
    }

    enableCors(&w)

    var req Request
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    visited := [20][20]bool{}
    var path []Point

    if dfs(req.Start.X, req.Start.Y, req.End.X, req.End.Y, &visited, &path) {
        json.NewEncoder(w).Encode(map[string]interface{}{"path": path})
    } else {
        http.Error(w, "No path found", http.StatusNotFound)
    }
}

func main() {
    http.HandleFunc("/find-path", findPath)

    fmt.Println("Starting server on :8080...")
    
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
