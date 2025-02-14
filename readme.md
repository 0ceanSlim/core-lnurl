# Core LNURL Server - CLN REST Integration

## 📌 Overview

This is a Go-based LNURL server that allows users to generate Lightning Network invoices via Core Lightning (CLN) REST API using LNURL-Pay.

This assumes you already have:  
✅ A Bitcoin full node running and synced.  
✅ A Core Lightning (CLN) node running with at least one active channel.  
✅ The CLN REST server enabled and working.  
✅ The ability to use either use `lightning-cli` or your CLN-Rest Server to issue a new rune with restrictions to only use the invoice method

### 🚀 Features

✅ Dynamic Invoice Generation → Users get a unique invoice every time.  
✅ Core Lightning REST API Integration  
✅ Configurable Settings → Easily modify ports, CLN node URL, and authentication Runes.  
✅ LNURL-Compatible → Works with LNURL-enabled wallets for seamless payments.

### 📡 API Endpoints

**LNURL-Pay** (`/lnurl/pay`)  
Request an invoice for a user:

```bash
curl "http://localhost:8085/lnurl/pay?username=user1&amount=1000000"
```

🔹 Example Response:

```bash
{
  "pr": "lnbc10u1pn6uegtsp..."
}
```

You can use any username and this server will fetch an invoice from your node for a payment in the amount requested. I'm hoping this makes self custodial lightning addresses on CLN a breeze.

Basically start this server with the proper configs and use any lightning address you want to this server and it will make an invoice that pays your node.

### Issue a rune for this server to use

For security, you should issue a new rune for this server to use that has restriction so it can only use the `invoice` command for your node.

There are a few ways of doing this.

- `lightning-cli`
  - Use the cli to issue a new rune with restrictions
  ```bash
  lightning-cli createrune -k "restrictions"='[["method=invoice"]]'
  ```
- You can visit the front end of your REST API (`127.0.0.1:3010` is common) and use the gui to log in with your master rune (`lightning-cli showrunes` to show all your runes). After you are logged in, you can use the gui to make a call to your rest server to create a new rune
  - Under the green POST Dropdown enter
    - **rpc_method :** `createrune`
    - **payload :** `{"restrictions": [["method=invoice"]]}`
- Lastly, you can make a curl request from a terminal (this does the same thing as your gui does above)
  ```bash
  curl -X 'POST' \
  'http://127.0.0.1:3010/v1/createrune' \
  -H 'accept: application/json' \
  -H 'Rune: YOUR_MASTER_RUNE_HERE' \
  -H 'Content-Type: application/json' \
  -d '{"restrictions": [["method=invoice"]]}'
  ```

In all of these cases you should get a response with the rune and it's summary

```bash
{
  "rune": "THIS_IS_YOUR_INVOICE_RUNE",
  "unique_id": "1"
}
```

### Installation

- Make a directory to put the config and program into
- Download the latest release binary for your system from the [releases page](https://github.com/0ceanSlim/core-lnurl/releases/latest)
- Create a `config.yml` in the same directory as your executable. There is an example [here](https://github.com/0ceanSlim/core-lnurl/blob/main/example-config.yml)
- Run the server `./core-lnurl`

#### Set up a system service on linux

- Guide and example service file coming soon if you don't already know how to do this

#### Use nssm to make a Windows service

- Guide at some point

#### Docker container by others

#### 🎯 Summary

🚀 Core LNURL Server lets you create Lightning invoices via CLN REST API.  
⚡ Works with LNURL-Pay wallets for easy payments.  
🔧 Fully configurable and self-hosted!

#### 📝 License

MIT License – Feel free to modify and contribute! 🚀

Open Source and made with 💦 by [OceanSlim](https://njump.me/npub1zmc6qyqdfnllhnzzxr5wpepfpnzcf8q6m3jdveflmgruqvd3qa9sjv7f60)
