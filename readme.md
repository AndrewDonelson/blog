# blog
**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/blog@latest! | sudo bash
```
`username/blog` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## serve

Cosmos SDK's version is: stargate - v0.46.6

🛠  Building proto...
📦 Installing dependencies...
🛠  Building the blockchain...
💿 Initializing the app...
🗂  Initialize accounts...
🙂 Added account alice with address blog1nad8tgwyr0pjprzgr3hjtky9a6fdsn7vul37sq and mnemonic: during govern gadget absent submit joke raccoon deal seminar divide pole claw foam grit style laptop monster zone clog yellow harvest slam marble basic
🙂 Added account bob with address blog16st0y37un0euqvswja7pjhv4l5djmzsrwphfx8 and mnemonic: disease license card party dry dog sport twice aim early change cereal combine master grant erupt raven neglect worth eternal unveil time allow van
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500

## create sample post

```
- address: blog17jh0y4m0pkwp987m5sn46tnnp3a5nu6zzk5z70
  name: andrew
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AoJgIXKz7yDp9EUkBgjhP97hMnMVBueXehDD3O3CMGUm"}'
  type: local
```
blogd tx blog create-post foo bar --from alice


## Commands

## Create Image Endpoint/Message

```
ignite scaffold message createImage ipfsurl tags views likes dislikes
```

Output:

        modify proto/blog/blog/tx.proto
        modify x/blog/client/cli/tx.go
        create x/blog/client/cli/tx_create_image.go
        create x/blog/keeper/msg_server_create_image.go
        modify x/blog/module_simulation.go
        create x/blog/simulation/create_image.go
        modify x/blog/types/codec.go
        create x/blog/types/message_create_image.go
        create x/blog/types/message_create_image_test.go

        🎉 Created a message `createImage`.