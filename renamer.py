"""
Script to add docnos to files in c4/no.clean
To process all files:
python renamer.py --path <path-to-c4-repo>
To process a subset, e.g. the first 20 files:
python renamer.py --path <path-to-c4-repo> --pattern 000[01]?
"""
import argparse
import glob
import gzip

parser = argparse.ArgumentParser(description='Add docnos to C4 collection.')
parser.add_argument('--path', type=str, help='Root of C4 git repo.', required=True)
parser.add_argument('--pattern', type=str, default="?????", help='File name patterns to process.')
args = parser.parse_args()
pattern = args.pattern
path = args.path


def new_docno(file_number, line_number):
    return f'en.noclean.c4-train.{file_number}-of-07168.{line_number}'


files = sorted(list(glob.iglob(f'{path}/en.noclean/c4-train.{pattern}-of-07168.json.gz')))

for filepath in files:
    with gzip.open(filepath) as f:
        file_number = filepath[-22:-22 + 5]
        file_name = filepath[-31:]
        print(f"adding docnos to file number {file_number} ...")
        with gzip.open(f'{path}/en.noclean.withdocnos/{file_name}', 'wb') as o:
            for line_number, line in enumerate(f.readlines()):
                line = line.decode('utf-8')
                new_line = f"{{\"docno\":\"{new_docno(file_number, line_number)}\",{line[1:]}"
                o.write(new_line.encode('utf-8'))
