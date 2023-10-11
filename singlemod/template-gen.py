import os
from dataclasses import dataclass

ignore_files = ['template-gen.py', 'template.go', "go.mod", "go.sum", "template_map"]
ignore_dirs = ['.git', 'mocks']

@dataclass
class TemplateMap:
    absolute_path: str
    relative_path: str
    var_name: str

def replace_pkg_app_names(content):
    # Replacing specific strings with given replacements
    content = content.replace('stktemplate', '{{ .AppName }}')
    content = content.replace('github.com/adharshmk96/stk-template/singlemod', '{{ .PkgName }}')
    return content

def replace_module_names(content):
    # Replacing specific strings with given replacements
    content = content.replace('ping', '{{ .ModName }}')
    content = content.replace('Ping', '{{ .ExportedName }}')
    return content

def ignore_dir(dir):
    for ignore_dir in ignore_dirs:
        if ignore_dir in dir:
            return True
    return False

def generate_var_name(file_path, base_path):
    relative_path = file_path.replace(base_path, '').replace('\\', '/')
    var_name = relative_path.replace(os.sep, '_').replace('.', '').strip('_').replace('-', '').replace('/', '')
    var_name = var_name.upper()
    return relative_path, var_name

def find_project_template_map(base_path):
    project_template_map = []
    for root, _, files in os.walk(base_path):
        for f in files:
            if ignore_dir(root) or (f in ignore_files):
                continue
                   
            absolute_path = os.path.join(root, f)
            relative_path, var_name = generate_var_name(absolute_path, base_path)
            project_template_map.append(TemplateMap(absolute_path, relative_path, var_name + "_TPL"))
    return project_template_map

def generate_project_template(base_path):
    project_template_map = find_project_template_map(base_path)
    project_template = "package tpl\n\n"
    for template in project_template_map:
        with open(template.absolute_path, 'r', encoding='utf-8') as input_file:
            content = input_file.read()
            content = replace_pkg_app_names(content)
            structure = f'var {template.var_name} = Template{{\n\tFilePath: "{template.relative_path}",\n\tContent: `{content}`,\n}}\n\n'
            project_template += structure
    return project_template

def write_project_template_to_file(target_path, file_content, project_template_map):
    with open(target_path, 'w') as output_file:
        output_file.write(file_content)
        output_file.write("var SingleModTemplates = []Template{\n")
        for template in project_template_map:
            output_file.write(f'\t{template.var_name},\n')
        output_file.write("}\n")

def create_project_template(base_directory, target_file_path):
    project_template_map = find_project_template_map(base_directory)
    project_template = generate_project_template(base_directory)
    write_project_template_to_file(target_file_path, project_template, project_template_map)


def find_module_template_map(base_path):
    module_template_map = []
    for root, _, files in os.walk(base_path):
        for f in files:
            if ("ping" not in f) or ignore_dir(root) or (f in ignore_files):
                continue

            absolute_path = os.path.join(root, f)
            relative_path, var_name = generate_var_name(absolute_path, base_path)
            module_template_map.append(TemplateMap(absolute_path, relative_path, var_name+"_MOD"))
    return module_template_map

def generate_module_template(base_path):
    module_template_map = find_module_template_map(base_path)
    module_template = "package tpl\n\n"
    for template in module_template_map:
        with open(template.absolute_path, 'r', encoding='utf-8') as input_file:
            content = input_file.read()
            content = replace_pkg_app_names(content)
            content = replace_module_names(content)

            file_path = replace_module_names(template.relative_path)

            structure = f'var {template.var_name} = Template{{\n\tFilePath: "{file_path}",\n\tContent: `{content}`,\n}}\n\n'
            module_template += structure
    return module_template

def write_module_template_to_file(target_path, file_content, module_template_map):
    with open(target_path, 'w') as output_file:
        output_file.write(file_content)
        output_file.write("var ModuleTemplates = []Template{\n")
        for template in module_template_map:
            output_file.write(f'\t{template.var_name},\n')
        output_file.write("}\n")

def create_module_template(base_directory, target_file_path):
    module_template_map = find_module_template_map(base_directory)
    module_template = generate_module_template(base_directory)
    write_module_template_to_file(target_file_path, module_template, module_template_map)

def remove_files(files):
    for file in files:
        try:
            os.remove(file)
        except:
            pass

if __name__ == '__main__':
    base_directory = "./"
    project_template_path = 'singlemod.go'
    module_template_path = 'ping.go'
    remove_files([project_template_path, module_template_path])
    create_project_template(base_directory, project_template_path)
    create_module_template(base_directory, module_template_path)
    print(f"All files have been written to {project_template_path}, {module_template_path}.")
