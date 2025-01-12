# KeyDash
A sleek, quick-access tool for KeyVault secrets.

No more scrolling through Keyvaults (**WITHOUT A SEARCH BAR**) to find the secret you need.

Just add it to your PATH and you should be good to go.

## WHERE'S THE EXE???
See latest workflow run for the latest release (under artifacts) [like this.](https://github.com/jibuene/KeyDash/actions/runs/12330361074)

You need to give execute permission on this file to run it.

Then you just `./keydash` and you're good to go.

## How to install?
1. Download the latest release from the actions tab.
2. Unzip the archive.
3. CD into the directory.
4. Run install.sh.
5. Run keydash in you terminal.
6. ???
7. Profit.


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

