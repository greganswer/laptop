# laptop

This repo contains assets that are used to setup a laptop for web development. This includes things like:

- dotfiles
- App settings
- Setup script
- Workflow script

## Tech Stack

Just Bash for now! I wanted something simple and portable. Try setting up a new laptop with a different version of Golang than the one used to create your setup script and you'll understand 😁.

## Installation

Quick and dirty installation (as of 0.3.0):

    mkdir -p ~/src
    cd ~/src
    curl https://raw.githubusercontent.com/greganswer/laptop/master/bin/setup --output temp_laptop_setup
    chmod +x temp_laptop_setup
    . ./temp_laptop_setup

## Update

    brew bundle check --file ~/Brewfile
    gca
    gp

## License

[MIT](https://choosealicense.com/licenses/mit/)
