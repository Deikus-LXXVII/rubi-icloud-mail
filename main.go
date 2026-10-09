// Command icloud-mail is the iCloud Mail plugin for Rubi: read, search and draft freely (by default), send
// only after the user's approval, get notified about replies and watched mail, with a privacy filter and
// folder access the user controls. Several iCloud accounts can be connected.
//
// Apple offers no OAuth for iCloud Mail; third-party apps use an app-specific password over IMAP/SMTP.
// The user enters it in the Rubi panel, never in the agent chat, and Rubi keeps it encrypted.
package main

import (
	"github.com/Deikus-LXXVII/rubi-mailkit"
)

var provider = mailkit.Provider{
	ID:          "icloud-mail",
	Name:        "iCloud Mail",
	Version:     "v2.0.0",
	MinRubi:     "v0.6.1",
	Description: "Read, search and draft iCloud email; send after your approval; get notified about replies. Several accounts.",
	Needs: "Your iCloud email address and an app-specific password created at account.apple.com " +
		"(Sign-In and Security > App-Specific Passwords). Apple offers no other way for apps to access iCloud Mail.",
	Source:     "https://github.com/Deikus-LXXVII/rubi-icloud-mail",
	ToolPrefix: "icloud_mail",

	AddressLabel:       "iCloud email address",
	AddressPlaceholder: "name@icloud.com",
	AddressHelp:        "Use your @icloud.com (or @me.com) address, even if your Apple ID uses another email.",
	PasswordLabel:      "App-specific password",
	PasswordHelp: "On account.apple.com open Sign-In and Security, then App-Specific Passwords, create one named \"Rubi\", " +
		"and paste it here. You can revoke it there at any time.",
	PasswordURL:  "https://account.apple.com/account/manage",
	PasswordLink: "Open account.apple.com",

	IMAPAddr:   "imap.mail.me.com:993",
	SMTPAddr:   "smtp.mail.me.com:587",
	Drafts:     "Drafts",
	Sent:       "Sent Messages",
	MailDomain: "icloud.com",
	AuthError:  "iCloud rejected the login. Check that you used your @icloud.com address and a current app-specific password",
}

func main() { mailkit.Main(provider) }
