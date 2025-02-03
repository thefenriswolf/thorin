{ pkgs ? import <nixpkgs> { } }:
let unstable = import <nixos-unstable> { config = { allowUnfree = true; }; };
in pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.go-tools
    pkgs.revive
    #pkgs.tinygo
    # keep this line if you use bash
    pkgs.bashInteractive
  ];
  shellHook = "CGO_ENABLED=0 go build -ldflags '-s -w'";
}
