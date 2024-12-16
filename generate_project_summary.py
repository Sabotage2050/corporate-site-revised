import os
import fnmatch
import chardet
import re

def is_binary(file_path):
    with open(file_path, 'rb') as file:
        return b'\0' in file.read(1024)

def read_file_contents(file_path):
    encodings = ['utf-8', 'shift_jis']
    for encoding in encodings:
        try:
            with open(file_path, 'r', encoding=encoding) as file:
                print(f'Reading file: {file_path}')
                return file.read()
        except UnicodeDecodeError:
            pass
    return ''

def generate_project_summary(target_dir):
    # プロジェクトディレクトリは常にカレントディレクトリ
    project_dir = os.getcwd()
    
    # summaryディレクトリをプロジェクトディレクトリ内に作成
    output_dir = os.path.join(project_dir, 'summary')
    os.makedirs(output_dir, exist_ok=True)
    
    # ターゲットディレクトリ名をサマリーファイル名に使用
    target_name = os.path.basename(target_dir)
    summary = f'# {target_name}\n\n## Directory Structure\n\n'

    gitignore_patterns = read_gitignore(project_dir)
    print(f"gitignore_patterns: {gitignore_patterns}")
    summaryignore_patterns = read_summaryignore(project_dir)
    print(f"summaryignore_patterns: {summaryignore_patterns}")
    additional_ignore_patterns = ['generate_project_summary.py', '.summaryignore', f'{target_name}_project_summary.txt', '.git', 'summary/']

    file_contents_section = "\n## File Contents\n\n"

    def traverse_directory(root, level):
        nonlocal summary, file_contents_section
        indent = '  ' * level
        relative_path = os.path.relpath(root, target_dir)
        if not is_ignored(relative_path, project_dir, gitignore_patterns, summaryignore_patterns, additional_ignore_patterns):
            summary += f'{indent}- {os.path.basename(root)}/\n'

            subindent = '  ' * (level + 1)
            for item in os.listdir(root):
                item_path = os.path.join(root, item)
                if os.path.isdir(item_path):
                    if not is_ignored(item_path, project_dir, gitignore_patterns, summaryignore_patterns, additional_ignore_patterns):
                        traverse_directory(item_path, level + 1)
                else:
                    if not is_ignored(item_path, project_dir, gitignore_patterns, summaryignore_patterns, additional_ignore_patterns):
                        if not is_binary(item_path):
                            summary += f'{subindent}- {item}\n'
                            content = read_file_contents(item_path)
                            if content.strip():
                                # ファイルパスはターゲットディレクトリからの相対パスで表示
                                relative_file_path = os.path.relpath(item_path, target_dir)
                                file_contents_section += f'### {relative_file_path}\n\n```\n{content}\n```\n\n'
                        else:
                            summary += f'{subindent}- {item} (binary file)\n'

    traverse_directory(target_dir, 0)
    
    output_file_path = os.path.join(output_dir, f'{target_name}_project_summary.txt')
    
    with open(output_file_path, 'w', encoding='utf-8') as file:
        file.write(summary + file_contents_section)
    
    print(f"Summary has been generated at: {output_file_path}")

def is_ignored(path, project_dir, gitignore_patterns, summaryignore_patterns, additional_ignore_patterns):
    relative_path = os.path.relpath(path, project_dir)
    all_patterns = gitignore_patterns + summaryignore_patterns + additional_ignore_patterns
    
    for pattern in all_patterns:
        escaped_pattern = re.escape(pattern).replace(r'\*', '.*').replace(r'\?', '.')
        
        if pattern.startswith('/'):
            escaped_pattern = '^' + escaped_pattern[1:]
        elif pattern.endswith('/'):
            escaped_pattern += '(?:/.+)?$'
        else:
            escaped_pattern = '(?:^|/)' + escaped_pattern + '$'

        if re.search(escaped_pattern, relative_path):
            print(f"Ignoring {relative_path} due to pattern: {pattern}")
            return True

    return False

def read_ignore_file(file_path):
    if os.path.exists(file_path):
        with open(file_path, 'r') as file:
            return [line.strip() for line in file if line.strip() and not line.startswith('#')]
    return []

def read_gitignore(project_dir):
    gitignore_path = os.path.join(project_dir, '.gitignore')
    return read_ignore_file(gitignore_path)

def read_summaryignore(project_dir):
    summaryignore_path = os.path.join(project_dir, '.summaryignore')
    return read_ignore_file(summaryignore_path)

if __name__ == '__main__':
    target_directory = input('Enter the target directory path to analyze: ')
    if not target_directory:
        target_directory = os.getcwd()
    
    generate_project_summary(target_directory)