## Repoguard
Repoguard is a CLI tool built in Go to get information about who has access to your GitHub Repository and revoke the access of a user or multiple users if required.

### How to Install
- Go to the [Release](https://github.com/harisheoran/repoguard/releases) section
- Download for your OS
- Linux
    - Unzip the tar and move binary to /usr/local/bin/
    ```
    repoguard rg --u=<username> --r=<repo>
    ```

### How to use Repoguard?
- First, provide a GitHub API token
```
export TOKEN=<your token here>
```
- ```rg``` command to get a list of collaborators.
- Use flag --u for providing the username of GitHub
- Use flag --r for providing the GitHub repository name

### Features
#### Implemented
- Get a list of users who have access to the GitHub Repository.

#### To be implemented
- Revoke access of single or multiple users.
- Generate a file of the list of users who have access to GitHub Repo.

#### Thanks to
- [Cobra](https://cobra.dev/)
- [go-github](https://github.com/google/go-github)

