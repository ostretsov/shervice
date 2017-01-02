# shervice
shervice makes any shell command as a RESTful HTTP service in a minute.

### Configuration

Here is a simple config of a service for `echo` command:

```yaml
services:
    - url: /echo
      arguments:
        - [name: message, type: string, raw: false]
      command: '/bin/echo ' ~ message
```