# Go-Shodan-API
# Go-Shodan-API: Go Script for Shodan API Interaction

Go-Shodan-API is a GoLang script that allows you to interact with the Shodan API, enabling you to query and retrieve internet-wide data from Shodan. This script simplifies the process of accessing Shodan's extensive database of network information using the Go programming language.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Authentication](#authentication)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- Perform various queries on Shodan, including IP, domain, and port searches.
- Retrieve and parse search results in a structured format.
- Easily integrate Shodan data into your Go applications for network analysis and security monitoring.

## Prerequisites

Before using this script, you need to have the following:

- [GoLang](https://golang.org/dl/) installed on your system.

## Installation

1. Clone the GoShodan repository to your local machine:

   ```shell
   git clone https://github.com/Scofield-777/Go-Shodan-API
   ```
## Usage
-To use the script, follow these steps:
  --Export the shodan API keys as environment varaible.
  --Run the script with the desired commands and parameters as per shodan documentation. For example, to search devices running specific service:
  ```sh
    go run main.go SMB
  ```
  -- IP Search 
  ```sh
    go run main.go "example.com"
  ```
  
## Authentication
- A Shodan API account with API key. You can obtain your API key by signing up for an API account on the [Shodan website](https://account.shodan.io/).
- Ensure that you export the Shodan API key as environment variable declared in the script for proper authentication.

## Contributing
- Contributions to Go-Shodan-API are welcome! If you encounter issues, have feature requests, or want to contribute code, please open an issue or create a pull request on the GitHub repository.
