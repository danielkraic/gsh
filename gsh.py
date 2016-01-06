#!/usr/bin/env python3

import sys
import server_list
from dialog import Dialog
import subprocess


def choose_server(servers):
    """
    choose server from dialog menu

    :param servers: list of servers
    :return: selected server or None
    """
    if len(servers) == 0:
        return None

    servers_list = []
    for s in matched_servers:
        if scpFlag:
            servers_list.append(tuple([str(len(servers_list)), s.name + ": " + s.get_scp_string()]))
        else:
            servers_list.append(tuple([str(len(servers_list)), s.name + ": " + s.get_ssh_string()]))

    d = Dialog(dialog="dialog")
    code, server_tag = d.menu("Select server", choices=servers_list)

    if code != d.OK:
        return None

    return matched_servers[int(server_tag)]


def connect(server, scp):
    """
    execute ssh/scp command

    :param server: server to connect
    :param scp: connect with scp if True
    """
    command = ''
    if scp:
        command = server.get_scp_string()
    else:
        command = server.get_ssh_string()

    print("command: ", command)
    subprocess.call(command, shell=True)


if __name__ == '__main__':
    usageString = "Usage: {0} [-c] servername".format(sys.argv[0])

    if len(sys.argv) == 1:
        print(usageString)
        sys.exit(0)

    args = sys.argv[1:]

    scpFlag = False
    if args[0] == '-c':
        scpFlag = True
        args = args[1:]

    if len(args) == 0:
        print(usageString)
        sys.exit(0)

    matched_servers = [server for server in server_list.all_servers if server.match(args)]

    if len(matched_servers) == 0:
        print("no server matched")
        sys.exit(0)

    if len(matched_servers) == 1:
        connect(matched_servers[0], scpFlag)
    else:
        selected_server = choose_server(matched_servers)
        if selected_server is None:
            sys.exit(1)

        connect(selected_server, scpFlag)
