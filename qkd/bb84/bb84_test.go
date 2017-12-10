package bb84

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"time"
	"math/rand"
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
)

func ExampleBB84Protocol(){
	rand.Seed(time.Now().UnixNano())

	n := 50
	aliceKey, bobKey := qkd.EstablishKeys(NewAlice(n), NewBob(n))

	// fmt.Printf("Alice's key: %s...\n", aliceKey[:20])
	// fmt.Printf("Bob's key  : %s...\n", bobKey[:20])
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

type fmtgingAlice struct {
	Alice
	Consumed int
}

func (alice *fmtgingAlice) EstablishKey(ch qkd.ChannelOnAlice){
	for len(alice.key) < alice.n {
		bits  := qkd.NewRandomBits(qkd.ProperBitCount)
		bases := qkd.NewRandomBits(qkd.ProperBitCount)

		ch.Qch() <- encode(bits, bases)
		matches := matchBases(bases, <- ch.FromBob())

		ch.ToBob() <- matches
		newKey, consumed := qkd.AppendMatchingBitsLogged(alice.key, bits, matches, alice.n)
		alice.key = newKey
		alice.Consumed += consumed
	}
}

func ExampleBB84ProtocolWithLoggingAlice(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	alice := &fmtgingAlice{*NewAlice(n), 0}
	aliceKey, bobKey := qkd.EstablishKeys(alice, NewBob(n))

	fmt.Printf("Accepted Bit Rate: %.2f\n", float32(n)/float32(alice.Consumed))
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// Accepted Bit Rate: 0.50
	// true
	// 1
}

func ExampleBB84ProtocolWithEve(){
	rand.Seed(time.Now().UnixNano())

	n := 100000
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEve())

	//fmt.Printf("Alice's key: %s...\n", aliceKey[:20])
	//fmt.Printf("Bob's key  : %s...\n", bobKey[:20])
	fmt.Printf("Concordance rate: %.2f\n", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// Concordance rate: 0.75
	// false
}

func ExampleSuccessRateOfEavesdropping() {
	rand.Seed(time.Now().UnixNano())

	const nTry= 10000
	nKeyMax := 20
	data := make(plotter.XYs, nKeyMax)
	for nKey := 1; nKey <= nKeyMax; nKey++ {
		matched := 0
		for i := 0; i < nTry; i++ {
			aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
				NewAlice(nKey), NewBob(nKey), qkd.NewObservingEve())

			if aliceKey.Equals(bobKey) {
				matched++
			}
		}
		rate := 1 - float64(matched)/float64(nTry)
		// fmt.Printf("%d: %.3f\n", nKey, rate)

		i := nKey - 1
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

	f := plotter.NewFunction(func(x float64) float64 { return 1-math.Pow(0.75, x)})
	f.Color = color.RGBA{B:255, A:255}

	p.Add(f, s)
	p.Legend.Add("Simulation", s)
	p.Legend.Add("Theory", f)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, file); err != nil {
		panic(err)
	}
}

func ExampleConcordanceRateBetweenAliceAndEve(){
	rand.Seed(time.Now().UnixNano())

	nKey := 100000
	eve := NewEve(nKey)
	aliceKey, bobKey, eveKey := qkd.EstablishKeysWithEavesdropping(
		NewAlice(nKey), NewBob(nKey), eve)

	fmt.Printf("Sure bit rate of Eve: %.2f\n",
		float32(eve.SureKeyBitCount())/float32(eve.n))

	fmt.Printf("Concordance Rate between Alice and Bob: %.2f\n",
		aliceKey.ConcordanceRate(bobKey))
	fmt.Printf("Concordance Rate between Alice and Eve: %.2f\n",
		aliceKey.ConcordanceRate(eveKey))
	fmt.Printf("Concordance Rate between Bob and Eve: %.2f\n",
		bobKey.ConcordanceRate(eveKey))
	// Output:
	// Sure bit rate of Eve: 0.50
  // Concordance Rate between Alice and Bob: 0.75
  // Concordance Rate between Alice and Eve: 0.75
  // Concordance Rate between Bob and Eve: 0.75
}

func ExampleConcordanceRateBetweenAliceAndEveWhenEavesdroppingSucceed(){
	rand.Seed(time.Now().UnixNano())

	nTry := 100000
	nKey := 10
	n, sure, conc := 0, 0, 0
	for i := 0; i < nTry; i++ {
		eve := NewEve(nKey)
		aliceKey, bobKey, eveKey := qkd.EstablishKeysWithEavesdropping(
			NewAlice(nKey), NewBob(nKey), eve)

		if aliceKey.Equals(bobKey) {
			n++
			sure += eve.SureKeyBitCount()
			conc += aliceKey.Concordance(eveKey)
		}
	}

	base := float32(n*nKey)
	fmt.Printf("Sure bit rate of Eve: %.2f\n", float32(sure)/base)
	fmt.Printf("Concordance Rate between Alice and Eve: %.2f\n", float32(conc)/base)
	// Output:
	// Sure bit rate of Eve: 0.67
  // Concordance Rate between Alice and Eve: 0.83
}