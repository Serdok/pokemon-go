import platform
import subprocess
import sys


def run_windows():
    process = subprocess.Popen(
        'powershell.exe -Command "& {firebase emulators:start}"',
        stdout=sys.stdout
    )
    process.communicate()


if __name__ == '__main__':
    if platform.system() == 'Windows':
        run_windows()
