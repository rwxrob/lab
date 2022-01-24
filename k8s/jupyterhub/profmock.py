#!/usr/bin/python3

class SimulatedSpawner:
    image = ''
    cpu_guarantee = 1
    user_options = {}

class ResourceProfile:
    name = 'default'
    description = ''
    image = 'jupyter/datascience-notebook:latest'
    cpu_guarantee = 1.0
    mem_guarantee = 1.0

default = ResourceProfile()

base = ResourceProfile()
base.name = 'base'
base.image = 'jupyter/base-notebook:latest'

def opt_sync(spawner):
    '''Sets the properties of the spawner to the user_options value of
    the same name'''
    pprint(base)

if __name__ == '__main__':
    sp = SimulatedSpawner()
    opt_sync(sp)

# async def custom_pre_spawn_hook(sp):
#   sp.image = sp.user_options.get("image",prof["def"]["image"])
#   sp.cpu_guarantee = sp.user_options.get("cpu_guarantee","1")

