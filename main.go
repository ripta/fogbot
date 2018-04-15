package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ripta/fogbot/client"
)

func main() {
	hostname := os.Getenv("FOGBOT_HOSTNAME")
	c, err := client.New(hostname)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if token, ok := os.LookupEnv("FOGBOT_TOKEN"); ok {
		c.Token = client.Token(token)
	} else {
		user := os.Getenv("FOGBOT_USERNAME")
		pass := os.Getenv("FOGBOT_PASSWORD")
		if err := c.Login(user, pass); err != nil {
			log.Fatalf("%v", err)
		}
	}

	// Uncomment one of:
	//
	// loadCases(c, []int{100844, 69732, 91084})
	// loadCases(c, []int{106175})
	// loadCases(c, []int{91084})
	// loadCase(c, 69732)
	// search(c, "SQS", 10)
	// listStatuses(c)
	// listTags(c)
}

func addCheckin(c *client.Client) {
	ci := &client.Checkin{
		BugID:        1337,
		Repository:   1,
		File:         "README.md",
		PrevRevision: "fdd2a83f8327a3f5068c2a30dd14fedd2d205c12",
		NewRevision:  "bedfc29ab9cebb0c2d97a632bbc7b56b3f7c7cf4",
	}
	if err := c.NewCheckin(ci); err != nil {
		log.Fatalf("%+v", err)
	}
}

func displayBug(bug *client.Case) {
	fmt.Printf("%6d: {%s} %s\n", bug.ID, bug.StatusName(), bug.Title)
	if tags := bug.Tags(); len(tags) > 0 {
		fmt.Printf("\tTags: %+v\n", bug.Tags())
	}

	fmt.Printf("\tLast Event ID: %d\n", bug.LatestBugEventID)
	es := bug.LatestEvents(5)
	// es := bug.Events
	for idx, evt := range es {
		fmt.Printf("\t(%d/%d) %d: %s\n", len(bug.Events)-len(es)+idx+1, len(bug.Events), evt.ID, evt.Description)
		if evt.Text != "" {
			for _, line := range evt.WrappedText() {
				fmt.Printf("\t\t%s\n", line)
			}
		}
	}

	cis, err := bug.ListCheckins()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	fmt.Printf("\tCommits: %d\n", len(cis))
	for rev, cl := range cis.ByRevision() {
		fmt.Printf("\t\t%s\n", rev)
		for idx, ci := range cl {
			fmt.Printf("\t\t(%d/%d): %s\n", idx+1, len(cl), ci.File)
		}
	}
	fmt.Printf("---\n")
}

func listStatuses(c *client.Client) {
	result, err := c.ListStatuses()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	for _, status := range result {
		fmt.Printf("%d: %s\n", status.Code, status.Name)
	}
}

func listTags(c *client.Client) {
	result, err := c.ListTags()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	for _, status := range result {
		fmt.Printf("%d\t%25s\t%d\n", status.ID, status.Name, status.UseCount)
	}
}

func loadCase(c *client.Client, id int) {
	bug, err := c.SearchByID(id)
	if err != nil {
		log.Fatalf("%v", err)
	}
	displayBug(bug)
}

func loadCases(c *client.Client, ids []int) {
	bugs, err := c.SearchByIDs(ids)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, bug := range bugs {
		displayBug(bug)
	}
}

func search(c *client.Client, query string, limit int) {
	result, err := c.Search(query, limit)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, bug := range result.Cases {
		displayBug(bug)
	}
}
