package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule [][]string

func parse_rule(s string) (int, Rule) {
	parts := strings.Split(s, ":")
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	var rule Rule

	for _, part := range strings.Split(parts[1], "|") {
		var option []string

		for _, item := range strings.Fields(part) {
            if item[0] != '"' {
                _, err := strconv.Atoi(item)

                if err != nil {
                    log.Fatal(err)
                }
            }

			option = append(option, item)
		}

		rule = append(rule, option)
	}

	return index, rule
}

func parse_input(filename string) (map[int]Rule, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	parsing_rules := true
	rules := make(map[int]Rule)
	var samples []string

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			parsing_rules = false
		} else if parsing_rules {
			idx, rule := parse_rule(line)
			rules[idx] = rule
		} else {
			samples = append(samples, line)
		}
	}

    return rules, samples
}

func check_rule_inner(s string, rules map[int]Rule, rule_idx int, base int) (bool, int) {
    // fmt.Println("matching", s, "versus ", rules[rule_idx])

    for _, option := range rules[rule_idx] {
        pos := base
        matched := true

        for _, part := range option {
            if part[0] == '"' {
                // fmt.Print("matching symbol", part[1], "vs", s[pos])

                if part[1] != s[pos] {
                   matched = false
                } else {
                   pos++;
                }
            } else {
                subrule_idx, _ := strconv.Atoi(part)
                matched, pos = check_rule_inner(s, rules, subrule_idx, pos)
            }

            if !matched {
                break;
            }
        }

        if matched {
            // fmt.Println("matched", pos)
            return matched, pos
        }
    }

    return false, 0
}

func check_rule(s string, rules map[int]Rule) bool {
    res, pos := check_rule_inner(s, rules, 0, 0)
    return res && pos == len(s)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ", os.Args[0], " input_case.txt")
	}

	rules, samples := parse_input(os.Args[1])

	// fmt.Println(rules)
	// fmt.Println(samples)

    total := 0

	for _, s := range samples {
        if check_rule(s, rules) {
            // fmt.Println(s, " matched")
            total++;
        }
    }

    fmt.Printf("total: %d\n", total)
}
