# Coinop - A pay-as-you-go helper service for the stellar network.

The notion of a ubiquitous digital currency seems to trigger visions of a world where everything is metered; Pay as you go cell data, pay as you go video streaming, etc.  Consumers seem to want it because it offers the ability to try more things for cheaper, whereas producers seem to want it to improve conversion from non-paying to paying customer.  

However, working with micro-payments involves writing a lot of custom software.  There are not a lot of open source solutions out there yet.  Coinop provides an api server and client libraries to make building a prepaid digital service easier.  Specifically, coinop provides facilities to open, close, and track the balances of application-specific accounts programmatically. In addition to manual account changes (made via authenticated api calls or client libraries), deposits can be made into accounts by payments from the stellar payment network.

Imagine, for example, an on-demand electrical service that takes payment from the stellar network.  At signup, a customer of this service would be given a stellar federated account name, for example `cust-1234*power.co`.  They would make an initial payment to provide some balance to their account.  As the customer uses electricity, the control software for the meter periodically makes an API to debit the appropriate customer account.  It may also periodically query an account's balance via the API and shut the power off if the account has been in the negative for too many consecutive checks.

Bullet points:
- Configuration file to define postgres database and horizon server.
- API server
- Audit log
- Database Models: DepositAddress, Account, LogEntry


## Command line examples

```bash

# Manage watched addresses
$ coinop add-deposit-address GSDS...
$ coinop remove-deposit-address GSDS...

# Manage accounts
$ coinop open-account
$ coinop edit-account ACCOUNT_ID
$ coinop excise-account ACCOUNT_ID

# Manage balances
$ coinop change-balances ACCOUNT_ID
```
