package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestBasic(t *testing.T) {
	data := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}

	logs := []LogEntry{}

	for _, d := range data {
		logs = append(logs, parseLog(d))
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].key() < logs[j].key()
	})

	guardLogs := captureShifts(logs)
	for g, mins := range guardLogs {
		fmt.Printf("%v : %+v\n", g, mins)
	}

	gid := findMaxSleeper(guardLogs)
	if gid != "#10" {
		t.Errorf("Wrong guard is max sleeper: %v", gid)
	}

	min, count := daySleptMost(guardLogs[gid])
	if min != 24 {
		t.Errorf("Wrong minute: %v (count: %v)", min, count)
	}

	guard, minute, amount := findGuardMostAsleepMinute(guardLogs)

	if guard != "#99" || minute != 45 {
		t.Errorf("Wrong: %v, %v (total times sleeping on that minute: %v)", guard, minute, amount)
	}
}
