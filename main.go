package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var GlobalVariance float64

func main() {
	//ouvrir le fichier, le scanner et stocker le contenu en slices
	var sumfile = []int{}
	file, _ := os.Open("data.txt")
	filelector := bufio.NewScanner(file) // lire le contenu du fichier ligne par ligne.
	filelector.Split(bufio.ScanLines)    //: CconfigureR le scanner pour diviser le contenu du fichier en lignes en utilisant la méthode bufio.ScanLines.
	var filetextscanner []string
	for filelector.Scan() { //lit le contenu du fichier ligne par ligne en utilisant la méthode Scan
		filetextscanner = append(filetextscanner, filelector.Text())
	}

	file.Close()

	n := 0
	for _, eachline := range filetextscanner {
		if eachline == "" {
			continue
		}

		n = n + 1
		j, _ := strconv.Atoi(eachline)

		sumfile = append(sumfile, j) // ajouter lecontenu de chaque ligne à la fin de la slice(sumfile)
		//donc sumfile est une liste successif des nombres contenus dans le fichier
	}

	a := int(math.Round(Average(sumfile)))
	fmt.Println("Average:", a)

	m := int(math.Round(Median(sumfile)))
	fmt.Println("Median:", m)

	v := int(math.Round(Variance(sumfile)))
	fmt.Println("Variance", v)

	s := StandardDeviation(sumfile)
	fmt.Println("Standard Deviation", s)

}

func Average(s []int) float64 {
	sum := 0
	for _, value := range s {
		sum += value
	}
	return float64(sum) / float64(len(s))
}

func Median(s []int) float64 {
	sort.Ints(s)
	middle := len(s) / 2
	if len(s)%2 == 0 {
		return float64((s[middle-1] + s[middle])) / 2.0
	}
	return float64(s[middle])
}

func Variance(s []int) float64 {
	avg := Average(s)
	var superieur float64
	for _, value := range s {
		vf := float64(value)
		superieur += (vf - avg) * (vf - avg)

	}
	return float64(superieur / float64(len(s)))
}

func StandardDeviation(s []int) int {
	return int(math.Round((math.Sqrt(float64(Variance(s))))))
}
