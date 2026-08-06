{
  pkgs,
  ...
}:

{
  packages = [ pkgs.git ];

  languages.go = {
    enable = true;
    enableHardeningWorkaround = true;
    lsp = {
      enable = true;
    };
    delve = {
      enable = true;
    };
  };

  enterShell = ''
    go version
  '';

  # https://devenv.sh/tests/
  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
    go test -v ./...
  '';

  # https://devenv.sh/git-hooks/
  # git-hooks.hooks.shellcheck.enable = true;

  # See full reference at https://devenv.sh/reference/options/
}
