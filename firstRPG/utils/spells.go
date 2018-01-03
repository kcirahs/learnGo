package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type magic struct {
	spell string
	mp int
	dmg int
	atr string
	desc string
}
var f1 = magic{"Flame", 10, 50, "fire", ""}
var i1 = magic{"Freeze", 10, 40, "ice", ""}


func GetSpells() {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "SPELL\tDMG\tMP\t")
	fmt.Fprintln(w, f1.spell + "\t" + fmt.Sprint(f1.dmg) + "\t" + fmt.Sprint(f1.mp) + "\t")
	fmt.Fprintln(w, i1.spell + "\t" + fmt.Sprint(i1.dmg) + "\t" + fmt.Sprint(i1.mp) + "\t")
	w.Flush()

}
