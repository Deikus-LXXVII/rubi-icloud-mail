# iCloud Mail for Rubi

A [Rubi](https://github.com/Deikus-LXXVII/rubi) plugin that lets your AI agent work with iCloud Mail, with
your approval:

- read and search mail (never marks anything as read);
- save drafts to your iCloud Drafts folder;
- send email only after you approve it in the Rubi panel (Face ID by default);
- optionally get notified when someone replies to an email you sent;
- let your agent watch for new mail from given senders or with given words (say, delivery updates for an
  order) and react to each one;
- keep private mail private: sign-in codes, password resets and anything you choose stay hidden from your
  agent unless you approve showing one with Face ID;
- choose which folders your agent can see, and whether it may ask for the others.

## Install

Ask your agent: "Connect my iCloud Mail through Rubi." It installs the plugin from the Rubi store (you
approve what it can do), then sends you a setup link. There you enter your @icloud.com address and an
app-specific password from [account.apple.com](https://account.apple.com/account/manage) (*Sign-In and
Security* > *App-Specific Passwords*). Apple offers no other way for apps to reach iCloud Mail. Rubi keeps
the password encrypted; your agent never sees it.

## Permissions

| | |
|---|---|
| Asks for | iCloud email address, app-specific password |
| Read and search mail | no approval by default |
| Save drafts | no approval by default |
| Send email | Face ID / password by default |
| Connects to | `imap.mail.me.com:993`, `smtp.mail.me.com:587` |
| Watch for new mail (wakes your agent) | no approval by default |
| Show a private email | Face ID / password, always |
| Open a closed folder for a while | Face ID / password, always |
| Can notify your agent about | replies to tracked emails, mail matching its watches |

You can change the approval levels in Rubi's settings, except the two marked "always".

## Your settings

Right after connecting, and later under *Settings* > *iCloud Mail* > *Settings* in the Rubi panel:

- **Folders your agent can see:** all, or only the ones you check; and whether it may ask for the others.
- **Hide sign-in codes and password emails:** on by default (verification codes, one-time passwords,
  password resets, sign-in alerts; English and Russian).
- **Hidden senders** and **hidden words:** anything else you want to keep from your agent.

Only you can change these, with Face ID or your password; your agent can't see or change them. Hidden mail
shows up to your agent only as "a private email from <sender>".

## Development

Built with Rubi's Go SDK (`github.com/Deikus-LXXVII/rubi/sdk/rubiplugin`).

```bash
go test ./...
go build -o icloud-mail . && ./icloud-mail --manifest
```

To develop against a local Rubi checkout, use a workspace (not committed):

```bash
go work init . ../rubi
```

## Releasing

Tag `vX.Y.Z`. The release workflow tests, builds for linux/amd64, linux/arm64 and darwin/arm64, writes
`rubi-plugin.json`, signs `SHA256SUMS` with the plugin key (`PLUGIN_SIGNING_KEY`; public half in
`publisher-key.pub.pem`), attests the build, and publishes the release. It prints the SHA-256 of
`SHA256SUMS` for the Rubi store catalog.

## License

[MIT](LICENSE)
