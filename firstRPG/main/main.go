package main


import "../utils"

func main() {
	utils.GetSpells()

}


	/*magic := []map[string]interface{}{
		{"Spell": "Fire", "MP": 10, "dmg": 50},
		{"Spell": "Lightning", "MP": 10, "dmg": 60},
		{"Spell": "Ice", "MP": 10, "dmg": 40},
	}
	fmt.Println(magic)
	fmt.Println(magic[0]["Spell"], magic[0]["MP"])

	for i:= range magic{
		for k, v := range magic[i]{


		fmt.Println(k, v)
		}
	}
	*///magic := []string{
	//	"Spell",
	//	"dmg",
	//}
	//fmt.Println(magic)
	//fmt.Println(magic[0])
