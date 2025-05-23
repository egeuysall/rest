# Rest

Rest is a simple CLI tool for sharing JSON data through one-time, public REST endpoints. It's built for developers who need to temporarily expose or test JSON payloads without setup or authentication. Rest is perfect for quick, disposable JSON sharing during development.

## How It Works

- Use `rest post payload.json` to upload your JSON.
- Get a public URL with `rest url`.
- The endpoint auto-deletes after one view or 15 minutes.
- Optionally delete early with `rest delete <id>`.

## Use Cases

- Share webhook payloads
- Mock API responses
- Test integrations instantly
