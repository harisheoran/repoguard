## Repoguard
Repoguard is a CLI tool built in Go to get information about who has access to your GitHub Repository and revoke the access of a user or multiple users if required. 

### How to use Repoguard?
- First, provide a GitHub API token
```
export TOKEN=<your token here>
```
- ```rg``` command to get a list of collaborators.
- Use flag --u for providing the username of GitHub
- Use flag --r for providing the GitHub repository name

example use
```
rg --u=harisheoran --r=sparrowbit
```

### Features
#### Implemented
- Get a list of users who have access to the GitHub Repository.

#### To be implemented
- Revoke access of single or multiple users.
- Generate a file of the list of users who have access to GitHub Repo.

#### Thanks to
- [Cobra](https://cobra.dev/)
- [go-github](https://github.com/google/go-github)