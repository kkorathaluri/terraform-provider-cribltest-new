import yaml
import sys

def get_schema_names_from_overlay(overlay_file):
    with open(overlay_file, 'r') as f:
        overlay = yaml.safe_load(f)
    
    schema_names = set()
    for override in overlay.get('overrides', []):
        for patch in override.get('patches', []):
            # Extract schema name from path like /components/schemas/InputKafka/properties/status
            path_parts = patch.get('path', '').split('/')
            if len(path_parts) > 3:
                schema_names.add(path_parts[3])
    return schema_names

def update_schema(schema):
    if isinstance(schema, dict):
        if 'properties' not in schema:
            schema['properties'] = {}
        schema['properties']['status'] = {"$ref": "#/components/schemas/TFStatus"}
    return schema

def process_file(filename, schema_names):
    data = yaml.safe_load(open(filename, "r"))
    for k, v in data.get("components", {}).get("schemas", {}).items():
        if k in schema_names:
            data["components"]["schemas"][k] = update_schema(v)
    yaml.dump(data, open(filename, "w"), sort_keys=False)

if __name__ == "__main__":
    schema_names = get_schema_names_from_overlay("overlay.yml")
    process_file("openapi.yml", schema_names)
