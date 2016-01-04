#!/usr/bin/env python3

import sys

class Server:
    '''server class'''

    DEFAULT_USER = 'user'
    DEFAULT_PORT = 22

    def __init__(self, name, host, user=DEFAULT_USER, port=DEFAULT_PORT):
        self.name = name
        self.host = host
        self.user = user
        self.port = port

    def getSsh(self):
        return "ssh -p {0} {1}@{2} # {3}".format(self.port, self.user, self.host, self.name)

    def getScp(self):
        return "scp -P {0} {1}@{2} # {3}".format(self.port, self.user, self.host, self.name)

    def match(self, patters):
        all_matched = True;

        for p in patters:
            p_matched = False
            if not p_matched and self.name.find(p) != -1:
                p_matched = True
            if not p_matched and self.host.find(p) != -1:
                p_matched = True
            if not p_matched and self.user.find(p) != -1:
                p_matched = True

            all_matched = all_matched and p_matched

        return all_matched


class ServerSearch:
    '''Search server class'''

    all_servers = [
        Server(name='server1', host='host1', user='user1', port=22),
        Server('server2', 'host2', 'user2', 22),
        Server('server3', 'host3')
    ]

    def getServers(self, patterns):
        if len(patterns) == 0:
            return []

        matched_servers = []

        for server in ServerSearch.all_servers:
            if server.match(patterns):
                matched_servers.append(server)

        return matched_servers


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

    matched_servers = ServerSearch().getServers(args)

    if len(matched_servers) == 0:
        print("no server matched")

    for server in matched_servers:
        if scpFlag:
            print(server.getScp())
        else:
            print(server.getSsh())
