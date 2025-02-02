{ stdenvNoCC, lib, fetchFromGitHub, makeWrapper, buildGoModule }:

buildGoModule rec {
  pname = "thorin";
  version = "v20250202";

  src = fetchFromGitHub {
    owner = "thefenriswolf";
    repo = "thorin";
    rev = "${version}";
    hash = "";
  };

  meta = with lib; {
    description = "Thorin money balance tool";
    homepage = "https://github.com/thefenriswolf/thorin";
    license = licenses.gpl3;
    maintainers = with maintainers; [ thefenriswolf ];
  };
}

