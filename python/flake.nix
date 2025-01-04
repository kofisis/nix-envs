{
  description = "This creates a python dev env with the specified package and builds the code in it";

  # Flake Inputs
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
    flake-utils.url = "github:numtide/flake-utils"
  };

  outputs = { self, nixpkgs, flake-utils }:
  let
    # Systems supported -> can be tailored to use users os and archtype
    mySystems = [
      "x86_64-linux" # 64-bit Intel/AMD Linux
      "aarch64-linux" # 64-bit ARM Linux
      "x86_64-darwin" # 64-bit Intel macOS
      "aarch64-darwin" # 64-bit ARM macOS
    ];

    # Define package to build (assuming a basic Python package)
    myPackage = pkgs.buildPythonPackage {
      pname = "my-python-package";
      src = ./main; # Path to your Python source code directory
      buildInputs = [ pkgs.python314 ]; # Assuming Python 3.14
      # Add any other necessary build inputs here
    };

    # Helper to get system-specific attribute
    forAllSystems = f:
      nixpkgs.lib.genAttrs mySystems
      (system: f {pkgs = import nixpkgs { inherit system };
      in {
        # My Dev Env
        envShell = forAllSystems ({pkgs}): {
          default = let
            python = pkgs.python314;
          in pkgs.mkshell {
            packages = [
              python.withPackages (ps:
                with ps; [
                  #sys utils
                  ansible-core
                  git
                  ssh
                  black
                ])
            ]
          }
        }
      }});
  in forAllSystems (system: system.envShell);
}