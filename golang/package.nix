'{ pkgs, pname, src, vendorHash}:

pkgs.buildGoModule {
    inherit pname src vendorHash;
    version = "0.0.3";

    # update this version which the what's in your go.mod file
    # go = pkgs.go_1_25;

    proxyVendor = true; # to fetch a fake vendorHash for go 

    # ensures output of bin/*.exe is same as name the packageName given
    # ldflags = ["-o ${pname}"]; # 
            


    # to ensure nix looks for go.mod file in each src path
    subPackages = [ "." ];

    postInstall = ''
        if [ -f $out/bin/myapp ]; then
            mv $out/bin/myapp $out/bin/${pname}
        fi
  '';
}'