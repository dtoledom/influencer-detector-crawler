package models

import (
	"cassandra"
	"log"
)

// Node defines a model to represent an node in the social graph.
type Node struct {
	// Page info
	ID          string
	Platform    string
	Name        string
	Description string

	// Relevance
	Categories []string
	FanCount   int64

	// Location
	City    string
	Country string
	ZIP     string
}

// Save the Node to the given database.
func (n *Node) Save() {
	session := cassandra.GetSession()

	// id, platform, name, description, categories, fan_count, city, country, zip
	q := session.Query(cassandra.InsertEdgeQuery,
		n.ID,
		n.Platform,
		n.Name,
		n.Description,
		n.Categories,
		n.FanCount,
		n.City,
		n.Country,
		n.ZIP,
	)

	if err := q.Exec(); err != nil {
		log.Fatal(err)
	}
}
