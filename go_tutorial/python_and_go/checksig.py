import ctypes
from distutils.sysconfig import get_config_var
from pathlib import Path

# Location of shared library
here = Path(__file__).absolute().parent
ext_suffix = get_config_var('EXT_SUFFIX')
so_file = here / ('_checksig' + ext_suffix)

# Load functions from shared library set their signatures
so = ctypes.cdll.LoadLibrary(so_file)
verify = so.verify
verify.argtypes = [ctypes.c_char_p]
verify.restype = ctypes.c_void_p
free = so.free
free.argtypes = [ctypes.c_void_p]


def check_signatures(root_dir):
    """Check (in parallel) digital signature of all files in root_dir.
    We assume there's a sha1sum.txt file under root_dir
    """
    res = verify(root_dir.encode('utf-8'))
    if res is not None:
        msg = ctypes.string_at(res).decode('utf-8')
        free(res)
        raise ValueError(msg)

if __name__ == "__main__":
    check_signatures("/<dir>/logs")