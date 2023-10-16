# Unfriended
Who Unfriended you on IG??

## Social Media Follower Analysis
This Golang script is designed to analyze your social media followers and following to identify users you are following but who are not following you back. It takes as input two JSON files, `followers.json` and `following.json`, which should contain information about your followers and the users you are following on a social media platform. 

## Prerequisites

- Go
- To use this script, you will need to have two JSON files: `followers.json` and `following.json`. You can get these files by exporting your followers and following lists from Instagram Accounts Centre.
- IDE (VS Code, GitHub Codespaces)
  
## How to Use

1. Make sure you have the required JSON files in the same directory as this script.
2. Run the script using go.

```bash
go run main.go
```
## Understanding the Code

- The code reads the `followers.json` and `following.json` files.
- It extracts relevant data from these files.
- It creates a list of users you are following but who are not following you back.
- Finally, it prints the usernames of users who fall into this category.

## Customization

You may need to modify the code to match the structure of your JSON files if they are organized differently. Specifically, make sure to update the keys and values used to access usernames in the JSON data.

## _Notes:_

> This script does not use the Instagram API, - so you do not need to have an Instagram - developer account to use it.
>
> The script is still under development, so there may be some bugs.
>
> Please use this script responsibly.

## License

See the [LICENSE](LICENSE) file for license rights and limitations (MIT).
