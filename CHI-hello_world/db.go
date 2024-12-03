package main

var dataBase map[string]*Athlete

func BuildDb() {
	startsAthlets := make(map[string]*Athlete)
	startsAthlets["Francisco Italo"] = &Athlete{
		ID:   "1",
		Name: "Francisco Italo",
	}
	startsAthlets["alfredo"] = &Athlete{
		ID:   "2",
		Name: "alfredo grugel",
	}

	dataBase = startsAthlets
}