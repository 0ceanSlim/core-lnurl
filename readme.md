# Core LNURL Server - CLN REST Integration

## 📌 Overview

This is a Go-based LNURL server that allows users to generate Lightning Network invoices via Core Lightning (CLN) REST API using LNURL-Pay.

This assumes you already have:
✅ A Bitcoin full node running and synced.
✅ A Core Lightning (CLN) node running with at least one active channel.
✅ The CLN REST server enabled and working.

### 🚀 Features

✅ Dynamic Invoice Generation → Users get a unique invoice every time.
✅ Core Lightning REST API Integration → No need for lightning-cli or SSH.
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

As you can see, you can put usernames in the lnurlp folder and use the lnurl server to fetch an invoice for a payment someone wants to send a user. I'm hoping this makes self custodial lightning addresses on CLN a breeze

#### 🎯 Summary

🚀 Core LNURL Server lets you create Lightning invoices via CLN REST API.
⚡ Works with LNURL-Pay wallets for easy payments.
🔧 Fully configurable and self-hosted!

#### 📝 License

MIT License – Feel free to modify and contribute! 🚀

Open Source and made with 💦 by [OceanSlim](https://njump.me/npub1zmc6qyqdfnllhnzzxr5wpepfpnzcf8q6m3jdveflmgruqvd3qa9sjv7f60)
