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

## Usage

### Push a new task to your inbox

	wlcli push remember the milk

creates a new task with title "remember the milk" in your inbix list.
