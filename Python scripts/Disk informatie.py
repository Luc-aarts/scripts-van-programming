import psutil
import platform
from datetime import datetime

#een functie voor het veranderen van bytes een andere formt
def get_size(bytes, suffix="B"):
    """
    
    e.g:
        1253656 => '1.20MB'
        1253656678 => '1.17GB'
    """
    factor = 1024
    for unit in ["", "K", "M", "G", "T", "P"]:
        if bytes < factor:
            return f"{bytes:.2f}{unit}{suffix}"
        bytes /= factor

# een functie voor het verwijderen van "="
def remove_char(string):
    string = list(string)
    string.remove('=')
    return ''.join(string)

# systeem informatie
print("="*40, "System Information", "="*40)
uname = platform.uname()
print(f"System: {uname.system}")
print(f"device Name: {uname.node}")
print(f"windows version: {uname.release}")
print(f"Machine: {uname.machine}")
print(f"Processor: {uname.processor}")


# Disk informatie
print("="*40, "Disk Information", "="*40)
print("Partitions and Usage:")
# alle disk partieties die aanwezig zijn
partitions = psutil.disk_partitions()
for partition in partitions:
    print(f"=== Device: {partition.device} ===")
    print(f"  Mountpoint: {partition.mountpoint}")
    print(f"  File system type: {partition.fstype}")
    try:
        partition_usage = psutil.disk_usage(partition.mountpoint)
    except PermissionError:

        continue
    print(f"  Total Size: {get_size(partition_usage.total)}")
    print(f"  Used: {get_size(partition_usage.used)}")
    print(f"  Free: {get_size(partition_usage.free)}")
    print(f"  Percentage: {partition_usage.percent}%")

#krijg IO-statistieken sinds het opstarten    
disk_io = psutil.disk_io_counters()
print(f"Total read: {get_size(disk_io.read_bytes)}")
print(f"Total write: {get_size(disk_io.write_bytes)}")



