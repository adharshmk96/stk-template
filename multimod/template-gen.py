import os
from dataclasses import dataclass

ignore_files = ['template-gen.py', 'template.go', "go.mod", "go.sum", "template_map"]

def transform_content(content):
    # Replacing specific strings with given replacements
    content = content.replace('stktemplate', '{{ .AppName }}')
    content = content.replace('github.com/adharshmk96/stk-template', '{{ .PkgName }}')
    return content

def ignore_dir(dir):
    if ".git" in dir:
        return True

def generate_var_name(file_path, base_path):
    relative_path = file_path.replace(base_path, '').replace('\\', '/')
    var_name = relative_path.replace(os.sep, '_').replace('.', '').strip('_').replace('-', '').replace('/', '')
    var_name = var_name.upper()
    return relative_path, var_name + "_TPL"

@dataclass
class TemplateMap:
    relative_path: str
    var_name: str

template_map = []

def write_to_go_template(base_path, target_path):
    with open(target_path, 'w') as output_file:
        output_file.write("package tpl\n\n")
        # Walk through the base directory
        for root, _, files in os.walk(base_path):
            for f in files:
                if ignore_dir(root):
                    continue

                if f in ignore_files:
                    continue
                file_path = os.path.join(root, f)
                # Create a valid Go variable name from the file path
                relative_path, var_name = generate_var_name(file_path, base_path)              
                
                template_map.append(TemplateMap(relative_path, var_name))

                with open(file_path, 'r', encoding='utf-8') as input_file:
                    content = input_file.read()
                    content = transform_content(content)

                    structure = f'var {var_name} = Template{{\n\tFilePath: "{relative_path}",\n\tContent: `{content}`,\n}}\n\n'
                    
                    # Write variable declaration and content to the output file
                    output_file.write(structure)

        # Write the template list to the output file
        output_file.write("var BoilerPlateTemplates = []Template{\n")
        for template in template_map:
            output_file.write(f'\t{template.var_name},\n')
        output_file.write("}\n")


if __name__ == '__main__':
    base_directory = "./"
    target_file_path = 'multimod.go'
    write_to_go_template(base_directory, target_file_path)
    print(f"All files have been written to {target_file_path}.")
