package main

import (
	"flag"
	"fmt"

	"github.com/ItzAfroBoy/ft/calc"
	"github.com/ItzAfroBoy/ft/styles"
	lg "github.com/charmbracelet/lipgloss"
)

func flagCheck(flags ...float64) bool {
	var total float64
	for _, value := range flags {
		total += value
	}

	if total > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	frontPtr := flag.Float64("front", 0, "Percentage of weight at the front")
	sminPtr := flag.Float64("smin", 0, "Min spring value")
	smaxPtr := flag.Float64("smax", 0, "Max spring value")
	rminPtr := flag.Float64("rmin", 0, "Min rebound value")
	rmaxPtr := flag.Float64("rmax", 0, "Max rebound value")

	flag.Parse()
	if !flagCheck(*frontPtr, *sminPtr, *smaxPtr, *rminPtr, *rmaxPtr) {
		flag.Usage()
		return
	}

	frontARB, rearARB := calc.ARB(*frontPtr)
	frontSpring, rearSpring := calc.Springs(*frontPtr, *sminPtr, *smaxPtr)
	frontRebound, rearRebound, frontBump, rearBump := calc.Damping(*frontPtr, *rminPtr, *rmaxPtr)

	input := styles.Input.Render(fmt.Sprintf("%s\n\n\nFront Percentage: %.f%s\n\nMin Spring: %.1f\nMax Spring: %.1f\n\nMin Rebound: %.1f\nMax Rebound: %.1f", styles.Heading.Render(" Input "), *frontPtr, "%", *sminPtr, *smaxPtr, *rminPtr, *rmaxPtr))
	output := styles.Output.Render(fmt.Sprintf("%s\n\n\nFront ARB: %.2f\nRear ARB: %.2f\n\nFront Springs: %.1f\nRear Springs: %.1f\n\nFront Rebound: %.1f\n Rear Rebound: %.1f\nFront Bump: %.1f\nRear Bump: %.1f", styles.Heading.Render(" Output "), frontARB, rearARB, frontSpring, rearSpring, frontRebound, rearRebound, frontBump, rearBump))
	fmt.Println(lg.JoinHorizontal(lg.Top, input, output))
}
