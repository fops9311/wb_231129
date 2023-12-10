package main

import "fmt"

// тип, который не соотвествует интерфейсу
type Human struct {
	Intellect int
	Strength  int
	Speed     int
}
type Result bool

const (
	Success Result = true
	Fail    Result = false
)

// методы исходного типа
func (h Human) IntellectTask(challenge int) Result {
	return challenge <= h.Intellect
}
func (h Human) StrengthTask(challenge int) Result {
	return challenge <= h.Strength
}
func (h Human) SpeedTask(challenge int) Result {
	return challenge <= h.Speed
}

type Difficulty int

const (
	Trivial Difficulty = iota
	Easy
	Demanding
	Hard
	Impossible
)

// интерфейс, которому мы хотим написать адаптер
type QuestParticipant interface {
	SolvePazzle(challenge Difficulty) Result
	RunAMile() Result
	CarryLoad(strChallenge Difficulty, speedChallenge Difficulty) Result
}

// адаптер, которые реализует методы интерфейса с использованием методов исходного типа
type HumanQuestParticipant struct {
	H Human
}

// реализация интерфейса
func (h HumanQuestParticipant) CarryLoad(strChallenge Difficulty, speedChallenge Difficulty) Result {
	return h.H.StrengthTask(int(strChallenge)) && h.H.SpeedTask(int(speedChallenge))
}
func (h HumanQuestParticipant) RunAMile() Result {
	return h.H.SpeedTask(int(Demanding))
}
func (h HumanQuestParticipant) SolvePazzle(challenge Difficulty) Result {
	return h.H.StrengthTask(int(Impossible)) || h.H.IntellectTask(int(challenge))
}

// класс, который не совместим с Human
type Quest struct {
	SolvePazzleDifficulty  Difficulty
	CarryLoadStrDifficulty Difficulty
	CarryLoadSpdDifficulty Difficulty
}

func (q Quest) Run(p QuestParticipant) {

	fmt.Println("SolvePazzle\tsuccess:", p.SolvePazzle(q.SolvePazzleDifficulty))
	fmt.Println("RunAMile\tsuccess:", p.RunAMile())
	fmt.Println("CarryLoad\tsuccess:", p.CarryLoad(q.CarryLoadSpdDifficulty, q.CarryLoadSpdDifficulty))
}

func main() {
	//объявление интерфейса с адаптером
	var participant QuestParticipant = HumanQuestParticipant{
		H: Human{
			Intellect: 2,
			Strength:  3,
			Speed:     2,
		},
	}
	//использование адаптера
	Quest{
		SolvePazzleDifficulty:  Hard,
		CarryLoadStrDifficulty: Demanding,
		CarryLoadSpdDifficulty: Demanding,
	}.Run(participant)
}
