import yaml

with open("openapi.yml") as f:
    spec = yaml.safe_load(f)

schemas = spec["components"]["schemas"]
patches = []

for name in schemas:
    if name.startswith("Input") or name.startswith("Output"):
        patches.append({
            "op": "add",
            "path": f"/components/schemas/{name}/properties/status",
            "value": {"$ref": "#/components/schemas/TFStatus"}
        })

overlay = {
    "overrides": [
        {
            "files": ["openapi.yml"],
            "patches": patches
        }
    ]
}

with open("overlay.yml", "w") as f:
    yaml.dump(overlay, f, sort_keys=False)
