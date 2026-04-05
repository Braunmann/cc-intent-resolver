#### Generate Bindings

```bash
cat contracts/out/IntentHub.sol/IntentHub.json | python3 -c "import json,sys; print(json.dumps(json.load(sys.stdin)['abi']))" > /tmp/IntentHub.abi.json
abigen --abi /tmp/IntentHub.abi.json --pkg bindings --out solver/internal/bindings/IntentHub.go --type IntentHub
```