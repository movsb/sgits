# SGITS

A **S**imple **GIT** **S**erver, using only your local GIT installation.

## Example Configuration File

Create a YAML file named `sgits.yml` in the working directory of SGITS.

For example:

```yml
# port to be listened
listen: :3558

# root directory of your projects
root: /home/tao/code

# username to push to remote
username: name

# password for username
password: pass
```

Now, SGITS will listen on `:3558` as a GIT server.

## Accessing the GIT server

```bash
$ git clone http://localhost:3558/repo

$ git push origin master
```

**Note:** No suffix `.git` required.

## Security

For security, you can set `username` and `password` in your configuration file.

If username and password are set:

- Authenticated Read
- Authenticated Write

If username and password are NOT set:

- Anonymous Read
- Authenticated Write
