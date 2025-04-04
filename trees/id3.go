// Demonstrates decision tree classification
package main

import (
	"fmt"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
)

func main() {

	var tree base.Classifier

	rand.Seed(44111342)

	// Load in the data dataset
	data, err := base.ParseCSVToInstances("../data/clean_creditcard.csv", true)
	if err != nil {
		panic(err)
	}

	// Discretise the data dataset with Chi-Merge
	filt := filters.NewChiMergeFilter(data, 0.999)
	for _, a := range base.NonClassFloatAttributes(data) {
		filt.AddAttribute(a)
	}
	filt.Train()
	dataf := base.NewLazilyFilteredInstances(data, filt)

	// Create a 60-40 training-test split
	trainData, testData := base.InstancesTrainTestSplit(dataf, 0.60)
	tree = ensemble.NewRandomForest(70, 30)
	err = tree.Fit(trainData)
	if err != nil {
		panic(err)
	}
	predictions, err := tree.Predict(testData)
	if err != nil {
		panic(err)
	}
	fmt.Println("RandomForest Performance")
	cf, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(cf))
}
