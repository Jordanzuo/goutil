package jsonUtil

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMarshalWithNumberType(t *testing.T) {
	src := make(map[string]int)
	src["Name"] = 123
	src["Money"] = 100000000

	var byteSlice []byte
	var err error
	if byteSlice, err = json.Marshal(src); err != nil {
		t.Errorf("Marshal src failed\n")
	}

	if target, err := UnMarshalWithNumberType(string(byteSlice)); err != nil {
		t.Errorf("Expected got nil, but got err:%s\n", err)
	} else {
		money, ok := target["Money"].(json.Number)
		money_int, err := money.Int64()
		if !ok || err != nil || money_int != 100000000 {
			t.Errorf("Expected got 100000000, but got %v, ok:%v, err:%s\n", money_int, ok, err)
		}

		fmt.Printf("target:%v\n", target)
	}
}

func TestDeepClone(t *testing.T) {
	p1 := NewPlayer(100000000, "Jordan")
	if p1_map, err := DeepClone(p1); err != nil {
		t.Errorf("Expected nil, but got err:%s", err)
	} else {
		fmt.Printf("p1:%s\n", p1)
		before := fmt.Sprintf("%v", p1_map)
		p1.Name = "Jordan Zuo"
		fmt.Printf("p1:%s\n", p1)
		after := fmt.Sprintf("%v", p1_map)
		fmt.Printf("before:%s\n", before)
		fmt.Printf("after:%s\n", after)
		if before != after {
			t.Errorf("Expected before and after same, but got different")
		}
	}
}

func TestDeepCloneWithNumberType(t *testing.T) {
	p1 := NewPlayer(100000000, "Jordan")
	if p1_map, err := DeepCloneWithNumberType(p1); err != nil {
		t.Errorf("Expected nil, but got err:%s", err)
	} else {
		fmt.Printf("p1:%s\n", p1)
		before := fmt.Sprintf("%v", p1_map)
		p1.Name = "Jordan Zuo"
		fmt.Printf("p1:%s\n", p1)
		after := fmt.Sprintf("%v", p1_map)
		fmt.Printf("before:%s\n", before)
		fmt.Printf("after:%s\n", after)
		if before != after {
			t.Errorf("Expected before and after same, but got different")
		}

		fmt.Printf("p1_interface_map:%v\n", p1_map)
		id, ok := p1_map["Id"].(json.Number)
		id_int, err := id.Int64()
		if !ok || err != nil || id_int != 100000000 {
			t.Errorf("Expected got 100000000, but got %v, ok:%v, err:%s\n", id_int, ok, err)
		}
	}
}

type Player struct {
	Id   int
	Name string
}

func (player *Player) String() string {
	return fmt.Sprintf("{Addr:%v, Id:%v, Name:%s}", &player, player.Id, player.Name)
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		Id:   id,
		Name: name,
	}
}