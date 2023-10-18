import os

def get_go_module_name(filepath='go.mod'):
    with open(filepath, 'r') as f:
        lines = f.readlines()
        for line in lines:
            if line.startswith('module'):
                return line.split()[1].strip()
            
proto_dir = "proto/dkvs"
go_mod_name = get_go_module_name()

for filename in os.listdir(proto_dir):
    if filename.endswith(".proto"):
        with open(os.path.join(proto_dir, filename)) as f:
            lines = f.readlines()
            for line in lines:
                if line.startswith('option go_package'):
                    # Extract the go_package value
                    go_package = line.split('=')[1].strip(' ";').replace('";\n', '')
                    print(f'go_package: {go_package}')
                    
                    if not os.path.exists(go_package):
                        os.system(f'mkdir -p {go_package}')
                        
                    # Build the protoc command
                    cmd = f'protoc -I=proto --go_out=. {proto_dir}/{filename}'
                    print(f'Running: {cmd}')
                    os.system(cmd)
                    break