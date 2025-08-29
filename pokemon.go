package main

import (
	"fmt"
	"strings"
)

func (p *Pokemon) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Name: %s\n", p.name))
	sb.WriteString(fmt.Sprintf("Height: %d\n", p.height))
	sb.WriteString(fmt.Sprintf("Weight: %d\n", p.weight))
	sb.WriteString("Stats:\n")
	sb.WriteString(fmt.Sprintf("  - hp: %d\n", p.stats.hp))
	sb.WriteString(fmt.Sprintf("  - attack: %d\n", p.stats.attack))
	sb.WriteString(fmt.Sprintf("  - defense: %d\n", p.stats.defense))
	sb.WriteString(fmt.Sprintf("  - special-attack: %d\n", p.stats.special_attack))
	sb.WriteString(fmt.Sprintf("  - special-defense: %d\n", p.stats.special_defense))
	sb.WriteString(fmt.Sprintf("  - speed: %d\n", p.stats.speed))
	sb.WriteString("Types:\n")
	for _, t := range p.types {
		sb.WriteString(fmt.Sprintf("  - %s\n", t))
	}
	return sb.String()
}
