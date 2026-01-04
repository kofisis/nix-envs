{ pkgs, pname, src, vendorHash}:

pkgs.buildGoModule {
    inherit pname src vendorHash;
    version = "0.0.1";

    # update this version which the what's in your go.mod file
    # go = pkgs.go_1_25;

    # to ensure nix looks for go.mod file in each src path
    subPackages = [ "." ];
}