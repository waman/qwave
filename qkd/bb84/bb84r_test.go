package bb84

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"log"
	"time"
	"math/rand"
	"gonum.org/v1/plot/plotter"
)

func ExampleBB84ProtocolWithEveR(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEveR())

	//log.Printf("Alice's key: %s", aliceKey)
	//log.Printf("Bob's key  : %s", bobKey)
	log.Printf("Concordance rate: %f", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
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
		log.Printf("%d: %.3f\n", nKey, rate)

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

	log.Printf("Concordance Rate between Alice and Bob: %.3f",
		aliceKey.ConcordanceRate(bobKey))
	log.Printf("Concordance Rate between Alice and Eve: %.3f",
		aliceKey.ConcordanceRate(eveKey))
	log.Printf("Concordance Rate between Bob and Eve: %.3f",
		bobKey.ConcordanceRate(eveKey))

	fmt.Println("Done.")
	// Output:
	// Done.
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

	log.Printf("Concordance Rate between Alice and Eve: %.3f",
		float32(conc)/float32(n*nKey))

	fmt.Println("Done.")
	// Output:
	// Done.
}