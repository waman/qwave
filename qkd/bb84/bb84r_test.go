package bb84

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"time"
	"math/rand"
	"gonum.org/v1/plot/plotter"
)

func ExampleBB84ProtocolWithEveR(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEveR())

	//fmt.Printf("Alice's key: %s...\n", aliceKey[:20])
	//fmt.Printf("Bob's key  : %s...\n", bobKey[:20])
	fmt.Printf("Concordance rate: %.2f\n", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// Concordance rate: 0.75
	// false
}

func ExampleSuccessRateOfEavesdroppingR(){
	rand.Seed(time.Now().UnixNano())

	const nTry = 10000
	nKeyMax := 20
	data := make(plotter.XYs, nKeyMax)
	for nKey := 1; nKey <= nKeyMax; nKey++ {
		matched := 0
		for i := 0; i < nTry; i++ {
			aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
				NewAlice(nKey), NewBob(nKey), qkd.NewObservingEveR())

			if aliceKey.Equals(bobKey) { matched++ }
		}
		rate := 1-float64(matched)/float64(nTry)
		// fmt.Printf("%d: %.3f\n", nKey, rate)

		i := nKey-1
		data[i].X = float64(nKey)
		data[i].Y = rate
	}
	plotData("points-r.png", data)
	fmt.Println("Done.")
	// Output:
	// Done.
}

func ExampleConcordanceRateBetweenAliceAndEveR(){
	rand.Seed(time.Now().UnixNano())

	nKey := 100000
	aliceKey, bobKey, eveKey := qkd.EstablishKeysWithEavesdropping(
		NewAlice(nKey), NewBob(nKey), NewEveR(nKey))

	fmt.Printf("Concordance Rate between Alice and Bob: %.3f\n",
		aliceKey.ConcordanceRate(bobKey))
	fmt.Printf("Concordance Rate between Alice and Eve: %.2f\n",
		aliceKey.ConcordanceRate(eveKey))
	fmt.Printf("Concordance Rate between Bob and Eve: %.2f\n",
		bobKey.ConcordanceRate(eveKey))
	// Output:
	// Concordance Rate between Alice and Bob: 0.750
  // Concordance Rate between Alice and Eve: 0.66
  // Concordance Rate between Bob and Eve: 0.66
}

func ExampleConcordanceRateBetweenAliceAndEveWhenEavesdroppingSucceedR(){
	rand.Seed(time.Now().UnixNano())

	nTry := 100000
	nKey := 10
	n, conc := 0, 0
	for i := 0; i < nTry; i++ {
		aliceKey, bobKey, eveKey := qkd.EstablishKeysWithEavesdropping(
			NewAlice(nKey), NewBob(nKey), NewEveR(nKey))

		if aliceKey.Equals(bobKey) {
			n++
			conc += aliceKey.Concordance(eveKey)
		}
	}

	fmt.Printf("Concordance Rate between Alice and Eve: %.3f\n",
		float32(conc)/float32(n*nKey))
	// Output:
  // Concordance Rate between Alice and Eve: 0.710
}