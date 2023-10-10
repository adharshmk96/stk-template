import os
from dataclasses import dataclass

ignore_files = ['template-gen.py', 'template.go', "go.mod", "go.sum", "template_map"]

def transform_content(content):
    # Replacing specific strings with given replacements
    content = content.replace('stktemplate', '{{ .AppName }}')
    content = content.replace('github.com/adharshmk96/stk-template/singlemod', '{{ .PkgName }}')
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
    absolute_path: str
    relative_path: str
    var_name: str

module_template_map = []

def find_project_template_map(base_path):
    project_template_map = []
    for root, _, files in os.walk(base_path):
        for f in files:
            if ignore_dir(root) or (f in ignore_files):
                continue
        
            absolute_path = os.path.join(root, f)
            relative_path, var_name = generate_var_name(absolute_path, base_path)
            project_template_map.append(TemplateMap(absolute_path, relative_path, var_name))
    return project_template_map


def write_to_go_template(base_path, target_path):
    project_template_map = find_project_template_map(base_path)
    file_content = "package tpl\n\n"
    for template in project_template_map:
        with open(template.absolute_path, 'r', encoding='utf-8') as input_file:
            content = input_file.read()
            content = transform_content(content)
            structure = f'var {template.var_name} = Template{{\n\tFilePath: "{template.relative_path}",\n\tContent: `{content}`,\n}}\n\n'
            file_content += structure

    with open(target_path, 'w') as output_file:
        output_file.write(file_content)
        output_file.write("var SingleModTemplates = []Template{\n")
        for template in project_template_map:
            output_file.write(f'\t{template.var_name},\n')
        output_file.write("}\n")

if __name__ == '__main__':
    base_directory = "./"
    target_file_path = 'singlemod.go'
    write_to_go_template(base_directory, target_file_path)
    print(f"All files have been written to {target_file_path}.")
