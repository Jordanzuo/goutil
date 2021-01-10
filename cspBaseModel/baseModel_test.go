package cspBaseModel

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

type Player struct {
	Id    int
	Name  string
	Level int
}

func (this *Player) Equal(other *Player) bool {
	return this.Id == other.Id && this.Name == other.Name && this.Level == other.Level
}

func (this *Player) String() string {
	return fmt.Sprintf("{Id:%d, Name:%s, Level:%d}", this.Id, this.Name, this.Level)
}

func NewPlayer(id int, name string, level int) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Level: level,
	}
}

// Use mutex as the way to prevent concurrency
type MutexPlayer struct {
	mutex     sync.RWMutex
	playerMap map[int]*Player
}

func (this *MutexPlayer) GetPlayerById(id int) (*Player, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	playerObj, exists := this.playerMap[id]
	return playerObj, exists
}

func (this *MutexPlayer) GetPlayerList() []*Player {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	playerList := make([]*Player, 0, len(this.playerMap))
	for _, value := range this.playerMap {
		playerList = append(playerList, value)
	}
	return playerList
}

func (this *MutexPlayer) AddPlayer(playerObj *Player) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.playerMap[playerObj.Id] = playerObj
}

func NewMutexPlayer() *MutexPlayer {
	return &MutexPlayer{
		playerMap: make(map[int]*Player, 1024),
	}
}

// Use csp as the way to prevent concurrency
type CspPlayer struct {
	playerMap    map[int]*Player
	baseModelObj *BaseModel
	cancel       context.CancelFunc
}

func NewCspPlayer() *CspPlayer {
	ctx, cancel := context.WithCancel(context.Background())
	cspPlayerObj := &CspPlayer{
		playerMap:    make(map[int]*Player, 1024),
		baseModelObj: NewBaseModel(),
		cancel:       cancel,
	}
	cspPlayerObj.baseModelObj.Start(ctx)

	// Register callback
	cspPlayerObj.baseModelObj.Register("GetPlayerById", cspPlayerObj.getPlayerByIdCallback)
	cspPlayerObj.baseModelObj.Register("GetPlayerList", cspPlayerObj.getPlayerListCallback)
	cspPlayerObj.baseModelObj.Register("AddPlayer", cspPlayerObj.addPlayerCallback)

	return cspPlayerObj
}

func (this *CspPlayer) Stop() {
	this.cancel()
}

type GetPlayerByIdResponse struct {
	Value  *Player
	Exists bool
}

func (this *CspPlayer) GetPlayerById(id int) (*Player, bool) {
	retCh := make(chan *GetPlayerByIdResponse)
	this.baseModelObj.RequestChannel <- NewRequest("GetPlayerById", id, retCh)
	responseObj := <-retCh
	return responseObj.Value, responseObj.Exists
}

func (this *CspPlayer) getPlayerByIdCallback(requestObj *Request) {
	id, _ := requestObj.Parameter.(int)
	playerObj, exists := this.playerMap[id]
	getPlayerByIdResponseObj := &GetPlayerByIdResponse{
		Value:  playerObj,
		Exists: exists,
	}
	retCh, _ := requestObj.ReturnCh.(chan *GetPlayerByIdResponse)
	retCh <- getPlayerByIdResponseObj
}

type GetPlayerListResponse struct {
	Value []*Player
}

func (this *CspPlayer) GetPlayerList() []*Player {
	retCh := make(chan *GetPlayerListResponse)
	this.baseModelObj.RequestChannel <- NewRequest("GetPlayerList", nil, retCh)
	responseObj := <-retCh
	return responseObj.Value
}

func (this *CspPlayer) getPlayerListCallback(requestObj *Request) {
	playerList := make([]*Player, 0, len(this.playerMap))
	for _, value := range this.playerMap {
		playerList = append(playerList, value)
	}
	getPlayerListResponseObj := &GetPlayerListResponse{
		Value: playerList,
	}

	retCh, _ := requestObj.ReturnCh.(chan *GetPlayerListResponse)
	retCh <- getPlayerListResponseObj
}

func (this *CspPlayer) AddPlayer(playerObj *Player) {
	this.baseModelObj.RequestChannel <- NewRequest("AddPlayer", playerObj, nil)
}

func (this *CspPlayer) addPlayerCallback(requestObj *Request) {
	playerObj, _ := requestObj.Parameter.(*Player)
	this.playerMap[playerObj.Id] = playerObj
}

func TestMutexPlayer(t *testing.T) {
	mutexPlayerObj := NewMutexPlayer()
	playerObj := NewPlayer(1, "Jordan", 100)
	mutexPlayerObj.AddPlayer(playerObj)
	playerList := mutexPlayerObj.GetPlayerList()
	if len(playerList) != 1 {
		t.Fatalf("Expected %d items in the list, now got:%d", 1, len(playerList))
	}

	gotPlayerObj, exists := mutexPlayerObj.GetPlayerById(1)
	if !exists {
		t.Fatalf("Expected:%v, Got:%v", true, exists)
	}
	if gotPlayerObj.Equal(playerObj) == false {
		t.Fatalf("Expected:%s, Got:%s", playerObj, gotPlayerObj)
	}
}

func TestCspPlayer(t *testing.T) {
	cspPlayerObj := NewCspPlayer()
	playerObj := NewPlayer(1, "Jordan", 100)
	cspPlayerObj.AddPlayer(playerObj)
	playerList := cspPlayerObj.GetPlayerList()
	if len(playerList) != 1 {
		t.Fatalf("Expected %d items in the list, now got:%d", 1, len(playerList))
	}

	gotPlayerObj, exists := cspPlayerObj.GetPlayerById(1)
	if !exists {
		t.Fatalf("Expected:%v, Got:%v", true, exists)
	}
	if gotPlayerObj.Equal(playerObj) == false {
		t.Fatalf("Expected:%s, Got:%s", playerObj, gotPlayerObj)
	}
}

func BenchmarkMutexPlayer(b *testing.B) {
	mutexPlayerObj := NewMutexPlayer()
	for i := 0; i < b.N; i++ {
		mutexPlayerObj.AddPlayer(NewPlayer(i, fmt.Sprintf("Player%d", i), i))
		mutexPlayerObj.GetPlayerById(i)
		mutexPlayerObj.GetPlayerList()
	}
}

func BenchmarkCspPlayer(b *testing.B) {
	cspPlayerObj := NewCspPlayer()
	for i := 0; i < b.N; i++ {
		cspPlayerObj.AddPlayer(NewPlayer(i, fmt.Sprintf("Player%d", i), i))
		cspPlayerObj.GetPlayerById(i)
		cspPlayerObj.GetPlayerList()
	}
	cspPlayerObj.Stop()
}
