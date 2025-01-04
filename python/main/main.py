import platform as p


class RegionLocatior():
    def find_me(self):

        my_aarchh = p.machine()
        my_system = p.uname().system
        my_node = p.uname().node
        rel_ver = p.uname().version

        print(f"Hi {my_system}, welcome to my {my_node}. My version is {rel_ver} & my architecture is {my_aarchh}")




if __name__ == "__main__":
    try:
        region_locator = RegionLocatior()
        region_locator.find_me()
    except Exception as local_error:
        print(f"An Error Occured : {local_error}")