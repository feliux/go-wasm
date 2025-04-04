// Example of how to use CART trees for both Classification and Regression
package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {
	classificationData, err := base.ParseCSVToInstances("../data/clean_creditcard.csv", false)
	if err != nil {
		panic(err)
	}
	trainData, testData := base.InstancesTrainTestSplit(classificationData, 0.7)

	// Create New Classification Tree
	// Hyperparameters - loss function, max Depth (-1 will split until pure), list of unique labels
	decTree := trees.NewDecisionTreeClassifier("entropy", 500, []int64{0, 1})

	// Train Tree
	err = decTree.Fit(trainData)
	if err != nil {
		panic(err)
	}
	// Print out tree for visualization - shows splits and feature and predictions
	fmt.Println(decTree.String())

	// Access Predictions
	classificationPreds := decTree.Predict(testData)

	fmt.Println("Titanic Predictions")
	fmt.Println(classificationPreds)

	// Evaluate Accuracy on Test Data
	fmt.Println(decTree.Evaluate(testData))

	// Load House Price Data For Regression
	regressionData, err := base.ParseCSVToInstances("../../datasets/boston_house_prices.csv", false)
	if err != nil {
		panic(err)
	}
	trainRegData, testRegData := base.InstancesTrainTestSplit(regressionData, 0.5)

	// Hyperparameters - Loss function, max Depth (-1 will split until pure)
	regTree := trees.NewDecisionTreeRegressor("mse", -1)

	// Train Tree
	err = regTree.Fit(trainRegData)
	if err != nil {
		panic(err)
	}

	// Print out tree for visualization
	fmt.Println(regTree.String())

	// Access Predictions
	regressionPreds := regTree.Predict(testRegData)

	fmt.Println("Boston House Price Predictions")
	fmt.Println(regressionPreds)

}
