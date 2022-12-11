package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []uint64 // worry levels exactly
	operation func(uint64, uint64) uint64
	div       uint64
	op2       uint64 // Operation second operand
	pos_rx    uint64
	neg_rx    uint64
	actions   int
}

func (m *monkey) divisible(new uint64) bool {
	return new%m.div == 0
}

func (m *monkey) addItem(item uint64) {
	m.items = append(m.items, item)
}

func (m *monkey) unsetItems() {
	m.items = make([]uint64, 0)
}

var (
	// F*ng monkeys
	monkeys = []monkey{}
	// Operations for new worry level
	ops = map[string]func(uint64, uint64) uint64{
		"*": func(old, op2 uint64) (new uint64) { return old * op2 },
		"+": func(old, op2 uint64) (new uint64) { return old + op2 },
	}
	// Parser for one monkey
	monkey_regex = []string{
		`Monkey (\d+):\n\s+Starting items: (?P<items>.*)`,
		`\s+Operation: new = old (?P<operation>.) (?P<op2>\S+)`,
		`\s+Test: divisible by (?P<div>\d+)`,
		`\s+If true: throw to monkey (?P<pos_rx>\d+)`,
		`\s+If false: throw to monkey (?P<neg_rx>\d+)`,
	}
)

func main() {
	p1, p2 := day_11("input.txt")
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)
}

func day_11(input string) (p1, p2 int) {
	data, _ := os.ReadFile(input)

	// Prepare regexps and parse file
	regex := fmt.Sprintf("(%s)", strings.Join(monkey_regex, `\n`))
	re := regexp.MustCompile(regex)
	res := re.FindAllStringSubmatch(string(data), -1)

	// Get monkeys
	// one outer match = one monkey
	for _, match := range res {
		monkey := monkey{}
		for grp_id, var_name := range re.SubexpNames() {
			switch var_name {
			case "items":
				monkey.items = getItems(match[grp_id])
			case "operation":
				monkey.operation = ops[match[grp_id]]
			case "div":
				monkey.div = getInt("divisor", match[grp_id])
			case "pos_rx":
				monkey.pos_rx = getInt("positive recepient", match[grp_id])
			case "neg_rx":
				monkey.neg_rx = getInt("negative recepient", match[grp_id])
			case "op2":
				if match[grp_id] == "old" {
					monkey.op2 = 0
				} else {
					monkey.op2 = getInt("second operand", match[grp_id])
				}
			}
		}
		monkeys = append(monkeys, monkey)
	}
	p1 = monkeyBusiness(monkeys, 20, 3)
	// p2 = monkeyBusiness(monkeys, 10000, 1)

	return
}

func monkeyBusiness(m []monkey, rounds int, damper float64) int {
	copy(monkeys, m)
	// Monkey business starts here
	// 20 rounds
	for k := 0; k < rounds; k++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeys[i].actions++
				// Apply operation
				// 0 = old as second operand
				var new_w_lvl uint64
				if monkey.op2 == 0 {
					new_w_lvl = monkey.operation(item, item)
				} else {
					new_w_lvl = monkey.operation(item, monkey.op2)
				}
				if item > new_w_lvl {
					panic("Uint64 overflowed")
				}
				// boring...
				new_w_lvl = uint64(math.Floor(float64(new_w_lvl) / damper))
				// throw to another monkey
				if new_w_lvl%monkey.div == 0 {
					monkeys[monkey.pos_rx].addItem(new_w_lvl)
				} else {
					monkeys[monkey.neg_rx].addItem(new_w_lvl)
				}
			}
			// Clear monkey items as processed
			monkeys[i].unsetItems()
		}
	}

	actions := make([]int, 0)
	for _, m := range monkeys {
		actions = append(actions, m.actions)
	}
	sort.Ints(actions)
	return actions[len(actions)-2] * actions[len(actions)-1]
}

// Get any int param
func getInt(p_name, s string) uint64 {
	d, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Wrong %s %s", p_name, s)
	}
	return uint64(d)
}

// Converts items from string to slice of int
func getItems(s string) (items []uint64) {
	for _, v := range strings.Split(s, ",") {
		item, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			log.Fatalf("Wrong item %s", v)
		}
		items = append(items, uint64(item))
	}
	return
}
