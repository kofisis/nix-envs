
/* Personal Nix Flakes with a managed package.nix file to build and run my small o practice applications
nix run .#task_no -> to run app from task_no/code
nix flake update -> to update flake should anything be changed in testing


 */


{
  description = "A monorepo setup for all my go tutorial projects";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };

        makeTask = dirName: hash: pkgs.callPackage ./package.nix {
          pname = dirName;
          # If this flake is meant to manage many tasks, it should be one level up.
          src = ./${dirName}/code;
          vendorHash = hash;
        };
      in
      {
        packages = {
          task1 = makeTask "task_1" "sha256-AQSGryR9YkZH8FlB/DS1cw8akCW73674JnrLHB1KBB4=";
          # task2 = makeTask "task_2" null;

          default = self.packages.${system}.task1;
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            golangci-lint
          ];

          shellHook = ''echo "Ready to work on tasks"'';
        };
      });
}