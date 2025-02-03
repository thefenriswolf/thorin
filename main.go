package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fp, tf := parseArgs(os.Args)
	if fp != "" {
		fc, err := getFileContents(fp)
		if err != nil {
			log.Fatal(err)
		}
		fle := len(fc)
		data := collectData(fc, fle)
		d := getTimeframe(data, tf)
		for i := range d {
			fmt.Printf("%v %v\t%v€ %v\n",
				d[i].date.Format(outputDateFormat),
				d[i].category,
				d[i].amount,
				d[i].note,
				// d[i].income
			)
		}
		fmt.Println("-----------------------------------------")
		sum := genSum(d, 0, len(data))
		fmt.Printf("Sum: %.2f\n", sum)
		//startWebview()
	}
}
