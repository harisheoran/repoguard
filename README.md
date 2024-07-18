## Repoguard
Repoguard is a CLI tool built in Go to get information about who has access to your GitHub Repository and revoke the access of a user or multiple users if required.

### How to Install
- Go to the [Release](https://github.com/harisheoran/repoguard/releases) section
- Download for your OS

```
wget <file link from release section for your OS>
tar -xzvf <downloaded tar file>
mv repoguard /usr/local/bin
```

### How to use Repoguard?
- Configure the CLI
``` repoguard configure ```

and provide it the GitHub Token

- ```users``` command to get a list of collaborators.
    - Use flag ```--r``` for providing the GitHub repository name.

- ```revoke``` to revoke access of a user.
    - Use flag ```--r``` for providing the repository name of GitHub account of that owner.
    - Use flag ```--u``` for providing the username whom access you want to revoke.

### Features
#### Implemented
- Get a list of users who have access to the GitHub Repository.
- Revoke access of user.
- Configure repoguard for your GitHub account.

#### Thanks to
- [Cobra](https://cobra.dev/)
- [go-github](https://github.com/google/go-github)

