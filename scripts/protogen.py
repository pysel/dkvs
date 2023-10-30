import os

def get_go_module_name(filepath='go.mod'):
    with open(filepath, 'r') as f:
        lines = f.readlines()
        for line in lines:
            if line.startswith('module'):
                return line.split()[1].strip()

def list_files_recursively(start_directory):
    for root, _, files in os.walk(start_directory):
        for filename in files:
            # full path of the file
            full_path = os.path.join(root, filename)
            yield full_path

proto_dir = "proto/dkvs"
go_mod_name = get_go_module_name()
subPaths = []
for filename in list_files_recursively(proto_dir):
    if filename.endswith(".proto"):
        with open(filename) as f:
            lines = f.readlines()
            for line in lines:
                if line.startswith('option go_package'):
                    # Extract the go_package value
                    go_package = line.split('=')[1].strip(' ";').replace('";\n', '')
                    relativeGoPackage = go_package.split(go_mod_name)[1].lstrip('/')
                    print("AAAA", go_package, relativeGoPackage)

                    rmrf = f'rm -rf {relativeGoPackage}'
                    print(f'Running: {rmrf}')
                    os.system(rmrf)
                    
                    if not os.path.exists(go_package):
                        os.system(f'mkdir -p {go_package}')
                        
                    # Build the protoc command
                    cmd = f'protoc -I=proto --go_out=. {filename}'
                    subPaths.append(relativeGoPackage)
                    print(f'Running: {cmd}')
                    os.system(cmd)
                    break

# Remove module prefix and move generated files from there
for path in subPaths:
    # get names of top directories to move
    topDir = path.split('/')[0]
    dirToMove = go_mod_name + "/" + topDir
    cmd = f'(mkdir {topDir}; mv {dirToMove} .)'
    print(f'Running: {cmd}')
    os.system(cmd)

topModDir = go_mod_name.split("/")[0]
rmrfcmd = f"rm -rf {topModDir}"
print(f'Running: {rmrfcmd}')
os.system(rmrfcmd)