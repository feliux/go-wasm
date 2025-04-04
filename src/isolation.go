package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

var (
	csvCleanedWithoutClass string = "../data/clean_creditcard_noclass.csv"
	csvCleanedWithClass    string = "../data/clean_creditcard_class.csv"
)

// Adding Support for Converting Dataframe to FixedDataGrid
// https://github.com/sjwhitworth/golearn/pull/256/files

func main() {
	// Isolation Forest is used for outlier detection
	// The algorithm works by randomly splitting the data, so results will not be exactly reproducible
	// but generally outliers will still be classified as outliers
	csvData, err := base.ParseCSVToInstances(csvCleanedWithoutClass, true)
	if err != nil {
		panic(err)
	}
	// trainData, testData := base.InstancesTrainTestSplit(csvData, 0.50)

	// Create isolation forest with 100 trees, max depth 100, and each tree will use 850 datapoints
	forest := trees.NewIsolationForest(100, 100, 500)
	// Fit the isolation forest to the data. Note that all class attributes are also used during training.
	// Remove all class attributes you do nott want to use before calling fit
	forest.Fit(csvData)
	// Make predictions. Generally, IsolationForest is used for interpolation, not extrapolation.
	// Predictions are returned as Anomaly Scores from 0 to 1. close to 0 - not outlier, close to 1 - outlier
	preds := forest.Predict(csvData)
	counter := 0
	for i, _ := range preds {
		if preds[i] > 0.65 {
			counter++
		}
	}
	fmt.Println(counter)

	// We have to load the csv again cause no interface between golearn and gonum
	f, err := os.Open(csvCleanedWithClass)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the csv file. The types of the columns will be inferred.
	csvDataframe := dataframe.ReadCSV(f)
	classVals := csvDataframe.Col("Class").Float() // 492 outliers

	// Convert class column to bool
	var classes []bool
	for i, _ := range classVals {
		if classVals[i] > 0.5 {
			classes = append(classes, true)
		} else if classVals[i] < 0.5 {
			classes = append(classes, false)
		}
	}

	// Gonum stat need the data ordered
	// sort.Float64s(preds) // Bad cause we have to oreder classes too
	stat.SortWeightedLabeled(preds, classes, nil)
	tpr, fpr, _ := stat.ROC(nil, preds, classes, nil)

	// Compute Area Under Curve.
	auc := integrate.Trapezoidal(fpr, tpr)
	// fmt.Printf("true  positive rate: %v\n", tpr)
	// fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}
