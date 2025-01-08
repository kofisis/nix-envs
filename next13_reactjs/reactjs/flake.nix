{
  description = "Nix flake to build devShell for node react projects";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # install node20
            nodejs_20
            # package managers
            nodePackages.yarn
            nodePackages.npm
            # system tools
            git
            watchman
          ];
          shellHook = '' # Commands to run postShell activation
            echo "Welcome to reatJS Dev Environment"
            echo "NodeJS Version: $(node --version)"
            echo "NPM Version: $(npm --version)"
            # Run the creation of a nextJS app if package.json not found
            if [ ! -f package.json ]; then
              echo "No package.json found. You can create a new project by running:"
              echo "npx create-react-app"
              echo "or: "
              echo "npm create vite"
            fi
          '';
        };
      }
    );
}