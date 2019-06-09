# SGITS

A **S**imple **GIT** **S**erver, using only your local GIT installation.

## Example Configuration File

Create a YAML file named `sgits.yml` in the working directory of SGITS.

For example:

```yml
# http listen address
listen: :3558

# root directory of your projects
root: /home/tao/code

# username to git server
username: name

# password for username
password: pass
```

Now, SGITS will listen on `:3558` as a GIT server.

## Accessing the GIT server

Because GIT server requires bare repositories (without working directory), you should first create it before pushing:

```sh
# at projects root
$ git init --bare repo
```

Now, you can push and clone:

```sh
# first add as remote
$ git remote add sgits http://localhost:3558/repo

# push to remote
$ git push -u sgits master
```

**Note:** No suffix `.git` is required.

## Security

For security, you can set `username` and `password` in your configuration file.

If username and password are set:

- Authenticated Read
- Authenticated Write

If username and password are NOT set:

- Anonymous Read
- Authenticated Write
