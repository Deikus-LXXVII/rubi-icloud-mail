// Command icloud-mail is the iCloud Mail plugin for Rubi: read, search and draft freely (by default),
// send only after the user's approval, and get notified when someone replies to a tracked email.
//
// Apple offers no OAuth for iCloud Mail; third-party apps use an app-specific password over IMAP/SMTP.
// The user enters it in the Rubi panel, never in the agent chat, and Rubi keeps it encrypted.
package main

import (
	"github.com/Deikus-LXXVII/rubi/sdk/rubiplugin"
)

const ID = "icloud-mail"

const (
	kindRead    = ID + ".read"
	kindDraft   = ID + ".draft"
	kindSend    = ID + ".send"
	kindWatch   = ID + ".watch"
	kindPrivate = ID + ".private"
	kindFolder  = ID + ".folder"
)

// PublisherKey is the Rubi-Project plugin signing key (public half).
const PublisherKey = "MCowBQYDK2VwAyEAxeDfKAkO77JdARN7Y2jJT3tXw9mN+GqqH8R5mhcxt8c="

var sendOptions = []rubiplugin.Option{
	{Key: "send", Label: "Send"},
	{Key: "send_track", Label: "Send and notify on reply", Meaning: "send, then watch for replies and notify you"},
}

func manifest() rubiplugin.Manifest {
	return rubiplugin.Manifest{
		ID:          ID,
		Name:        "iCloud Mail",
		Version:     "v1.0.0",
		Description: "Read, search and draft iCloud email; send after your approval; get notified about replies.",
		Needs: "Your iCloud email address and an app-specific password created at account.apple.com " +
			"(Sign-In and Security > App-Specific Passwords). Apple offers no other way for apps to access iCloud Mail.",
		Publisher: rubiplugin.Publisher{Name: "Rubi-Project", Key: PublisherKey,
			URL: "https://github.com/Deikus-LXXVII/rubi-icloud-mail"},
		Source:  "https://github.com/Deikus-LXXVII/rubi-icloud-mail",
		MinRubi: "v0.4.0",
		Entry:   "icloud-mail",
		Fields: []rubiplugin.Field{
			{Key: "address", Label: "iCloud email address", Type: "email", Placeholder: "name@icloud.com", Required: true,
				Help: "Use your @icloud.com (or @me.com) address, even if your Apple ID uses another email."},
			{Key: "from_name", Label: "Your name (shown to recipients)", Type: "text", Placeholder: "Optional"},
		},
		Secrets: []rubiplugin.Secret{{
			Key:     "app_password",
			Label:   "App-specific password",
			Help:    "On account.apple.com open Sign-In and Security, then App-Specific Passwords, create one named \"Rubi\", and paste it here. You can revoke it there at any time.",
			HelpURL: "https://account.apple.com/account/manage",
		}},
		Actions: []rubiplugin.Action{
			{Kind: kindRead, Title: "Read and search mail", DefaultLevel: rubiplugin.None},
			{Kind: kindDraft, Title: "Save drafts", DefaultLevel: rubiplugin.None},
			{Kind: kindSend, Title: "Send email", DefaultLevel: rubiplugin.Strong, Options: sendOptions},
			{Kind: kindWatch, Title: "Watch for new mail (wakes your agent)", DefaultLevel: rubiplugin.None},
			{Kind: kindPrivate, Title: "Show a private email", DefaultLevel: rubiplugin.Strong, Locked: true,
				Options: []rubiplugin.Option{{Key: "show", Label: "Show it to my agent"}}},
			{Kind: kindFolder, Title: "Open a closed folder for a while", DefaultLevel: rubiplugin.Strong, Locked: true},
		},
		Events: []rubiplugin.EventType{
			{Type: "reply", Untrusted: []string{"reply.from", "reply.subject"}},
			{Type: "watch", Untrusted: []string{"message.from", "message.subject", "message.snippet"}},
		},
		Config: append(append([]rubiplugin.ConfigField{}, folderSettings...),
			rubiplugin.ConfigField{Key: "hide_codes", Label: "Hide sign-in codes, one-time passwords and confirmation links", Type: "bool", Default: true,
				Help: "Built-in list, English and Russian, plus subjects like \"482913 is your code\"."},
			rubiplugin.ConfigField{Key: "hide_password_resets", Label: "Hide password reset emails", Type: "bool", Default: true},
			rubiplugin.ConfigField{Key: "hide_sign_in_alerts", Label: "Hide sign-in and security alerts", Type: "bool", Default: false,
				Help: "New sign-ins, new devices, suspicious activity. Off by default, so your agent can warn you; an alert that contains a code stays hidden by the first switch."},
			rubiplugin.ConfigField{Key: "hidden_senders", Label: "Hidden senders", Type: "list", Default: []string{},
				Help: "Email addresses or domains, one per line (e.g. bank.com). Your agent can't see mail from them."},
			rubiplugin.ConfigField{Key: "hidden_keywords", Label: "Hidden words", Type: "list", Default: []string{},
				Help: "Words or phrases, one per line. Your agent can't see mail whose subject or text contains one."},
		),
		Egress: []string{"imap.mail.me.com:993", "smtp.mail.me.com:587"},
	}
}

func main() { newPlugin(&integration{}).Main() }
