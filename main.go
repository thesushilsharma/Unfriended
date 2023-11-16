package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Relationship struct {
	StringListData []struct {
		Value string `json:"value"`
	} `json:"string_list_data"`
}

func main() {
	// Read 'followers.json' and 'following.json' files
	followersData, err := os.ReadFile("followers.json")
	if err != nil {
		fmt.Println("Error reading followers.json:", err)
		return
	}

	followingData, err := os.ReadFile("following.json")
	if err != nil {
		fmt.Println("Error reading following.json:", err)
		return
	}

	// Define a map to store users you are following
	followingMap := make(map[string]bool)

	// Unmarshal JSON data for 'following.json'
	var followingJSON struct {
		RelationshipsFollowing []Relationship `json:"relationships_following"`
	}

	if err := json.Unmarshal(followingData, &followingJSON); err != nil {
		fmt.Println("Error unmarshaling following.json:", err)
		return
	}

	// Unmarshal JSON data for 'followers.json'
	var followersJSON []Relationship

	if err := json.Unmarshal(followersData, &followersJSON); err != nil {
		fmt.Println("Error unmarshaling followers.json:", err)
		return
	}

	// Extract values from followers.json
	unfollowers := make([]string, 0)

	for _, following := range followingJSON.RelationshipsFollowing {
		if len(following.StringListData) > 0 {
			username := following.StringListData[0].Value
			followingMap[username] = true
		}
	}

	// Identify users you are following but not followed back
	for _, follower := range followersJSON {
		if len(follower.StringListData) > 0 {
			username := follower.StringListData[0].Value
			if followingMap[username] {
				delete(followingMap, username)
			}
		}
	}

	// Populate the list of unfollowers
	for user := range followingMap {
		unfollowers = append(unfollowers, user)
	}

	// Print users you are following but not followed back
	fmt.Println("Users you are following but not followed back:")
	for _, user := range unfollowers {
		fmt.Println(user)
	}
}
