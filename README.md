# KeyDash
A sleek, quick-access tool for KeyVault secrets.

No more scrolling through Keyvaults (**WITHOUT A SEARCH BAR**) to find the secret you need.

Just add it to your PATH and you should be good to go.

## WHERE'S THE EXE???
See latest workflow run for the latest release (under artifacts) [like this.](https://github.com/jibuene/KeyDash/actions/runs/12732060185)

You need to give execute permission on this file to run it.

Then you just `./keydash` and you're good to go.

## How to install (Only for linux)?
1. Download the latest artifact from the actions tab.
2. Unzip the archive.
3. CD into the directory.
4. Unzip the keydash.zip file.
5. Run install.sh.
6. Run keydash in your terminal.
7. ???
8. Profit.

## How to install (on MacOS)
1. Download the latest ARM64 artifact from the Actions tab
2. Unzip the archive
3. `cd` into the directory
4. Run `install.sh` (with `sudo`)
5. Run `keydash` in the terminal
   1. The execution will be blocked, because the `keydash` binary is not verified by Apple
6. Manually allow the `keydash` binary
   1. Click `Done` to close the warning dialog
   2. Open System Settings
   3. Select Privacy & Security in the left-hand menu
   4. Scroll down and click `Allow Anyway` on the warning about `keydash` being blocked
7. Run `keydash` again in the terminal
8. In the new warning dialog, click `Open Anyway` and enter local account password
9. Profit!

## Features
Simple searching for secrets by using the start of its name.

Search multiple KeyVaults at once.


## Usage
```bash
keydash help                      # Displays help
keydash --keyvault add mykeyvault # Adds a keyvault to the config file.
keydash --secret mysecret         # Retrieves the secret named 'mysecret'
keydash secret                    # Retrieves the secret named 'secret'
```

