import getpass


class Server:
    """
    server class
    """
    default_user = getpass.getuser()
    default_port = 22

    def __init__(self, name, host, user=default_user, port=default_port):
        self.name = name
        self.host = host
        self.user = user
        self.port = port

    def get_ssh_string(self):
        return "ssh -p {0} {1}@{2} # {3}".format(self.port, self.user, self.host, self.name)

    def get_scp_string(self):
        return "scp -P {0} {1}@{2} # {3}".format(self.port, self.user, self.host, self.name)

    def get_pretty_string(self):
        return "{0}: ssh -p {1} {2}@{3}".format(self.name, self.port, self.user, self.host)

    def match(self, patters):
        all_matched = True

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
