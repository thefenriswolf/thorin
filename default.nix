{ stdenvNoCC, lib, fetchFromGitHub, makeWrapper, buildGoModule }:

#with import <nixpkgs> { };

buildGoModule rec {
  pname = "thorin";
  version = "5b4d46f0e05c3f88ba1567462189a7236e244aff";

  src = fetchFromGitHub {
    owner = "thefenriswolf";
    repo = "thorin";
    rev = "${version}";
    hash = "sha256-4zXFxSawLI3VXVZKqiLPIrOfwVXlKFT/92ZO56XSoz0";
  };
  vendorHash = null; # go stdlib does not need a vendor hash

  ldflags = [ "-s -w" ];
  CGO_ENABLED = 0;

  meta = with lib; {
    description = "Thorin money balance tool";
    homepage = "https://github.com/thefenriswolf/thorin";
    license = licenses.gpl3;
    maintainers = with maintainers; [ thefenriswolf ];
  };
}

