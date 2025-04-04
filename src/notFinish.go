package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

var (
	csvCleanedWithoutClass string = "../data/clean_creditcard_noclass.csv"
	csvCleanedWithClass    string = "../data/clean_creditcard_class.csv"
	// SeriesInit{}
)

func main() {
	ctx := context.Background()
	f, err := os.Open(csvCleanedWithClass)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	df, err := imports.LoadFromCSV(ctx, f)
	if err != nil {
		panic(err)
	}
	// fmt.Print(df.Table())
	dataGrid := base.ConvertDataFrameToInstances(df, 28)
	forest := trees.NewIsolationForest(100, 100, 500)
	forest.Fit(dataGrid)
	preds := forest.Predict(dataGrid)
	counter := 0
	for i, _ := range preds {
		if preds[i] > 0.65 {
			counter++
		}
	}
	fmt.Println(counter)
}
