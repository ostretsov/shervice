# shervice
shervice makes any shell command as a RESTful HTTP service in a minute.

### Configuration

Here is a simple config of a service for `echo` command:

```yaml
services:
  - url: /echo
    args:
      - name: message
    command: /bin/echo %message%
```

* raw
* type