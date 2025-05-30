<p align="center">
  <a href="https://www.rest.egeuysal.com/">
    <img src="https://res.cloudinary.com/ddjnqljd8/image/upload/v1748623162/rest-logo.png" height="96">
    <h3 align="center">Spot</h3>
  </a>
</p>

<p align="center">
    Post. Expire. Vanish.
</p>

<p align=center>
    <strong>
        <a href="CONTRIBUTING.md">Contributing</a>
    </strong>
</p>

## Rest 🛠️

**Rest** is a minimal CLI tool for sharing JSON data via one-time, public REST endpoints.  
Built for developers who need to quickly expose or test JSON payloads **without setup or authentication**.

Perfect for:

- Sharing payloads with teammates
- Testing webhooks or API clients
- Temporary data exposure during development

### 🚀 Installation

Install `rest` using the following one-liner:

```sh
curl -fsSL https://raw.githubusercontent.com/egeuysall/rest/master/install.sh | sh
```

### ⚙️ Usage

```sh
rest -d <path-to-json> [-e <expires-in-minutes>] [-t <max-access-count>]
```

#### Options:

| Flag        | Description                                     | Default |
| ----------- | ----------------------------------------------- | ------- |
| `-d` string | **Required.** Path to your JSON file            | —       |
| `-e` int    | Expiration time in minutes                      | `10`    |
| `-t` int    | Max number of times the payload can be accessed | `1`     |

#### Example

```sh
rest -d ./payload.json -e 10 -t 3
```

This shares `payload.json` for **10 minutes** and allows **up to 3 accesses**.

### 🧹 Auto-Expiration

Payloads are automatically deleted after:

- The expiration time (`-e`) is reached
- They have been accessed `-t` times

### 📦 Use Cases

- Share a JSON response with your frontend team
- Provide sample payloads to test third-party integrations
- Simulate webhooks with disposable endpoints

### 🧑‍💻 Contributing

Contributions are welcome!  
Feel free to open issues or pull requests on [GitHub](https://github.com/egeuysall/rest).

### 📄 License

Licensed under the [Apache License 2.0](./LICENSE).
