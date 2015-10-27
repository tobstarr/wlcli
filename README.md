# wlcli

Wunderlist command line client.

## Configuration

Place a file at `$HOME/.wlcli/config.json` with this content:

	{
		"client_id": "OATH_APP_ACCESS_CLIENT_ID",
		"access_token": "OAUTH_APP_ACCESS_TOKEN"
	}

You can get this information by going to https://developer.wunderlist.com/apps

* Create a new Application
* Fill in some dummy URLs at `APP URL` and `Auth Callback URL`
* click `Create Access Token`

## Installation

	go get -v github.com/tobstarr/wlcli`

## Usage

	wlcli inbox list                    List Inbox
	wlcli inbox push       <Payload>... Push a task to inbox
	wlcli tasks complete   <IDs>...     Complete Tasks
	wlcli tasks delete     <IDs>...     Delete Tasks

### the current list

You can specify a default list_id on a per directory basis:

	$ cat .wlcli
	list_id = 12345

In that case you can just run `wlcli list` to list that very list and `wlcli push` to create new tasks in it.

### useful shell aliases

	alias wl="wlcli $@"
	alias inbox="wlcli inbox $@"

### push a new task to your inbox

	wlcli inbox push remember the milk

creates a new task with title "remember the milk" in your inbix list.

### list inbox

	wlcli inbox list

lists all tasks in the inbox list

## to come

* "dispatch inbox": move tasks from inbox into other lists
