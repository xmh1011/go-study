package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	// now 为现在的时间，yearAgo 为距现在一年的时间，monthAgo 为距现在一月的时间。
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)
	monthAgo := now.AddDate(0, -1, 0)

	// 三个切片，用来存储 不足一个月的问题，不足一年的问题，超过一年的问题。
	yearAgos := make([]*Issue, 0)
	monthAgos := make([]*Issue, 0)
	lessMonths := make([]*Issue, 0)

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// 如果 yearAgo 比 创建时间晚，说明超过一年
		if yearAgo.After(item.CreatedAt) {
			yearAgos = append(yearAgos, item)
			// 如果 monthAgo 比 创建时间晚，说明超过一月 不足一年
		} else if monthAgo.After(item.CreatedAt) {
			monthAgos = append(monthAgos, item)
			// 如果 monthAgo 比 创建时间早，说明不足一月。
		} else if monthAgo.Before(item.CreatedAt) {
			lessMonths = append(lessMonths, item)
		}
	}

	fmt.Printf("\n一年前\n")
	for _, item := range yearAgos {
		fmt.Printf("#%-5d %9.9s %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("\n一月前\n")
	for _, item := range monthAgos {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Printf("\n不足一月\n")
	for _, item := range lessMonths {
		fmt.Printf("#%-5d %9.9s %.55s %-40v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
