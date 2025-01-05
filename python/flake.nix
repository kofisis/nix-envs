{
  description = "This creates a python dev env with the specified package and builds the code in it";
  # Flake Inputs
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
  # a better way to auto populate the eachSystems is to use
  # flake-utils.lib.eachDefaultSystem -> this auto fills the systems with build systems arch and os
  # https://github.com/numtide/flake-utils/tree/main
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ] (system: let
      pkgs = import nixpkgs { inherit system; };

      myPackage = pkgs.python3.pkgs.buildPythonPackage {
        pname = "myPackage";
        version = "0.1.0";
        src = ./main;
        buildInputs = [ pkgs.python3 ];
      };
    in {
      packages.default = myPackage;

      devShells.default = pkgs.mkShell {
        packages = [
          (pkgs.python3.withPackages (ps: with ps; [
            ansible-core
            black
          ]))
        ];
      };
    });
}