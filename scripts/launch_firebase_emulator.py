import platform
import subprocess
import sys


def run_windows():
    try:
        process = subprocess.Popen(
            'powershell.exe -Command "& {firebase emulators:start}"',
            stdout=sys.stdout
        )
        process.communicate()
    except RuntimeError:
        return 1


if __name__ == '__main__':
    if platform.system() == 'Windows':
        run_windows()
