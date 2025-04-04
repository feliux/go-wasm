package main

import (
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	f, err := os.Open("../data/clean_creditcard_class.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create a dataframe from the CSV file. The types of the columns will be inferred.
	adDF := dataframe.ReadCSV(f)
	colName := "V10"
	colTarget := "V14"
	// Extract the target column.
	yVals := adDF.Col(colTarget).Float()
	classVals := adDF.Col("Class").Float()
	// Create a scatter plot for each of the features in the dataset.
	// pts will hold the values for plotting
	pts := make(plotter.XYs, adDF.Nrow())
	ptsAnomal := make(plotter.XYs, adDF.Nrow())
	// Fill pts with data.
	for i, floatVal := range adDF.Col(colName).Float() {
		if int(classVals[i]) == 0 {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		} else if int(classVals[i]) == 1 {
			ptsAnomal[i].X = floatVal
			ptsAnomal[i].Y = yVals[i]
		}
	}
	// Create the plot.
	p := plot.New()
	p.X.Label.Text = colName
	p.Y.Label.Text = colTarget
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	sAnomal, err := plotter.NewScatter(ptsAnomal)
	if err != nil {
		log.Fatal(err)
	}
	// s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	s.GlyphStyle.Radius = vg.Points(0.5)
	sAnomal.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	sAnomal.GlyphStyle.Radius = vg.Points(0.5)
	// Save the plot to a PNG file.
	p.Add(s)
	p.Add(sAnomal)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/"+colName+"_scatter.png"); err != nil {
		log.Fatal(err)
	}
}
