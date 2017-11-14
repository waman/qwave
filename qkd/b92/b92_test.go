package b92

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"log"
	"time"
	"math/rand"
	"image/color"
	"math"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func ExampleB92Protocol(){
	rand.Seed(time.Now().UnixNano())

	n := 50
	aliceKey, bobKey := qkd.EstablishKeys(NewAlice(n), NewBob(n))

	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

type LoggingAlice struct {
	Alice
	Consumed int
}

func (alice *LoggingAlice) EstablishKey(ch qkd.ChannelOnAlice){
	for len(alice.key) < alice.n {
		bits  := qkd.NewRandomBits(qkd.ProperBitCount)
		ch.Qch() <- encode(bits)

		matches := <- ch.FromBob()
		newKey, m := qkd.AppendMatchingBitsLogged(alice.key, bits, matches, alice.n)
		alice.key = newKey
		alice.Consumed += m
	}
}

func ExampleB92ProtocolWithLoggingAlice(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	alice := &LoggingAlice{*NewAlice(n), 0}
	aliceKey, bobKey := qkd.EstablishKeys(alice, NewBob(n))

	log.Printf("Accepted Bit Rate: %f", float32(n)/float32(alice.Consumed))
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

func ExampleB92ProtocolWithEve(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEve())

	//log.Printf("Alice's key: %s", aliceKey)
	//log.Printf("Bob's key  : %s", bobKey)
	log.Printf("Concordance rate: %f", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// false
}

func ExampleSuccessRateOfEavesdropping(){
	rand.Seed(time.Now().UnixNano())

	const nTry = 10000
	nKeyMax := 20
	data := make(plotter.XYs, nKeyMax)
	for nKey := 1; nKey <= nKeyMax; nKey++ {
		matched := 0
		for i := 0; i < nTry; i++ {
			aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
				NewAlice(nKey), NewBob(nKey), qkd.NewObservingEve())

			if aliceKey.Equals(bobKey) { matched++ }
		}
		rate := 1-float64(matched)/float64(nTry)
		log.Printf("%d: %.3f\n", nKey, rate)

		i := nKey-1
		data[i].X = float64(nKey)
		data[i].Y = rate
	}
	plotData("points.png", data)
	fmt.Println("Done.")
	// Output:
	// Done.
}

func plotData(file string, data plotter.XYs){
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Aware Rate of Eavesdropping"
	p.X.Label.Text = "Key Length"
	p.Y.Label.Text = "Rate"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R:255, B:128, A:255}

	f := plotter.NewFunction(func(x float64) float64 { return 1-math.Pow(2.0/3.0, x)})
	f.Color = color.RGBA{B:255, A:255}

	p.Add(f, s)
	p.Legend.Add("Simulation", s)
	p.Legend.Add("Theoretical", f)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, file); err != nil {
		panic(err)
	}
}

func ExampleConcordanceRateBetweenAliceAndEveR(){
	rand.Seed(time.Now().UnixNano())

	nKey := 100000

	aliceKey, bobKey, eveKey := qkd.EstablishKeysWithEavesdropping(
		NewAlice(nKey), NewBob(nKey), NewEve(nKey))

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
			NewAlice(nKey), NewBob(nKey), NewEve(nKey))

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