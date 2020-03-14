package store

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/lorandfazakas/grpc-rest-openapi-demo/pb"
)

var ToDos pb.ToDos

type ToDo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Urgency     string `json:"urgency"`
	IsDone      bool   `json:"isDone"`
}

// ReadJson reads the data.json for initial ToDos
func ReadJson() {
	var toDos []ToDo
	dataFile, err := ioutil.ReadFile("testdata/data.json")
	if err != nil {
		log.Println("error reading data file: ", err)
	}
	if err := json.Unmarshal(dataFile, &toDos); err != nil {
		log.Println("error parsing data file ", err)
	}
	for _, t := range toDos {
		ToDo := new(pb.ToDo)
		ToDo.Id = int32(t.ID)
		ToDo.Title = t.Title
		ToDo.Description = t.Description
		ToDo.Urgency = getUrgencyType(t.Urgency)
		ToDo.IsDone = t.IsDone
		ToDos.ToDo = append(ToDos.ToDo, ToDo)
	}
}

// GetAllToDos returns all ToDos
func GetAllToDos() (list pb.ToDos) {
	return ToDos
}

// GetToDoByID returns single ToDo by id
func GetToDoByID(id int32) *pb.ToDo {
	toDo := new(pb.ToDo)
	for _, res := range ToDos.ToDo {
		if res.Id == id {
			toDo = res
		}
	}
	return toDo
}

// CreateToDo creates a ToDo
func CreateToDo(t pb.ToDo) {
	ToDos.ToDo = append(ToDos.ToDo, &t)
}

// FinishToDo marks a ToDo as done
func FinishToDo(id int32) {
	for _, res := range ToDos.ToDo {
		if res.Id == id {
			res.IsDone = true
		}
	}
}

func getUrgencyType(s string) pb.ToDo_ToDoUrgency {
	var ret pb.ToDo_ToDoUrgency
	for v, k := range pb.ToDo_ToDoUrgency_value {
		if v == s {
			ret = pb.ToDo_ToDoUrgency(k)
			return ret
		}
	}
	return ret
}
