# iCloud Mail for Rubi

A [Rubi](https://github.com/Deikus-LXXVII/rubi) plugin that lets your AI agent work with iCloud Mail, with
your approval:

- read and search mail (never marks anything as read);
- save drafts to your iCloud Drafts folder;
- send email only after you approve it in the Rubi panel (with your passkey by default);
- optionally get notified when someone replies to an email you sent;
- let your agent watch for new mail from given senders or with given words (say, delivery updates for an
  order) and react to each one;
- keep private mail private: sign-in codes, password resets and anything you choose stay hidden from your
  agent unless you approve showing one;
- choose which folders your agent can see, and whether it may ask for the others;
- connect several iCloud Mail accounts. Your agent picks one by its address; otherwise it uses the default
  (the first one you connected).

## Install

Ask your agent: "Connect my iCloud Mail through Rubi." It installs the plugin from the Rubi store (you
approve what it can do), then sends you a setup link. There you enter your @icloud.com address and an
app-specific password from [account.apple.com](https://account.apple.com/account/manage) (*Sign-In and
Security* > *App-Specific Passwords*). Apple offers no other way for apps to reach iCloud Mail. Rubi keeps
the password encrypted; your agent never sees it.

To add another account, open *Settings* in the Rubi panel and press *Add account* at iCloud Mail.

## Permissions

| | |
|---|---|
| Asks for | iCloud email address, app-specific password |
| Read and search mail | no approval by default |
| Save drafts | no approval by default |
| Send email | passkey / password by default |
| Connects to | `imap.mail.me.com:993`, `smtp.mail.me.com:587` |
| Watch for new mail (wakes your agent) | no approval by default |
| Show a private email | passkey / password, always |
| Give an attachment to your agent (when set to "may ask") | passkey / password, always |
| Open a closed folder for a while | passkey / password, always |
| Can notify your agent about | replies to tracked emails, mail matching its watches |

You can change the approval levels in Rubi's settings, except the two marked "always".

## Your settings

Right after connecting, and later in the Rubi panel under *Settings*, at each account of iCloud Mail:

- **Folders your agent can see** (for each account): all, or only the ones you check; and whether it may
  ask for the others.
- **Hide sign-in codes, one-time passwords and confirmation links:** on by default.
- **Hide password reset emails:** on by default.
- **Hide sign-in and security alerts:** off by default, so your agent can warn you about new sign-ins or
  suspicious activity. (An alert that contains a code stays hidden by the first switch.)
- **Hidden senders** and **hidden words:** anything else you want to keep from your agent.

Formatted emails: your agent reads emails as text, and can ask for the formatted version when it needs a
button the text doesn't show (an "Unsubscribe" link, say). It gets the links and a copy of the email that
loads nothing from the internet, so opening it doesn't tell the sender you read it.

Attachments: per mailbox, choose whether your agent can open the documents and photos attached to emails,
may ask you first (the default), or can't open them at all and sees only how many there are.

The privacy switches apply to all your iCloud Mail accounts. Only you can change these settings, with your
passkey or password; your agent can't see or change them. Hidden mail shows up to your agent only as
"a private email from <sender>".

## Development

The mail engine lives in [rubi-mailkit](https://github.com/Deikus-LXXVII/rubi-mailkit), shared with the
Gmail plugin; this repository only describes iCloud Mail (`main.go`). Tests are in rubi-mailkit.

```bash
go build -o icloud-mail . && ./icloud-mail --manifest
```

To develop against local checkouts, use a workspace (not committed):

```bash
go work init . ../rubi-mailkit ../rubi
```

## Releasing

Tag `vX.Y.Z`. The release workflow tests, builds for linux/amd64, linux/arm64 and darwin/arm64, writes
`rubi-plugin.json`, signs `SHA256SUMS` with the plugin key (`PLUGIN_SIGNING_KEY`; public half in
`publisher-key.pub.pem`), attests the build, and publishes the release. It prints the SHA-256 of
`SHA256SUMS` for the Rubi store catalog.

## License

[MIT](LICENSE)
