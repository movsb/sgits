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
```

Now, SGITS will listen on `:3558` as a GIT server.

## Accessing the GIT server

```bash
$ git clone http://localhost:3558/repo

$ git push origin master
```
