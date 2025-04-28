package main

type cliCommand struct {
	name		string
	description string
	callback	func(arg string) error
}

type config struct {
	Next		string
	Previous	string
	Results		[]struct {
					Name	string
					Url		string
	}
}

type APIclient struct {
	URL		string
}

type LocationArea struct {
	Pokemon_encounters []struct {
		Pokemon struct {
			Name 	string 
			URL  	string 
		}
	}
}

type Pokemon struct {
	Name				string
	Base_experience		int
	Height				int
	Weight				int
	Sprites				Sprite
	Stats 				[]struct {
						Base_stat		int
						Stat			Stat

	}
	Types				[]struct {
						Type		Type
	}
}

type Stat struct {
	Name			string

}

type Type struct {
	Name			string
}

type Sprite struct {
	Back_default	string
	Front_default	string
}
