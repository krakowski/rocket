This is a small application designed for sending notifications to RocketChat. It was developed with the idea of sending notifications from CI pipelines to channels.

## :wrench: &nbsp; Installation

The binary is available for Linux, macOS and Windows from the [GitHub Releases page](https://github.com/krakowski/rocket).

## :bulb: &nbsp; Example

First a ".rocket" folder must be created in the project directory.

```console
mkdir .rocket
```

In the next step a template for the notification is created in the folder just created.

```console
cat << EOF > .rocket/my-template.yml
channel: '#my-channel'
text: 'Hello from {{.Env.SENDER}}'
EOF
```

Since the YAML file is processed with the [`template`](https://golang.org/pkg/text/template) package, as in this
example, environment variables can be used. The created template can now be used to send a notification.

```console
ROCKET_HOST=https://rocketchat.com ROCKET_USER=myuser ROCKET_PASS=mypass SENDER=Rocket rocket notify my-template
```

The environment variables can of course also be set in advance. This includes variables set in CI Pipelines.

## :triangular_ruler: &nbsp; Configuration Options


The `rocket` tool accepts the following arguments.

|         Name        |   Type   |    Required    | Description                                                         |
|:-------------------:|:--------:|:--------------:|---------------------------------------------------------------------|
|      `TEMPLATE`     | `string` | :black_circle: | The name of the template to be used                                 |
| `--directory`, `-d` | `string` |                | The directory in which to search for templates (default: `.rocket`) |

Additionally some options are configured via environment variables.

|      Name     |   Type   |    Required    | Description                                    |
|:-------------:|:--------:|:--------------:|------------------------------------------------|
| `ROCKET_HOST` | `string` | :black_circle: | The URL of the RocketChat server to connect to |
| `ROCKET_USER` | `string` |                | The RocketChat user's username                 |
| `ROCKET_PASS` | `string` |                | The RocketChat user's password                 |

`ROCKET_USER` and `ROCKET_PASS` are not required, since they are interactively queried from the console when not available.

Please refer to the [RocketChat documentation](https://docs.rocket.chat/api/rest-api/methods/chat/postmessage#payload)
to find out which fields can be set in the template file(s).

## :scroll: &nbsp; License

This project is licensed under the GNU GPLv3 License - see the [LICENSE](LICENSE) file for details.