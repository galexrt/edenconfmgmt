# edenconfmgmt

Configuration management with automatic clustering, events and stuff.

Inspired by [mgmt](https://github.com/purpleidea/mgmt), [Ansible](https://www.ansible.com/) and [Salt](https://www.saltstack.com/).

## Why?

("I" and "me" are standing for myself, @galexrt)
I was not statisfied with Salt's approach to being able to access "data" of "previous" states like in Ansbile where one would simply put `register: VARNAME` and be able to access the data.
Maybe I was just doing it wrong. Who knows?
On the other hand, I like Ansible's simplicity but totally missed the event driven part in it. At least I am more of a "I want an agent on the nodes" type of person, so the agent is able to react when e.g., there is no network connectivity anymore.
After some interweb searching, I found `mgmt`. though for me some examples didn't work. I was more and more fedup with configuration management tools in general, so I didn't bother joining the IRC..

### What do I want to solve with this?

Looking at least at Salt here, I wanted to achieve the following tasks:
* Copy files from minion to minion.
    - Note: Yes there are ways to do that, either using `publish`ers or `cp.push` but the files (certificates, private keys, etc.) I wanted to transfer must have been securely transfered.
* Ansible like state/task variable sharing.
    - "Task to variable", so it can be used in the next state/template to determine certain things.
* And some more or less smaller things that just didn't make me love Salt and Ansible for what I wanted to do/use it for.

## Goal

This project will kind of go back to the roots and in the first "phase" simply allowing executing commands/scripts on machines from a central point.
But allow extension using a [Terraform](https://github.com/hashicorp/terraform) like plugin system.

I (personally) don't want hundreds of integrations in this repo. You, the user, the community should share their plugins openly to help each other.
Let me be honest here, for the first part this project should solve my problems, if you see this project help you solve your problems I'd be happy to extend the functionality/feature set(s).

The following list contains the primary goals of this project:

1. Throw an event for "everything":
    * Black/Whitelist for events thrown.
1. Simple input webhook event API:
    * Allowing e.g., Prometheus webhooks to be accepted and an annotation/a label to be used to run a certain task.
1. React(or)ion to events:
    * Allow doing things based on events.
        - [Salt's Thorium Reactor](https://docs.saltstack.com/en/latest/topics/thorium/index.html) like functionality.
1. Allow sharing files between nodes:
    * Not too big files though.
        - If the file is bigger than X MB, it is not a "configuration part" that is shared here ;).
            E.g., if one shares binaries using this mechanism, "you are doing it wrong". Binaries should be "shared" through packages.
            Only thing I want (to hear is) shared through this, are config snippets, certificates, public/private keys and other such things.
1. More to come.

### Example Use Cases

1. _Server A_ generates certificates (and their respective private keys) which are needed on the other `master` "labeled" servers. The other `master` labeled servers "request" the file(s) from server A. _Server C_ can't request the files from _server A_ nor the other `master` labeled servers (this basically boils to a dynamic ACL system).
1. _Server B_ joins the configuration management cluster, an event is thrown and _server A_ generates a token which in return triggers a task to be run on _server B_ which consumes the token generated.

## Designs

The design of components and most workflows can be found in the [design/](/design/) directory.
