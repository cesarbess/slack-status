# slack-status

**slack-status** updates your status across multiple teams.

## Usage

To use slack-status, you must first obtain personal tokens for each of your
teams. Those tokens can be acquired in the [Legacy Tokens](https://api.slack.com/custom-integrations/legacy-tokens)
section of the Slack API website.

After obtaning token for your teams, create a new `~/.slack-status` file with
the following contents:

```yaml
token:
  Team-Name: 'xoxp-your-very-secret-slack-token'

default: ':wave: Hello!'
```

The structure above is the minimum required to use `slack-status`. Please do
notice that:

1. `Team-Name` don't need to be the actual name of your Slack team. It is just
used to identify a team among others, since the tool supports more than one
team.
2. The `default` key is special. The tool will use this key in case no key is
provided.
3. Status keys may contain an Array. In this case, the tool will randomly pick a
status and set it.
4. The `:emoji-name: text` format is required. Although you must always provide
an `:emoji:`, the `text` part is optional.
5. Always wrap the `:emoji: text` within quotes or double-quotes.

For the following examples, assume the following configuration:

```yaml
token:
  Team-Name: 'xoxp-your-very-secret-slack-token'

default: ':wave: Hello!'
lunch:
    - ':hamburger: Lunch time!'
    - ':hamburger: Lunching'
    - ':hamburger: BRB! Lunching!'
```

### Setting the default state
If `slack-status` is invoked without arguments, it will look-up for the
`default` group and set it:

```
$ slack-status
Set on Team-Name: (wave) Hello!

$
```

### Setting a status from a named group
```bash
$ slack-status lunch
Set on Team-Name: (hamburguer) Lunching!

$ slack-status lunch
Set on Team-Name: (hamburguer) Lunch time!
```

### Unsetting your status on all teams
```bash
$ slack-status unset
Cleared your status on Team-Name.
```

### Listing available statuses
```bash
$ slack-status list
default: (wave) Hello!

lunch
  - (hamburger) Lunch time!
  - (hamburger) Lunching
  - (hamburger) BRB! Lunching!
```


## Contributing
Bug reports and pull requests are welcome on GitHub at
https://github.com/victorgama/slack-status. This project is intended to be a safe,
welcoming space for collaboration, and contributors are expected to adhere to
the [Contributor Covenant](http://contributor-covenant.org) code of conduct.


## Code of Conduct
Everyone interacting in the slack-status project’s codebases, issue trackers, chat
rooms and mailing lists is expected to follow the
[code of conduct](https://github.com/victorgama/slack-status/blob/master/CODE_OF_CONDUCT.md).


## Acknowledgements
slack-status was built using the following awesome Open Source projects:

| Name                                                                | License      |
|---------------------------------------------------------------------|--------------|
| [fatih/color](https://github.com/fatih/color)                       | MIT          |
| [go-yaml/yaml](https://github.com/go-yaml/yaml)                     | Apache 2.0   |
| [hashicorp/go-syslog](https://github.com/hashicorp/go-syslog)       | MIT          |
| [urfave/cli](https://github.com/urfave/cli)                         | MIT          |
| [levigross/grequests](https://github.com/levigross/grequests)       | Apache 2.0   |


## License

```
The MIT License (MIT)

Copyright (c) 2019 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```